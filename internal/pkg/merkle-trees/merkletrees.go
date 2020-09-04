package merklehashing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/cbergoon/merkletree"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	// Determine frequency of scraping
	refreshDelay = time.Second * 10 * 1
)

// GetHashTopics returns a map listing all hashing topics
func GetHashTopics() map[int]string {
	topicMap := map[int]string{
		0: "hash-interestrates",
		1: "hash-trades",
	}
	return topicMap
}

// GetNumTopics returns the number of hashing topics
func GetNumTopics() int {
	return len(GetHashTopics())
}

type nothing struct{}

// KafkaChannel (rename) is basically a channel for kafka messages
type KafkaChannel struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock   sync.RWMutex
	error       error
	closed      bool
	ticker      *time.Ticker
	datastore   models.Datastore
	chanMessage chan *kafka.Message
}

// StartKafkaReader starts a kafka reader that listens to @topic and
// sends the messages to the channel of KafkaChannel
func (kc *KafkaChannel) StartKafkaReader(topic string) {
	config := kafka.ReaderConfig{
		// TO DO: Production switch
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		MaxBytes: 10,
	}

	reader := kafka.NewReader(config)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error ocurred: ", err)
			continue
		}
		kc.chanMessage <- &m
	}
}

// ActivateKafkaChannel makes a new KafkaChannel struct and gets continuous
// input to its channel from a kafka reader listening to @topic in a go routine
func ActivateKafkaChannel(topic string) *KafkaChannel {
	kc := &KafkaChannel{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		ticker:       time.NewTicker(refreshDelay),
		chanMessage:  make(chan *kafka.Message),
	}

	log.Info("KafkaReader is built and triggered")
	go kc.StartKafkaReader(topic)
	return kc
}

// FillPools streams data from the kafka channel into pools and directs
// them into @poolChannel to be flushed afterwards.
func FillPools(topic string, numBucket, sizeBucket uint64, poolChannel chan *merkletree.BucketPool, topicChan chan *kafka.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	bp := merkletree.NewBucketPool(numBucket, sizeBucket, topic)
	ok := true
	bucket, err := bp.Get()
	if err != nil {
		log.Error("error with Get bucket: ", err)
	}

	// Fill pools with messages from kafka channel
	for {
		message := <-topicChan
		if ok {
			ok = bucket.WriteContent(message.Value)

		} else {
			fmt.Println("bucket full. Return bucket to pool.")

			// TO DO: put the timestamping into the Put() function?
			bucket.Timestamp = time.Now()
			bp.Put(bucket)

			// Check, whether there is a fresh bucket in the pool
			bucket, err = bp.Get()
			if err != nil {
				// In this case, the pool is exhausted. Flush it and make a new one...
				poolChannel <- bp
				bp = merkletree.NewBucketPool(numBucket, sizeBucket, topic)
				fmt.Println(err)
			}
			// ... otherwise go on filling the fresh bucket.
			ok = bucket.WriteContent(message.Value)
			fmt.Println("new content written: ", message.Value)

		}
	}
}

// FlushPool flushes pools coming through a channel: It turns the pool into a merkle Tree
// and stores the tree in influx.
func FlushPool(poolChannel chan *merkletree.BucketPool, wg *sync.WaitGroup, ds models.AuditStore) {

	for {

		bp := <-poolChannel
		tree, err := merkletree.MakeTree(bp)
		if err != nil {
			log.Error(err)
			return
		}
		// Once a day, a script fetches all entries with today's date. Ordering of trees can be done
		// with influx timestamps. Ordering in merkle tree can be done using timestamps of buckets.
		err = ds.SaveMerkletreeInflux(*tree, bp.Topic)
		if err != nil {
			log.Error("error saving tree to influx: ", err)
		}

	}
}

// HashPoolLoop opens a kafka channel for data of type @topic and fills and saves bucketpools with
// the corresponding marshalled data in influx.
func HashPoolLoop(topic string) {

	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}

	kc := ActivateKafkaChannel(topic)
	defer kc.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	pChan := make(chan *merkletree.BucketPool)
	go FillPools(topic, 4, 512, pChan, kc.chanMessage, &wg)

	wg.Add(1)
	go FlushPool(pChan, &wg, ds)
	defer wg.Wait()
}

