package models

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// AuditStore is a datastore for the DIA audit trail
type AuditStore interface {
	FlushAuditBatch() error
	// Merkle Audit Trail methods
	SaveMerkletreeInflux(tree merkletree.MerkleTree, topic string) error
	GetMerkletreeInflux(topic string, timeInit, timeFinal time.Time) ([]merkletree.MerkleTree, error)
	GetMerkletreesInflux(topic string, timeInit, timeFinal time.Time) ([][]interface{}, error)
	SaveDailyTreeInflux(tree merkletree.MerkleTree, topic, level string, lastTimestamp time.Time) error
	GetLastTimestamp(topic, level string) (time.Time, error)
	GetLastID(topic, level string) (int64, error)
}

const (
	auditDBName         = "audit"
	influxDBTopicsTable = "hash"
	influxDBMerkleTable = "merkle"
)

// queryAuditDB convenience function to query the audit database
func queryAuditDB(clnt clientInfluxdb.Client, cmd string) (res []clientInfluxdb.Result, err error) {
	q := clientInfluxdb.Query{
		Command:  cmd,
		Database: auditDBName,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func NewAuditStore() (*DB, error) {
	return NewAuditStoreWithOptions(true, true)
}
func NewInfluxAuditStore() (*DB, error) {
	return NewAuditStoreWithOptions(false, true)
}

func NewRedisAuditStore() (*DB, error) {
	return NewAuditStoreWithOptions(true, false)
}

func NewAuditStoreWithoutInflux() (*DB, error) {
	return NewAuditStoreWithOptions(true, false)
}

func NewAuditStoreWithoutRedis() (*DB, error) {
	return NewAuditStoreWithOptions(false, true)
}

// NewAuditStoreWithOptions returns an audit store for either  influx or redis, depending
// on the boolean inputs
func NewAuditStoreWithOptions(withRedis bool, withInflux bool) (*DB, error) {
	var ci clientInfluxdb.Client
	var bp clientInfluxdb.BatchPoints
	var r *redis.Client
	var err error
	// This environment variable is either set in docker-compose or empty
	executionMode := os.Getenv("EXEC_MODE")
	address := ""

	if withRedis {
		// Run localhost for testing and server for production
		if executionMode == "production" {
			address = "redis:6379"
		} else {
			address = "localhost:6379"
		}
		r = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong2, err := r.Ping().Result()
		if err != nil {
			log.Error("NewAuditStore redis", err)
		}
		log.Debug("NewDB", pong2)
	}
	if withInflux {
		if executionMode == "production" {
			address = "http://influxdb:8086"
		} else {
			address = "http://localhost:8086"
		}
		ci, err = clientInfluxdb.NewHTTPClient(clientInfluxdb.HTTPConfig{
			Addr:     address,
			Username: "",
			Password: "",
		})
		if err != nil {
			log.Error("NewAuditStore influxdb", err)
		}
		bp, _ = createAuditBatchInflux()
		_, err = queryAuditDB(ci, fmt.Sprintf("CREATE DATABASE %s", auditDBName))
		if err != nil {
			log.Errorln("queryAuditDB CREATE DATABASE", err)
		}
	}
	return &DB{r, ci, bp, 0}, nil
}

func createAuditBatchInflux() (clientInfluxdb.BatchPoints, error) {
	bp, err := clientInfluxdb.NewBatchPoints(clientInfluxdb.BatchPointsConfig{
		Database:  auditDBName,
		Precision: "s",
	})
	if err != nil {
		log.Errorln("NewBatchPoints", err)
	}
	return bp, err
}

// FlushAuditBatch flushes a batch and writes it to influx
func (db *DB) FlushAuditBatch() error {
	var err error
	if db.influxBatchPoints != nil {
		err = db.WriteAuditBatchInflux()
	}
	return err
}

// WriteAuditBatchInflux writes a batch to influx
func (db *DB) WriteAuditBatchInflux() error {
	err := db.influxClient.Write(db.influxBatchPoints)
	if err != nil {
		log.Errorln("WriteBatchInflux", err)
		db.influxBatchPoints, _ = createAuditBatchInflux()
	} else {
		db.influxPointsInBatch = 0
	}
	return err
}

func (db *DB) addAuditPoint(pt *clientInfluxdb.Point) {
	db.influxBatchPoints.AddPoint(pt)
	db.influxPointsInBatch++
	if db.influxPointsInBatch >= influxMaxPointsInBatch {
		log.Debug("AddPoint forcing write Bash")
		db.WriteAuditBatchInflux()
	}
}

// ----------------------------------------------------------------------------------------
// Merkle Audit Trail Functionality
// ----------------------------------------------------------------------------------------

// HashingLayer activates a kafka writer content is written to.
// @topic is the category of hashed data in the merkle tree
// @content is a marshalled data point of the corresponding category
func HashingLayer(topic string, content []byte) error {
	config := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	}
	writer := kafka.NewWriter(config)
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte{},
			Value: content,
		},
	)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}

// Saving and retrieving from Topics Table (hashed data) ----------------------------------

