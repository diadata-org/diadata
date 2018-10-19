package kafkaHelper

import (
	"context"
	"errors"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"time"
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
	TopicFiltersBlock = 1
	TopicTrades       = 2
	TopicTradesBlock  = 3
	retryDelay        = 2 * time.Second
)

type Config struct {
	KafkaUrl []string
}

var KafkaConfig Config

func GetTopic(topic int) string {
	return getTopic(topic)
}

func getTopic(topic int) string {
	topicMap := map[int]string{
		1: "filtersBlock",
		2: "trades",
		3: "tradesBlock",
	}
	result, ok := topicMap[topic]
	if !ok {
		log.Error("getTopic cant fine topic", topic)
	}
	return result
}

func init() {
	KafkaConfig.KafkaUrl = []string{}
	var kafkaCount = 0
	for {
		kafkaName := fmt.Sprintf("kafka%d", kafkaCount)
		ip, err := net.LookupIP(kafkaName)
		if err != nil {
			log.Println("found", kafkaCount, "kafkas by name resolution")
			break
		} else {
			log.Println(kafkaName, " -> ", ip)
			kafkaName = fmt.Sprintf("kafka%d:9094", kafkaCount)
			KafkaConfig.KafkaUrl = append(KafkaConfig.KafkaUrl, kafkaName)
			kafkaCount++
		}
	}
	if kafkaCount == 0 {
		l := os.Getenv("LOCALHOST_KAFKA")
		if l == "" {
			log.Warning("could not find the kafka0 names, using kafka0")
			KafkaConfig.KafkaUrl = []string{"kafka0:9094"}
		} else {
			log.Println("LOCALHOST_KAFKA is set, Adding localhost, probably runned outside of kafka")
			KafkaConfig.KafkaUrl = []string{"localhost:9094"}
		}
	}
	log.Printf("brokers: %v", KafkaConfig.KafkaUrl)
}

// WithRetryOnError
func ReadOffset(topic int) (int64, error) {
	var err error
	for _, ip := range KafkaConfig.KafkaUrl {
		conn, err := kafka.DialLeader(context.Background(), "tcp", ip, getTopic(topic), 0)
		if err != nil {
			log.Errorln("ReadOffset conn error: <", err, "> ", ip)
		} else {
			defer conn.Close()
			offset, err := conn.ReadLastOffset()
			if err != nil {
				log.Errorln("ReadOffset ReadLastOffset error: <", err, "> ")
			} else {
				return offset, nil
			}
		}
	}
	return 0, err
}

func ReadOffsetWithRetryOnError(topic int) (offset int64) {
	for {
		for {
			for _, ip := range KafkaConfig.KafkaUrl {
				conn, err := kafka.DialLeader(context.Background(), "tcp", ip, getTopic(topic), 0)
				if err != nil {
					log.Errorln("ReadOffsetWithRetryOnError conn error: <", err, "> ", ip, " topic:", topic)
					time.Sleep(retryDelay)
				} else {
					defer conn.Close()
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
		log.Println("ReadOffsetWithRetryOnError retrying...")
	}
}

func NewWriter(topic int) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  KafkaConfig.KafkaUrl,
		Topic:    getTopic(topic),
		Balancer: &kafka.LeastBytes{},
		Async:    true,
	})
}

func NewSyncWriter(topic int) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  KafkaConfig.KafkaUrl,
		Topic:    getTopic(topic),
		Balancer: &kafka.LeastBytes{},
		Async:    false,
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
		err := w.WriteMessages(context.Background(),
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
		log.Warning("err %v on readOffset on topic %v", err, topic)
	}

	log.Println("NewReaderXElementsBeforeLastMessage: setting offset ", offset, "/", o)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   KafkaConfig.KafkaUrl,
		Topic:     getTopic(topic),
		Partition: 0,
		MinBytes:  0,
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(offset)
	return r
}

func NewReaderNextMessage(topic int) *kafka.Reader {
	offset := ReadOffsetWithRetryOnError(topic)
	r := NewReader(topic)
	r.SetOffset(offset)
	log.Printf("Reading from offset %d/%d on topic %s", offset, offset, getTopic(topic))
	return r
}

func IsTopicEmpty(topic int) bool {
	log.Println("IsTopicEmpty: ", topic)
	offset := ReadOffsetWithRetryOnError(topic)
	offset--
	if offset < 0 {
		return true
	}
	return false
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

			err = nil

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
				return nil, errors.New("Missing case unknown topic in switch... function GetElements / Kafka.go")
			}

			if err != nil {
				errorMsg := fmt.Sprintf("Parsing error while processing offset: %v/%v", c, maxOffset)
				return nil, errors.New(errorMsg)
			}
			if len(result) == nbElements {
				break
			}
		}

		return result, nil
	}
}