// DailyTreeTopic retrieves all merkle trees corresponding to @topic from influx and
// hashes them in a merkle tree. The tree's (influx-)timestamps are ranging until at most @timeFinal.
// The root hash of the resulting merkle tree is returned.
// This functionality implements Level2 from the Merkle Documentation.
func DailyTreeTopic(topic string, timeFinal time.Time) (dailyTopicTree *merkletree.MerkleTree, err error) {
	level := "2"
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	// Get last timestamp of trees from hashing table
	timeInit, err := ds.GetLastTimestamp(topic, level)
	if err != nil {
		log.Error(err)
	}
	// Get merkle trees from the data storage table
	vals, err := ds.GetMerkletreesInflux(topic, timeInit, timeFinal)
	if err != nil {
		log.Error(err)
	}
	var merkleTrees []merkletree.MerkleTree
	var lastTimestamp time.Time
	for i := range vals {
		// Collect merkle trees
		var auxTree merkletree.MerkleTree
		err = json.Unmarshal([]byte(vals[i][2].(string)), &auxTree)
		if err != nil {
			log.Error(err)
			return
		}
		merkleTrees = append(merkleTrees, auxTree)
		// Find last timestamp. It will be the initial time for the next iteration.
		tstamp, _ := time.Parse(time.RFC3339, vals[i][0].(string))
		if tstamp.After(lastTimestamp) {
			lastTimestamp = tstamp
		}
	}
	dailyTopicTree, err = merkletree.TreesToTree(merkleTrees)
	if err != nil {
		log.Error(err)
		return
	}

	err = ds.SaveDailyTreeInflux(*dailyTopicTree, topic, level, lastTimestamp)
	return
}

// DailyTree returns a merkle tree which is constructed from the root hashes of the DailyTopicTrees.
// It includes all Level2 trees which have not been hashed into a Level1 tree yet, up to timeFinal.
// This functionality implements Level1 from the Merkle Documentation
func DailyTree(timeFinal time.Time) (dailyTree *merkletree.MerkleTree, err error) {
	level := "1"
	var dailyTrees []merkletree.MerkleTree

	// Retrieve daily trees for all topics
	numTopics := GetNumTopics()
	topicMap := GetHashTopics()
	for i := 0; i < numTopics; i++ {
		topic := topicMap[i]
		dailyTopicTree, err := DailyTreeTopic(topic, timeFinal)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("daily topic tree: ", dailyTopicTree)
		dailyTrees = append(dailyTrees, *dailyTopicTree)
	}
	dailyTree, err = merkletree.TreesToTree(dailyTrees)
	if err != nil {
		return
	}

	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	err = ds.SaveDailyTreeInflux(*dailyTree, "", level, time.Time{})
	return
}

// MasterTree returns the master merkle tree resulting from the (daily) hashing procedure
func MasterTree() (masterTree merkletree.MerkleTree, err error) {
	level := "0"
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
		return
	}

	// Get today's merkle root from the level 1 tree
	timestamp := time.Now()
	dailyTree, err := DailyTree(timestamp)
	if err != nil {
		log.Error(err)
		return
	}
	dailyRootHash := dailyTree.MerkleRoot

	// Get last master tree
	lastID, err := ds.GetLastIDMerkle("", level)
	if err != nil {
		log.Error(err)
		return
	}
	ID := strconv.Itoa(int(lastID))
	masterTree, err = ds.GetDailyTreeByID("", level, ID)
	if err != nil {
		log.Error(err)
		return
	}

	// Extend master tree by today's merkle root
	newHash := merkletree.ByteContent(dailyRootHash)
	err = masterTree.ExtendTree([]merkletree.Content{newHash})
	if err != nil {
		log.Error(err)
		return
	}
	// Save newMasterTree
	ds.SaveDailyTreeInflux(masterTree, "", level, time.Time{})

	return
}

// Close closes the channel of KafkaChannel @kc if not done yet
func (kc *KafkaChannel) Close() error {
	if kc.closed {
		return errors.New("TopicSwitch: Already closed")
	}
	close(kc.shutdown)
	<-kc.shutdownDone
	kc.errorLock.RLock()
	defer kc.errorLock.RUnlock()
	return kc.error
}