// SaveMerkletreeInflux stores a tree from the merkletree package in Influx
func (db *DB) SaveMerkletreeInflux(tree merkletree.MerkleTree, topic string) error {

	// First marshal tree
	marshTree, err := json.Marshal(tree)
	if err != nil {
		log.Error(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"topic": topic}
	fields := map[string]interface{}{
		"value": string(marshTree),
	}

	pt, err := clientInfluxdb.NewPoint(influxDBTopicsTable, tags, fields, time.Now())
	if err != nil {
		log.Errorln("NewRateInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteAuditBatchInflux()
	if err != nil {
		log.Errorln("SaveRate: ", err)
	}
	log.Info("Batch written")
	return err
}

// GetMerkletreeInflux returns a slice of merkletrees of a given topic in a given time range
func (db *DB) GetMerkletreeInflux(topic string, timeInit, timeFinal time.Time) ([]merkletree.MerkleTree, error) {

	retval := merkletree.MerkleTree{}
	// q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s'", influxDBTreeTable, topic)
	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' and time > %d and time <= %d", influxDBTopicsTable, topic, timeInit.UnixNano(), timeFinal.UnixNano())
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return []merkletree.MerkleTree{}, err
	}
	// val is organized as: val[0]:(influx-)timestamp, val[1]:topic, val[2]:data
	val := res[0].Series[0].Values[0]
	err = json.Unmarshal([]byte(val[2].(string)), &retval)
	return []merkletree.MerkleTree{retval}, err
}

// GetMerkletreesInflux returns a slice of merkletrees of a given topic in a given time range.
// More precisely, the two-dimensional interface val is returned. It has length 3 and can be cast as follows:
// val[0]:(influx-)timestamp, val[1]:topic, val[2]:MerkleTree
func (db *DB) GetMerkletreesInflux(topic string, timeInit, timeFinal time.Time) (val [][]interface{}, err error) {

	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' and time > %d and time <= %d", influxDBTopicsTable, topic, timeInit.UnixNano(), timeFinal.UnixNano())
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return [][]interface{}{}, err
	}
	if len(res[0].Series) == 0 {
		return
	}
	val = res[0].Series[0].Values
	return
}

// Saving and retrieving from Merkle Table (hashed trees) ----------------------------------

// SaveDailyTreeInflux stores the trees which are produced on a daily basis in order to publish
// the master root hash.
// @topic only concerns level 2 and should be the empty string for level 1 and 0
// @level is an int corresponding to the level in the merkle documentation (currently 0<level<3)
// @lastTimestamp is the last timestamp of hashed trees from the data layer. Only applies to level 2
func (db *DB) SaveDailyTreeInflux(tree merkletree.MerkleTree, topic, level string, lastTimestamp time.Time) error {

	// Marshal tree
	marshTree, err := json.Marshal(tree)
	if err != nil {
		log.Error(err)
	}
	// Get last id and increment it
	lastID, err := db.GetLastID(topic, level)
	if err != nil {
		return err
	}
	id := strconv.FormatInt(lastID+1, 10)
	// Create a point and add to batch
	tags := map[string]string{"topic": topic, "level": level, "id": id, "lastTimestamp": strconv.Itoa(int(lastTimestamp.Unix()))}
	fields := map[string]interface{}{
		"value": string(marshTree),
	}
	pt, err := clientInfluxdb.NewPoint(influxDBMerkleTable, tags, fields, time.Now())
	if err != nil {
		log.Errorln("NewRateInflux:", err)
	} else {
		db.addPoint(pt)
	}
	err = db.WriteAuditBatchInflux()
	if err != nil {
		log.Errorln("SaveRate: ", err)
	}
	if topic == "" {
		log.Infof("Daily tree at level %s written \n", level)
	} else {
		log.Infof("Daily tree at level %s for topic %s written \n", level, topic)
	}
	return err
}

// GetDailyTreeByID returns the daily merkletree of a given topic, level and ID.
func (db *DB) GetDailyTreeByID(topic, level, ID string) (tree merkletree.MerkleTree, err error) {

	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' and level='%s' and id='%s'", influxDBMerkleTable, topic, level, ID)
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return merkletree.MerkleTree{}, err
	}
	if len(res[0].Series) == 0 {
		return
	}
	val := res[0].Series[0].Values[0]
	err = json.Unmarshal([]byte(val[2].(string)), &tree)
	return
}

// GetLastTimestamp retrieves the last timestamp for @topic (if given) and @level from the merkle table
func (db *DB) GetLastTimestamp(topic, level string) (time.Time, error) {

	q := fmt.Sprintf("SELECT lastTimestamp FROM (SELECT * FROM %s GROUP BY id) WHERE topic='%s' AND level='%s' ORDER BY DESC LIMIT 1", influxDBMerkleTable, topic, level)
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return time.Time{}, err
	}
	if len(res[0].Series) == 0 {
		// In this case, database is still empty, so set last timestamp to now-x
		return time.Now().AddDate(0, 0, -2), nil
	}
	val := res[0].Series[0].Values[0]
	i, err := strconv.ParseInt(val[1].(string), 10, 64)
	if err != nil {
		log.Error(err)
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

// GetLastID retrieves the highest current id for @topic (if given) and @level from the merkle table
func (db *DB) GetLastID(topic, level string) (int64, error) {

	q := fmt.Sprintf("SELECT id FROM (SELECT * FROM %s WHERE topic='%s' AND level='%s' GROUP BY id) ORDER BY DESC LIMIT 1", influxDBMerkleTable, topic, level)
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return 0, err
	}
	if len(res[0].Series) == 0 {
		// In this case, database is still empty, so begin with id=0
		return -1, nil
	}
	val := res[0].Series[0].Values[0]
	// We could also return last timestamp (in merkle table): val[0].(time.Time))
	lastID, err := strconv.ParseInt(val[1].(string), 10, 64)
	return lastID, err
}
