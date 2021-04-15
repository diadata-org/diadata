package main

// ----------------------------------------------------------------------------------
// THIS FILE IS JUST FOR TESTING PURPOSES AND CAN BE DELETED ONCE HASHING IS FINISHED
// ----------------------------------------------------------------------------------

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cbergoon/merkletree"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/go-redis/redis/v8"
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
		err = ds.SetStorageTreeInflux(*tree, bp.Topic)
		if err != nil {
			log.Error("error saving tree to influx: ", err)
		}

	}
}

type TestSHA256Content struct {
	x string
}

//CalculateHash hashes the values of a TestSHA256Content
func (t TestSHA256Content) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t TestSHA256Content) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestSHA256Content).x, nil
}

// makeTestTree makes a tree with two nodes
func makeTestTree(numpools uint64, ContentNode1, ContentNode2 []string) (*merkletree.MerkleTree, error) {
	bp := merkletree.NewBucketPool(numpools, 216, "test")
	bucket1, err := bp.Get()
	if err != nil {
		fmt.Println("error: ", err)
	}
	for _, content := range ContentNode1 {
		bucket1.WriteContent([]byte(content))
	}
	bucket2, err := bp.Get()
	if err != nil {
		fmt.Println("error: ", err)
	}
	for _, content := range ContentNode2 {
		bucket2.WriteContent([]byte(content))
	}
	if bp.Put(bucket1) {
		fmt.Println("put bucket p")
	}
	if bp.Put(bucket2) {
		fmt.Println("put bucket p2")
	}
	return merkletree.MakeTree(bp)
}

func marshAndUnmarshTree(tree merkletree.MerkleTree) (merkletree.MerkleTree, error) {
	marshTree, err := json.Marshal(&tree)
	if err != nil {
		return merkletree.MerkleTree{}, err
	}
	var recoveredTree merkletree.MerkleTree
	err = json.Unmarshal(marshTree, &recoveredTree)
	if err != nil {
		return merkletree.MerkleTree{}, err
	}
	return recoveredTree, nil
}

// GetPoolsParentID returns the ID of level 2 tree such that hashed pool with @id is a leaf
func GetPoolsParentID(id, topic string, r *redis.Client) (string, error) {
	key := getKeyPoolIDs(topic)
	response := r.HGet(context.Background(), key, id)
	val, err := response.Result()
	if err != nil {
		if err == redis.Nil {
			errorstring := fmt.Sprintf("no redis entry for pool ID %s with topic %s \n", id, topic)
			return "", errors.New(errorstring)
		} else {
			return "", err
		}
	}
	return val, nil
	// fmt.Println("Args: ", response.Args())
	// b, err := response.Bool()
	// fmt.Println("bool, err: ", b, err)
	// res, err := response.Result()
	// fmt.Println("res, err: ", res, err)
	// fmt.Println("string: ", response.String())
	// if len(response.Val()) > 0 {
	// 	return response.Val(), nil
	// }

}

// SetPoolID sets a key value map for retrieval of parent trees of hashed pools.
// It is important to notice that this just facilitates the retrieval. The map can be reconstructed
// by id information stored in influx. Hence, the system does not rely on correct function/constant.
func SetPoolID(topic string, children []string, ID int64, r *redis.Client) error {
	log.Infof("Set pool IDs for %s: %v", topic, ID)
	poolMap := make(map[string]interface{})
	for _, num := range children {
		poolMap[num] = int(ID)
	}
	key := getKeyPoolIDs(topic)
	fmt.Printf("key, map: %s, %v \n", key, poolMap)
	resp := r.HSet(context.Background(), key, poolMap)
	_, err := resp.Result()
	if err != nil {
		return err
	}
	return nil
}

func getKeyPoolIDs(topic string) string {
	return "HashedPoolsMap_" + topic
}

