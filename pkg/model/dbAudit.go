package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

// AuditStore is a datastore for the DIA audit trail
type AuditStore interface {
	FlushAuditBatch() error
	// Merkle Audit Trail methods
	SaveMerkletreeInflux(tree merkletree.MerkleTree, topic string) error
	GetMerkletreeInflux(topic string, timeInit, timeFinal time.Time) (merkletree.MerkleTree, error)
}

const (
	auditDBName       = "audit"
	influxDBTreeTable = "merkle"
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

func (db *DB) FlushAuditBatch() error {
	var err error
	if db.influxBatchPoints != nil {
		err = db.WriteAuditBatchInflux()
	}
	return err
}

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

// SaveMerkletreeInflux stores a tree from the merkletree package in Influx
func (db *DB) SaveMerkletreeInflux(tree merkletree.MerkleTree, topic string) error {

	// First marshal tree
	marshTree, err := json.Marshal(tree)
	if err != nil {
		log.Error(err)
	}

	// Check whether unmarshalling tree works:
	retrievedTree := merkletree.MerkleTree{}
	err = json.Unmarshal(marshTree, &retrievedTree)
	if err != nil {
		fmt.Println("error retrieving tree: ", err)
	}
	fmt.Println("retrieved tree: ", retrievedTree)

	// Create a point and add to batch
	tags := map[string]string{"topic": topic}
	fields := map[string]interface{}{
		"value": string(marshTree),
	}

	pt, err := clientInfluxdb.NewPoint(influxDBTreeTable, tags, fields, time.Now())
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
func (db *DB) GetMerkletreeInflux(topic string, timeInit, timeFinal time.Time) (merkletree.MerkleTree, error) {

	retval := merkletree.MerkleTree{}

	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s'", influxDBTreeTable, topic)
	// q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' and time > %d and time < %d", influxDBTreeTable, topic, timeInit.Unix(), timeFinal.Unix())
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	val := res[0].Series[0].Values[0]
	err = json.Unmarshal([]byte(val[2].(string)), &retval)
	return retval, err
}
