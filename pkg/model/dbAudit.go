package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// AuditStore is a datastore for the DIA audit trail
type AuditStore interface {
	FlushAuditBatch() error
	// Storage methods
	SetStorageTreeInflux(tree merkletree.MerkleTree, topic string) error
	GetStorageTreeInflux(topic string, timeLower time.Time) (merkletree.MerkleTree, error)
	GetStorageTreesInflux(topic string, timeInit, timeFinal time.Time) ([][]interface{}, error)
	GetStorageTreeByID(topic, ID string) (merkletree.MerkleTree, error)
	GetLastID(topic string) (string, error)

	// merkle tree methods
	SetDailyTreeInflux(tree merkletree.MerkleTree, topic, level string, children []string, lastTimestamp time.Time) error
	GetDailyTreesInflux(topic, level string, timeInit, timeFinal string) ([][]interface{}, error)
	GetDailyTreeByID(topic, level, ID string) (merkletree.MerkleTree, error)
	GetLastTimestamp(topic, level string) (time.Time, error)
	GetLastIDMerkle(topic, level string) (int64, error)
	SetPoolID(topic string, children []string, ID int64) error
	GetPoolsParentID(id, topic string) (string, error)
	GetYoungestChildMerkle(topic string) (int64, error)
}

const (
	// @auditDBName is the database for all data connected to the DIA audit trail
	auditDBName = "audit"
	// @influxDBStorageTable stores the structures (trees) which are continuously produced by the merkleService
	influxDBStorageTable = "storage"
	// @influxDBMerkleTable stores all trees which result from the hashing of trees from the storage table
	influxDBMerkleTable = "merkle"
)

func getKeyPoolIDs(topic string) string {
	return "HashedPoolsMap_" + topic
}

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
		Precision: "ns",
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

// HashingLayer activates a kafka writer to which content is written.
// @topic is the category of hashed data in the merkle tree. This list of contents can be
// 		  found in Kafka.go
// @content is a marshalled data point of the corresponding category
func HashingLayer(topic string, content []byte) error {
	// Determine topic index for kafka writer...
	topicNumber, err := kafkaHelper.GetTopicNumber(topic)
	if err != nil {
		return err
	}
	// ...and write into corresponding channel.
	writer := kafkaHelper.NewSyncWriter(topicNumber)
	err = writer.WriteMessages(context.Background(),
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

// Saving and retrieving from storage Table (hashed bucket pools) ------------------------------

// SetStorageTreeInflux stores a tree from the merkletree package in Influx.
// It is mainly used when flushing the bucket pools.
func (db *DB) SetStorageTreeInflux(tree merkletree.MerkleTree, topic string) error {

	// Set ID for buckets. IDs have the form i.j where i is the ID of the parent pool
	// and j is the ID of the bucket. IDs i of parent pools are (nanosecond Unix) times.
	influxTimeID := time.Now()
	var bucketsWithID []merkletree.Content
	firstDate := time.Now()
	lastDate := time.Time{}
	for i := range tree.Leafs {
		bucket := tree.Leafs[i].C.(merkletree.StorageBucket)
		bucket.ID = strconv.FormatInt(influxTimeID.UnixNano(), 10) + "." + strconv.FormatInt(int64(i), 10)
		bucketsWithID = append(bucketsWithID, bucket)
		// Get the time range of buckets for tags in order to uniquely define a storage tree.
		// This way, even if two storage trees are saved at the same time (go routines!), they differ in tags.
		if lastDate.Before(bucket.Timestamp) {
			lastDate = bucket.Timestamp
		}
		if bucket.Timestamp.Before(firstDate) {
			firstDate = bucket.Timestamp
		}
	}

	treeWithID, err := merkletree.NewTree(bucketsWithID)
	if err != nil {
		return err
	}

	// Marshal tree
	marshTree, err := json.Marshal(treeWithID)
	if err != nil {
		log.Error(err)
	}

	// Create a point and add to batch
	tags := map[string]string{
		"topic":     topic,
		"firstDate": strconv.FormatInt(firstDate.UnixNano(), 10),
		"lastDate":  strconv.FormatInt(lastDate.UnixNano(), 10),
	}
	fields := map[string]interface{}{
		"value": string(marshTree),
	}

	pt, err := clientInfluxdb.NewPoint(influxDBStorageTable, tags, fields, influxTimeID)
	if err != nil {
		log.Errorln("NewRateInflux:", err)
	} else {
		db.addAuditPoint(pt)
	}

	err = db.WriteAuditBatchInflux()
	if err != nil {
		log.Errorln("SaveRate: ", err)
	}
	log.Infof("Batch written for topic: %s\n", topic)
	return err
}

// GetStorageTreeInflux returns the first merkletree of a given topic with timestamp after timeLower
func (db *DB) GetStorageTreeInflux(topic string, timeLower time.Time) (merkletree.MerkleTree, error) {

	retval := merkletree.MerkleTree{}
	q := fmt.Sprintf("SELECT time, value FROM (SELECT * FROM %s WHERE topic='%s' and time >= %d) ORDER BY ASC LIMIT 1", influxDBStorageTable, topic, timeLower.UnixNano())
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return merkletree.MerkleTree{}, err
	}
	if len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		val := res[0].Series[0].Values[0]
		err = json.Unmarshal([]byte(val[1].(string)), &retval)
		return retval, err
	}
	return merkletree.MerkleTree{}, errors.New("empty response")
}

