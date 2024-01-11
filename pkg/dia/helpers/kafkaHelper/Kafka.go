package kafkaHelper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	log "github.com/sirupsen/logrus"
)

const (
	messageSizeMax = 1e8
)

type KafkaMessage interface {
	MarshalBinary() ([]byte, error)
}

type KafkaMessageWithAHash interface {
	Hash() string
}

const (
	TopicIndexBlock = 0

	TopicFiltersBlock = 1
	TopicTrades       = 2
	TopicTradesBlock  = 3

	// The replica topics can be used to forward trades and blocks to other services in parallel.
	TopicFiltersBlockReplica = 4
	TopicTradesReplica       = 5
	TopicTradesBlockReplica  = 6

	TopicTradesEstimation = 7

	TopicFiltersBlockDone = 14

	TopicFiltersBlockTest = 21
	TopicTradesTest       = 22
	TopicTradesBlockTest  = 23
	TopicNFTTrades        = 24
	TopicNFTTradesTest    = 25

	retryDelay = 2 * time.Second
)

type Config struct {
	KafkaUrl []string
}

var (
	KafkaConfig Config
	topicSuffix string
)

func GetTopic(topic int) string {
	return getTopic(topic)
}

func getTopic(topic int) string {
	topicMap := map[int]string{
		1:  "filtersBlock",
		2:  "trades",
		3:  "tradesBlock",
		4:  "filtersBlockReplica" + topicSuffix,
		5:  "tradesReplica" + topicSuffix,
		6:  "tradesBlockReplica" + topicSuffix,
		7:  "tradesEstimation",
		14: "filtersblockHistoricalDone",
		21: "filtersblocktest",
		22: "tradestest",
		23: "tradesblocktest",
		24: "nfttrades",
		25: "nfttradestest",
	}
	result, ok := topicMap[topic]
	if !ok {
		log.Error("getTopic cant find topic", topic)
	}
	return result
}

func init() {
	KafkaConfig.KafkaUrl = []string{os.Getenv("KAFKAURL")}
	topicSuffix = utils.Getenv("KAFKA_TOPIC_SUFFIX", "")
}

// WithRetryOnError
func ReadOffset(topic int) (offset int64, err error) {
	for _, ip := range KafkaConfig.KafkaUrl {
		var conn *kafka.Conn
		conn, err = kafka.DialLeader(context.Background(), "tcp", ip, getTopic(topic), 0)
		if err != nil {
			log.Errorln("ReadOffset conn error: <", err, "> ", ip)
		} else {
			offset, err = conn.ReadLastOffset()
			if err != nil {
				log.Errorln("ReadOffset ReadLastOffset error: <", err, "> ")
			} else {
				return
			}
			defer func() {
				cerr := conn.Close()
				if err == nil {
					err = cerr
				}
			}()
		}
	}
	return
}

func ReadOffsetWithRetryOnError(topic int) (offset int64) {
	// TO DO: check double infinite for loops.
	for {
		for {
			for _, ip := range KafkaConfig.KafkaUrl {
				conn, err := kafka.DialLeader(context.Background(), "tcp", ip, getTopic(topic), 0)
				if err != nil {
					log.Errorln("ReadOffsetWithRetryOnError conn error: <", err, "> ", ip, " topic:", topic)
					time.Sleep(retryDelay)
				} else {
					defer func() {
						err = conn.Close()
						if err != nil {
							log.Error(err)
						}
					}()

					offset, err = conn.ReadLastOffset()
					if err != nil {
						log.Errorln("ReadOffsetWithRetryOnError ReadLastOffset error: <", err, "> ", ip, " topic:", topic)
						time.Sleep(retryDelay)
					} else {
						return offset
					}
				}
			}
		}
	}
}

func NewWriter(topic int) *kafka.Writer {
	kafkaTopic := getTopic(topic)
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  KafkaConfig.KafkaUrl,
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
		Async:    true,
	})
}

func NewSyncWriter(topic int) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:    KafkaConfig.KafkaUrl,
		Topic:      getTopic(topic),
		Balancer:   &kafka.LeastBytes{},
		Async:      false,
		BatchBytes: 1e9, // 1GB
	})
}

func NewSyncWriterWithCompression(topic int) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:          KafkaConfig.KafkaUrl,
		Topic:            getTopic(topic),
		Balancer:         &kafka.LeastBytes{},
		Async:            false,
		BatchBytes:       1e9, // 1GB
		CompressionCodec: &compress.GzipCodec,
	})
}

