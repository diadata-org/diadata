package main

// ----------------------------------------------------------------------------------
// THIS FILE IS JUST FOR TESTING PURPOSES AND CAN BE DELETED ONCE HASHING IS FINISHED
// ----------------------------------------------------------------------------------

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/cbergoon/merkletree"
	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	// Determine frequency of scraping
	refreshDelay = time.Second * 10 * 1
)

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

// FlushPool flushes pools coming through a channel: It stores the pool in a database
// and makes a merkle Tree.
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

func main() {

	// Test verification ----------------------------------------------
	dataType := flag.String("type", "hash-interestrates", "Type of data")
	flag.Parse()

	ds, err := models.NewAuditStore()
	if err != nil {
		log.Error("NewAuditStore: ", err)
	}

	poolNum := "10"
	tree, err := ds.GetMerkletreeByID(*dataType, poolNum)
	if err != nil {
		log.Error(err)
	}
	num := 3
	fmt.Printf("bucket type %T and value %v\n", tree.Leafs[num].C.(merkletree.StorageBucket).ID, tree.Leafs[num].C.(merkletree.StorageBucket).ID)
	verif, _ := merklehashing.VerifyBucket(tree.Leafs[num].C.(merkletree.StorageBucket))
	fmt.Println("verification: ", verif)
	newBucket := tree.Leafs[num].C.(merkletree.StorageBucket)
	verifInterm, _ := merklehashing.VerifyBucket(newBucket)
	fmt.Println("verification before modification: ", verifInterm)
	newBucket.Content[0] = 1
	verifnew, _ := merklehashing.VerifyBucket(newBucket)
	fmt.Println("new verification: ", verifnew)

	verif2, _ := merklehashing.VerifyPool(tree, *dataType, "0")
	fmt.Println("verification of pool: ", verif2)

	level2Tree, err := ds.GetDailyTreeByID(*dataType, "2", "0")
	if err != nil {
		log.Error(err)
	}
	verif3, err := merklehashing.VerifyTree(level2Tree, "2", "0")
	fmt.Println(verif3, err)

	level1Tree, err := ds.GetDailyTreeByID("", "1", "0")
	if err != nil {
		log.Error(err)
	}
	fmt.Println("level1tree root: ", level1Tree.MerkleRoot)
	verif4, err := merklehashing.VerifyTree(level1Tree, "1", "1")
	fmt.Println(verif4, err)

	level0Tree, err := ds.GetDailyTreeByID("", "0", "0")
	if err != nil {
		log.Error(err)
	}
	fmt.Println("leaf 1 duplicated: ", level0Tree.Leafs[0])
	fmt.Println("leaf 2 duplicated: ", level0Tree.Leafs[1])
	// ---------------------------------------------------------------

	// level0Tree, _ := ds.GetDailyTreeByID("", "0", "0")

	// timeInit := time.Now().Add(time.Hour * (-800))
	// timeFinal := time.Now()
	// retval, err := ds.GetMerkletreeInflux(*dataType, timeInit, timeFinal)
	// if err != nil {
	// 	log.Error("error getting merkle tree from influx: ", err)
	// }
	// bucket := retval[0].Root.Left.Left.C.(merkletree.StorageBucket)
	// data, err := bucket.ReadContent()
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(string(data[i]))
	// }

	// vals, err := ds.GetMerkletreesInflux(*dataType, timeInit, timeFinal)
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println("vals are: ", vals[0][1])
	// myTree := merkletree.MerkleTree{}
	// json.Unmarshal([]byte(vals[0][2].(string)), &myTree)

	// ds.SaveDailyTreeInflux(myTree, "", "1")
	// ds.SaveDailyTreeInflux(myTree, "", "1")
	// ds.SaveDailyTreeInflux(myTree, "", "1")

	// -------------------------------------------------------------

	// // Run this section for data storage in hashing tables
	// // One instance of main for each data type
	// dataType := flag.String("type", "hash-trades", "Type of data")
	// flag.Parse()

	// kc := ActivateKafkaChannel(*dataType)
	// defer kc.Close()

	// ds, err := models.NewInfluxAuditStore()
	// if err != nil {
	// 	log.Error("NewInfluxDataStore: ", err)
	// }

	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// pChan := make(chan *merkletree.BucketPool)
	// go FillPools(*dataType, 4, 512, pChan, kc.chanMessage, &wg)

	// wg.Add(1)
	// go FlushPool(pChan, &wg, ds)
	// defer wg.Wait()

	// -------------------------------------------------------------

	// dataType := flag.String("type", "mytopic", "Type of data")
	// flag.Parse()

	// kc := ActivateKafkaChannel(*dataType)
	// defer kc.Close()

	// bp := models.NewBucketPool(2, 256)
	// fmt.Println("length of newly created bucketpool: ", bp.Len())
	// full := FillBucketPool(bp, kc.chanMessage)
	// fmt.Println(bp.Len())
	// bucket, _ := bp.Get()
	// fmt.Println(bucket.Content)
	// bucket, _ = bp.Get()
	// fmt.Println(bucket.Content)

	// fmt.Println(full)

	// bp := models.NewBucketPool(2, 2)
	// bucket, err := bp.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// ok := bucket.WriteContent([]byte("hi"))
	// fmt.Println(ok)
	// bp.Put(bucket)
	// bucket, err = bp.Get()
	// fmt.Println(bucket.Used())
	// ok = bucket.WriteContent([]byte("ya"))
	// fmt.Println(bucket.Content)
	// bp.Put(bucket)
	// fmt.Println("length of pool: ", bp.Len())

	// filledbucket, err := bp.Get()
	// fmt.Println(filledbucket.Used())
	// fmt.Println(filledbucket.Content)

	// filledbucket2, err := bp.Get()
	// fmt.Println(filledbucket2.Used())
	// fmt.Println(filledbucket2.Content)

	// filledbucket3, err := bp.Get()
	// fmt.Println(err)
	// fmt.Println(filledbucket3.Used())
	// fmt.Println(filledbucket3.Content)

	// dataType := flag.String("type", "mytopic", "Type of data")
	// flag.Parse()

	// kc := ActivateKafkaChannel(*dataType)
	// defer kc.Close()

	// bp := models.NewBucketPool(2, 256)
	// fmt.Println("length of newly created bucketpool: ", bp.Len())

	// ok := true
	// bucket, err := bp.Get()
	// for i := 0; i < 5; i++ {
	// 	message := <-kc.chanMessage
	// 	fmt.Printf("message %v: %v \n", i, message.Value)
	// 	if ok {
	// 		ok = bucket.WriteContent(message.Value)
	// 		fmt.Println("if condition")
	// 		// fmt.Println("bucket content: ", bucket.Content)
	// 	} else {
	// 		fmt.Println("bucket full. Return bucket to pool.")
	// 		bp.Put(bucket)
	// 		bucket, err = bp.Get()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 			// TO DO: Hash the bucket pool and make a new one
	// 		}
	// 		// fmt.Println("new bucket's content: ", bucket.Content)
	// 		ok = bucket.WriteContent(message.Value)
	// 	}
	// }
	// fmt.Println("bucket before put: ", bucket)
	// bp.Put(bucket)
	// fmt.Println("number of buckets: ", bp.Len())
	// bucket1, _ := bp.Get()
	// bucket2, _ := bp.Get()
	// fmt.Println("bucket1: ", bucket1.Used())
	// fmt.Println("bucket2: ", bucket2.Used())

	// ir := models.InterestRate{}
	// err := ir.UnmarshalBinary(val)
	// fmt.Println(err, ir)

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

// // FillBucketPool fills a bucket pool of finite size. It returns true when the pool is full.
// func FillBucketPool(bp *models.BucketPool, topicChan chan *kafka.Message) bool {

// 	used := false
// 	for !used {
// 		// Fill pool as long as unused buckets are in the channel
// 		bucket, err := bp.Get()
// 		if err != nil {
// 			log.Error(err)
// 			return true
// 		}
// 		used = bucket.Used()
// 		ok := true
// 		for ok {
// 			// Fill bucket as long as not full
// 			message := <-topicChan
// 			ok = bucket.WriteContent(message.Value)
// 		}
// 		bp.Put(bucket)
// 	}
// 	return used
// }

// func FillCategory(pools []*models.BucketPool, topic string, wg *sync.WaitGroup) {

// 	defer wg.Done()
// 	kc := ActivateKafkaChannel(topic)
// 	defer kc.Close()
// 	wg.Add(1)
// 	defer wg.Wait()
// 	for {
// 		bp := models.NewBucketPool(8, 2560)
// 		go FillBucketPool(bp, kc.chanMessage)
// 		pools = append(pools, bp)
// 	}

// }
