package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

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

	log.Info("TopicSwitch is built and triggered")
	go kc.StartKafkaReader(topic)
	return kc
}

func fillBucketPool(bp *models.BucketPool, topic string, topicChan chan *kafka.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	ok := true
	bucket, err := bp.Get()
	if err != nil {
		log.Error("error with Get bucket: ", err)
	}

	for {
		message := <-topicChan
		if ok {
			ok = bucket.WriteContent(message.Value)
			// fmt.Println("bucket content: ", bucket.Content)
		} else {
			fmt.Println("bucket full. Return bucket to pool.")
			bp.Put(bucket)
			bucket, err = bp.Get()
			if err != nil {
				fmt.Println(err)
				return
				// TO DO: Hash the bucket pool and make a new one
			}
			fmt.Println("new bucket's content: ", bucket.Content)
			ok = bucket.WriteContent(message.Value)
		}
	}
}

func main() {

	// preliminary main
	// One instance of main for each data type
	dataType := flag.String("type", "mytopic", "Type of data")
	flag.Parse()

	kc := ActivateKafkaChannel(*dataType)
	defer kc.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	bp := models.NewBucketPool(32, 2560)
	go fillBucketPool(bp, *dataType, kc.chanMessage, &wg)
	fmt.Println(bp)
	defer wg.Wait()

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