// GetStorageTreesInflux returns a slice of merkletrees from the storage table corresponding to a given topic in a given time range.
// More precisely, the two-dimensional interface val is returned. It has length 5 and can be cast as follows:
// val[0]:(influx-)timestamp, val[1]:firstDate, val[2]:lastDate, val[3]:topic, val[4]:Content/MerkleTree
func (db *DB) GetStorageTreesInflux(topic string, timeInit, timeFinal time.Time) (val [][]interface{}, err error) {

	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' and time > %d and time <= %d", influxDBStorageTable, topic, timeInit.UnixNano(), timeFinal.UnixNano())
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return [][]interface{}{}, err
	}
	if len(res[0].Series) == 0 {
		err = errors.New("empty response")
		return
	}
	val = res[0].Series[0].Values
	return
}

// GetStorageTreeByID returns a merkletree from the storage table with @ID and @topic
// We use primary key 'time' for storage trees.
func (db *DB) GetStorageTreeByID(topic, ID string) (merkletree.MerkleTree, error) {
	retval := merkletree.MerkleTree{}
	q := fmt.Sprintf("SELECT value FROM %s WHERE time=%s", influxDBStorageTable, ID)
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return merkletree.MerkleTree{}, err
	}
	if len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		val := res[0].Series[0].Values[0]
		err = json.Unmarshal([]byte(val[1].(string)), &retval)
		return retval, err
	}
	return merkletree.MerkleTree{}, errors.New("empty response")
}

// GetLastID retrieves the highest current id for @topic (if given) from the storage table
// as a string version of an int64 representing a unix nano time.
// Only used in DailyTreeTopic so not critical for scaling.
func (db *DB) GetLastID(topic string) (string, error) {

	// As ID in storage is identified with timestamp, we have the following query
	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' ORDER BY DESC LIMIT 1", influxDBStorageTable, topic)
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return "0", err
	}
	if len(res[0].Series) == 0 {
		// In this case, database is still empty, so begin with time.Now()
		return strconv.FormatInt(time.Now().UnixNano(), 10), nil
	}
	if len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		val := res[0].Series[0].Values[0]
		tstamp, _ := time.Parse(time.RFC3339, val[0].(string))
		return strconv.FormatInt(tstamp.UnixNano(), 10), nil
	}
	return "0", errors.New("empty response")
}

// Saving and retrieving from Merkle Table (hashed trees) ----------------------------------

