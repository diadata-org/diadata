package merklehashing

import (
	"strconv"

	"github.com/cbergoon/merkletree"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// VerifyBucket checks whether a given storage bucket is corrupted at the first step, i.e.
// regarding its containing bucket pool
func VerifyBucket(sb merkletree.StorageBucket) (bool, error) {
	// TO DO
	// Get HashedPool/Tree by ID
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	tree, err := ds.GetMerkletreeInflux(sb.Topic, sb.Timestamp)
	if err != nil {
		return false, err
	}
	return tree.VerifyContent(sb)
}

// // TO DO: Integrate as case into VerifyTree
func VerifyPool(tree merkletree.MerkleTree, topic, ID string) (bool, error) {
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	parentTree, err := ds.GetDailyTreeByID(topic, "2", ID)
	if err != nil {
		return false, err
	}
	cont := merkletree.StorageBucket{Content: tree.MerkleRoot}
	return parentTree.VerifyContent(cont)
}

// VerifyTree verifies a tree in the hashing hierarchy with respect to the tree one level down
func VerifyTree(tree merkletree.MerkleTree, topic, level, ID string) (bool, error) {
	// TO DO: Add case if level=0
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	levelInt, err := strconv.ParseInt(level, 10, 64)
	if err != nil {
		return false, err
	}
	levelInt--
	levelBelow := strconv.Itoa(int(levelInt))
	// Get parent tree by ID
	parentTree, err := ds.GetDailyTreeByID(topic, levelBelow, ID)
	cont := merkletree.StorageBucket{Content: tree.MerkleRoot}
	return parentTree.VerifyContent(cont)
}