func main() {

	r := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := SetPoolID("testtopic", []string{"111", "222", "333"}, 0, r)
	if err != nil {
		fmt.Println("error1: ", err)
	}

	parent, err := GetPoolsParentID("222", "testtopic", r)
	fmt.Println("parent, err: ", parent, err)

	// // unmarshalling from influx
	// ds, err := models.NewAuditStore()
	// if err != nil {
	// 	log.Error("NewAuditStore: ", err)
	// }

	// tree0, err := makeTestTree(2, []string{"hello ", "world."}, []string{"how ", "are ", "you?"})
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }

	// err = ds.SetStorageTreeInflux(*tree0, "test")
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }

	// time.Sleep(10 * time.Second)
	// vals, err := ds.GetStorageTreesInflux("test", time.Now().Add(-20*time.Second), time.Now())
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// log.Infof("got %v trees: ", len(vals))
	// if len(vals) > 0 {
	// 	var auxTree merkletree.MerkleTree
	// 	err = json.Unmarshal([]byte(vals[0][4].(string)), &auxTree)
	// 	if err != nil {
	// 		fmt.Println("error: ", err)
	// 	}
	// 	fmt.Println("tree: ", auxTree)
	// }

	// // marshalling and unmarshalling trees ----------

	// tree1, err := makeTestTree(2, []string{"hello ", "world."}, []string{"how ", "are ", "you?"})
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// fmt.Println("tree: ", tree1)

	// recoveredTree, err := marshAndUnmarshTree(*tree1)
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println("recovered tree: ", recoveredTree)

	// node1 := recoveredTree.Leafs[0]
	// storageBucket1 := node1.C.(*merkletree.StorageBucket)
	// data1, err := storageBucket1.ReadContent()

	// node2 := recoveredTree.Leafs[1]
	// storageBucket2 := node2.C.(*merkletree.StorageBucket)
	// data2, err := storageBucket2.ReadContent()

	// for _, word := range data1 {
	// 	fmt.Println(string(word))
	// }
	// for _, word := range data2 {
	// 	fmt.Println(string(word))
	// }

	// // second tree
	// tree2, err := makeTestTree(2, []string{"this ", "is "}, []string{"just  ", "another ", "tree."})
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }

	// forest, err := merkletree.ForestToTree([]merkletree.MerkleTree{*tree1, *tree2})
	// recoveredForest, err := marshAndUnmarshTree(*forest)
	// fmt.Println("recovered Forest: ", hex.EncodeToString(recoveredForest.Leafs[0].C.(*merkletree.ByteContent).Content))
	// fmt.Println("tree1 root: ", hex.EncodeToString(tree1.MerkleRoot))

	// // Marshaling and Unmarshaling mixed tree, i.e. with ByteContent and StorageBucket
	// mixedTree, err := merkletree.NewTree([]merkletree.Content{storageBucket1, merkletree.ByteContent{Content: tree2.MerkleRoot}})
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// fmt.Println("mixedTree leaf 0: ", hex.EncodeToString(mixedTree.Leafs[0].C.(*merkletree.StorageBucket).Content))
	// fmt.Println("mixedTree leaf 1: ", hex.EncodeToString(mixedTree.Leafs[1].C.(merkletree.ByteContent).Content))

	// recoveredMixedTree, err := marshAndUnmarshTree(*mixedTree)
	// fmt.Println("recoveredMixedTree leaf 0: ", hex.EncodeToString(recoveredMixedTree.Leafs[0].C.(*merkletree.StorageBucket).Content))
	// fmt.Println("recoveredMixedTree leaf 1: ", hex.EncodeToString(recoveredMixedTree.Leafs[1].C.(*merkletree.ByteContent).Content))

	// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------
	// p2 := merkletree.Bucket{}
	// fmt.Println("content: ", p2.Content)
	// n, err := p2.Content.Write([]byte("bye"))
	// fmt.Println("n, err: ", n, err)
	// fmt.Println(p2)

	// // Test verification ------------------------------------------------
	// dataType := flag.String("type", "hash-interestrates", "Type of data")
	// flag.Parse()

	// ds, err := models.NewAuditStore()
	// if err != nil {
	// 	log.Error("NewAuditStore: ", err)
	// }

	// poolNum := "0"
	// tree, err := ds.GetStorageTreeByID(*dataType, poolNum)
	// if err != nil {
	// 	log.Error(err)
	// }
	// num := 0
	// fmt.Printf("bucket type %T and value %v\n", tree.Leafs[num].C.(merkletree.StorageBucket).ID, tree.Leafs[num].C.(merkletree.StorageBucket).ID)
	// verif, _ := merklehashing.VerifyBucket(tree.Leafs[num].C.(merkletree.StorageBucket), ds)
	// fmt.Println("verification: ", verif)
	// newBucket := tree.Leafs[num].C.(merkletree.StorageBucket)
	// verifInterm, _ := merklehashing.VerifyBucket(newBucket, ds)
	// fmt.Println("verification before modification: ", verifInterm)
	// newBucket.Content[0] = 1
	// verifnew, _ := merklehashing.VerifyBucket(newBucket, ds)
	// fmt.Println("new verification: ", verifnew)

	// verif2, _ := merklehashing.VerifyPool(tree, *dataType, "0", ds)
	// fmt.Println("verification of pool: ", verif2)

	// level2Tree, err := ds.GetDailyTreeByID(*dataType, "2", "0")
	// if err != nil {
	// 	log.Error(err)
	// }
	// verif3, err := merklehashing.VerifyTree(level2Tree, "2", "0", ds)
	// fmt.Println(verif3, err)

	// level1Tree, err := ds.GetDailyTreeByID("", "1", "0")
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println("level1tree root: ", level1Tree.MerkleRoot)
	// verif4, err := merklehashing.VerifyTree(level1Tree, "1", "1", ds)
	// fmt.Println(verif4, err)

	// level0Tree, err := ds.GetDailyTreeByID("", "0", "0")
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println("leaf 1 duplicated: ", level0Tree.Leafs[0])
	// fmt.Println("leaf 2 duplicated: ", level0Tree.Leafs[1])
	// ---------------------------------------------------------------

	// dataType := flag.String("type", "hash-trades", "Type of data")
	// flag.Parse()
	// ds, err := models.NewAuditStore()
	// if err != nil {
	// 	log.Error("NewAuditStore: ", err)
	// }
	// // level0Tree, _ := ds.GetDailyTreeByID("", "0", "0")

	// timeFinal := time.Now().AddDate(0, 0, -1)
	// retval, err := ds.GetStorageTreeInflux(*dataType, timeFinal)
	// if err != nil {
	// 	log.Error("error getting merkle tree from influx: ", err)
	// }

	// bucket := retval.Root.Left.Left.C.(merkletree.StorageBucket)
	// data, err := bucket.ReadContent()
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(string(data[i]))
	// }

	// vals, err := ds.GetStorageTreesInflux(*dataType, timeInit, timeFinal)
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println("vals are: ", vals[0][1])
	// myTree := merkletree.MerkleTree{}
	// json.Unmarshal([]byte(vals[0][2].(string)), &myTree)

	// ds.SetDailyTreeInflux(myTree, "", "1")
	// ds.SetDailyTreeInflux(myTree, "", "1")
	// ds.SetDailyTreeInflux(myTree, "", "1")

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