// SetDailyTreeInflux stores the trees which are produced on a daily basis in order to publish
// the master root hash.
// @topic only concerns level 2 and should be the empty string for level 1 and 0
// @level is an int corresponding to the level in the merkle documentation (currently 0<level<3)
// @lastTimestamp is the last timestamp of hashed trees from the data layer. Only applies to level 2
func (db *DB) SetDailyTreeInflux(tree merkletree.MerkleTree, topic, level string, children []string, lastTimestamp time.Time) error {

	// Marshal tree
	marshTree, err := json.Marshal(tree)
	if err != nil {
		return err
	}

	// Get last id and increment it
	lastID, err := db.GetLastIDMerkle(topic, level)
	if err != nil {
		return err
	}
	id := strconv.FormatInt(lastID+1, 10)

	// Extend poolMap in Redis if level == 2
	if level == "2" {
		db.SetPoolID(topic, children, lastID+1)
	}
	// Encode children in order to store in influx
	childrenData, err := json.Marshal(children)
	if err != nil {
		return err
	}
	// Create a point and add to batch
	tags := map[string]string{
		"topic": topic,
		"level": level,
		"id":    id,
	}
	fields := map[string]interface{}{
		"value":         string(marshTree),
		"children":      string(childrenData),
		"lastTimestamp": strconv.Itoa(int(lastTimestamp.Unix())),
	}
	pt, err := clientInfluxdb.NewPoint(influxDBMerkleTable, tags, fields, time.Now())
	if err != nil {
		log.Errorln("NewRateInflux:", err)
	} else {
		db.addAuditPoint(pt)
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

// SetPoolID sets a key value map for retrieval of parent trees of hashed pools.
// It is important to notice that this just facilitates the retrieval. The map can be reconstructed
// by id information stored in influx. Hence, the system does not rely on correct function/constant
// connection of/to redis.
func (db *DB) SetPoolID(topic string, children []string, ID int64) error {
	log.Infof("Set pool IDs for %s: %v\n", topic, ID)
	poolMap := make(map[string]interface{})
	for _, num := range children {
		poolMap[num] = int(ID)
	}
	key := getKeyPoolIDs(topic)
	fmt.Printf("key, map: %s, %v \n", key, poolMap)
	// TO DO: Switch to HSet. atm HSet does not seem to be this:
	// https://github.com/go-redis/redis/blob/v8.1.3/commands.go#L1072
	// Check for go-redis version used here
	resp := db.redisClient.HMSet(key, poolMap)
	res, err := resp.Result()
	fmt.Println("response: ", res, err)
	return nil
}

// GetPoolsParentID returns the ID of level 2 tree such that hashed pool with @id is a leaf
func (db *DB) GetPoolsParentID(id, topic string) (string, error) {
	key := getKeyPoolIDs(topic)
	res := db.redisClient.HMGet(key, id)
	if res.Val()[0] != nil {
		return res.Val()[0].(string), nil
	}
	errorstring := fmt.Sprintf("no database entry for pool ID %s with topic %s \n", id, topic)
	return "", errors.New(errorstring)
}

// GetDailyTreesInflux returns a slice of merkletrees of a given topic in a given time range.
func (db *DB) GetDailyTreesInflux(topic, level string, timeInit, timeFinal string) (val [][]interface{}, err error) {

	q := fmt.Sprintf("SELECT * FROM %s WHERE topic='%s' and level='%s' and time > %s and time <= %s", influxDBMerkleTable, topic, level, timeInit, timeFinal)
	fmt.Println("influxquery: ", q)
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
	err = json.Unmarshal([]byte(val[6].(string)), &tree)
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
		return time.Now().AddDate(0, 0, -10), nil
	}
	val := res[0].Series[0].Values[0]
	i, err := strconv.ParseInt(val[1].(string), 10, 64)
	if err != nil {
		log.Error(err)
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

// GetLastIDMerkle retrieves the highest current id for @topic (if given) and @level from the merkle table
func (db *DB) GetLastIDMerkle(topic, level string) (int64, error) {

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
	lastID, err := strconv.ParseInt(val[1].(string), 10, 64)
	return lastID, err
}

// GetYoungestChildMerkle returns the highest ID from all pools hashed to level 2 trees.
// ID corresponds to a unix nano timestamp.
func (db *DB) GetYoungestChildMerkle(topic string) (int64, error) {
	// Get children from level 2 merkle tree with highest id
	q := fmt.Sprintf("SELECT children FROM (SELECT * FROM %s WHERE topic='%s' AND level='%s') ORDER BY DESC LIMIT 1", influxDBMerkleTable, topic, "2")
	res, err := queryAuditDB(db.influxClient, q)
	if err != nil {
		return 0, err
	}
	// Retrieve child with highest id (corresponding to youngest date)
	if len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		val := res[0].Series[0].Values[0][1].(string)
		childrenString := []string{}
		err = json.Unmarshal([]byte(val), &childrenString)
		if err != nil {
			log.Error(err)
			return 0, err
		}
		children, err := utils.StringsliceToInt(childrenString)
		if err != nil {
			return 0, err
		}
		youngestChild, err := utils.MaxIntSlice(children)
		if err != nil {
			return 0, err
		}
		return int64(youngestChild), nil
	}
	return 0, nil

}