func NewReader(topic int) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   KafkaConfig.KafkaUrl,
		Topic:     getTopic(topic),
		Partition: 0,
		MinBytes:  0,
		MaxBytes:  10e6, // 10MB
	})
	return r
}

func WriteMessage(w *kafka.Writer, m KafkaMessage) error {
	key := []byte("helloKafka")
	value, err := m.MarshalBinary()
	if err == nil && value != nil {
		err = w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   key,
				Value: value,
			},
		)
		if err != nil {
			log.Errorln("WriteMessage error:", err, "sizeMessage:", float64(len(value))/(1024.0*1024.0), "MB")
		}
	} else {
		log.Errorln("Skipping write of message ", err, m)
	}
	return err
}

func NewReaderXElementsBeforeLastMessage(topic int, x int64) *kafka.Reader {

	var offset int64
	o, err := ReadOffset(topic)

	if err == nil && o-x > 0 {
		offset = o - x
	} else {
		log.Warningf("err %v on readOffset on topic %v", err, topic)
	}

	log.Println("NewReaderXElementsBeforeLastMessage: setting offset ", offset, "/", o)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   KafkaConfig.KafkaUrl,
		Topic:     getTopic(topic),
		Partition: 0,
		MinBytes:  0,
		MaxBytes:  10e6, // 10MB
	})
	err = r.SetOffset(offset)
	if err != nil {
		log.Error(err)
	}
	return r
}

func NewReaderNextMessage(topic int) *kafka.Reader {
	offset := ReadOffsetWithRetryOnError(topic)
	r := NewReader(topic)
	err := r.SetOffset(offset)
	if err != nil {
		log.Error(err)
	}
	log.Printf("Reading from offset %d/%d on topic %s", offset, offset, getTopic(topic))
	return r
}

func IsTopicEmpty(topic int) bool {
	log.Println("IsTopicEmpty: ", topic)
	offset := ReadOffsetWithRetryOnError(topic)
	offset--
	return offset < 0
}

func GetLastElementWithRetryOnError(topic int) interface{} {
	for {
		i, err := GetLastElement(topic)
		if err == nil {
			return i
		}
		time.Sleep(retryDelay)
		log.Println("GetLastElementWithRetryOnError retrying...")
	}
}

func GetLastElement(topic int) (interface{}, error) {
	offset := ReadOffsetWithRetryOnError(topic)
	offset--
	if offset < 0 {
		return nil, io.EOF
	} else {
		e, err := GetElements(topic, offset, 1)
		if err == nil {
			return e[0], nil
		} else {
			return nil, err
		}
	}
}

func GetElements(topic int, offset int64, nbElements int) ([]interface{}, error) {

	var result []interface{}

	var maxOffset = offset + int64(nbElements)

	conn, err := kafka.DialLeader(context.Background(), "tcp", KafkaConfig.KafkaUrl[0], getTopic(topic), 0)

	if err != nil {
		log.Errorln("kafka error:", err)
		return nil, err
	} else {

		newSeek, err := conn.Seek(int64(offset), kafka.SeekAbsolute)

		if err != nil {
			log.Errorln("kafka error on seek:", err)
			return nil, err
		}

		log.Printf("kafka newSeek:%v", newSeek)

		first, last, err := conn.ReadOffsets()
		if err != nil {
			return nil, err
		}
		log.Printf("kafka ReadOffsets:%v, %v", first, last)

		batch := conn.ReadBatch(0, messageSizeMax*nbElements)

		b := make([]byte, messageSizeMax)
		for c := offset; c <= maxOffset; c++ {
			z, err := batch.Read(b)
			if err != nil {
				log.Printf("error on batch read: %v", err)
				return nil, err
			}
			b2 := b[:z]

			switch topic {
			case TopicFiltersBlock:
				var e dia.FiltersBlock
				err = e.UnmarshalBinary(b2)
				if err == nil {
					result = append(result, e)
				}
			case TopicTrades:
				var e dia.Trade
				err = e.UnmarshalBinary(b2)
				if err == nil {
					result = append(result, e)
				}
			case TopicTradesBlock:
				var e dia.TradesBlock
				err = e.UnmarshalBinary(b2)
				if err == nil {
					result = append(result, e)
				}
			default:
				return nil, errors.New("missing case unknown topic in switch... function GetElements / Kafka.go")
			}

			if err != nil {
				errorMsg := fmt.Sprintf("parsing error while processing offset: %v/%v", c, maxOffset)
				return nil, errors.New(errorMsg)
			}
			if len(result) == nbElements {
				break
			}
		}

		return result, nil
	}
}
