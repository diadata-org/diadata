package merklehashing

import (
	"strconv"
	"strings"

	"github.com/cbergoon/merkletree"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// VerifyBucket checks whether a given storage bucket is corrupted at the first step, i.e.
// regarding its containing bucket pool
func VerifyBucket(sb merkletree.StorageBucket) (bool, error) {

	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}
	// Get ID of parent pool
	id := strings.Split(sb.ID, ".")[0]
	// Get tree corresponding to the pool
	tree, err := ds.GetMerkletreeByID(sb.Topic, id)
	if err != nil {
		return false, err
	}
	return tree.VerifyContent(sb)
}

// VerifyPool verifies a topic tree as content of a daily tree
func VerifyPool(tree merkletree.MerkleTree, topic, ID string) (bool, error) {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}

	// Get ID of pool's parent tree
	parentID, err := ds.GetPoolsParentID(ID, topic)
	if err != nil {
		return false, err
	}

	parentTree, err := ds.GetDailyTreeByID(topic, "2", parentID)
	if err != nil {
		return false, err
	}
	cont := merkletree.StorageBucket{Content: tree.MerkleRoot}
	return parentTree.VerifyContent(cont)
}

// VerifyTree verifies a tree in the hashing hierarchy with respect to the tree one level down
func VerifyTree(tree merkletree.MerkleTree, level, ID string) (bool, error) {
	// TO DO: Add case if level=0
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}
	levelInt, err := strconv.ParseInt(level, 10, 64)
	if err != nil {
		return false, err
	}
	levelInt--
	levelBelow := strconv.Itoa(int(levelInt))
	// Get parent tree by ID
	parentTree, err := ds.GetDailyTreeByID("", levelBelow, ID)
	cont := merkletree.StorageBucket{Content: tree.MerkleRoot}
	return parentTree.VerifyContent(cont)
}

// VerifyContent checks whether a content/bucket is uncorrupted all the way
// up to the merkle root.
// Alternatively, we could make the above methods return the containing tree.
// This would shorten the below code
func VerifyContent(sb merkletree.StorageBucket) (bool, error) {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}

	// Verify bucket in pool
	// Get ID of parent pool
	id := strings.Split(sb.ID, ".")[0]
	// Get tree corresponding to the pool
	tree, err := ds.GetMerkletreeByID(sb.Topic, id)
	if err != nil {
		return false, err
	}
	// Verify bucket in pool
	val, err := tree.VerifyContent(sb)
	if err != nil {
		return false, err
	}
	if val == false {
		return false, nil
	}

	// Verify pool in tree level 2
	// Get ID of pool's parent tree (all trees with lower level have the same ID by construction)
	parentID, err := ds.GetPoolsParentID(id, sb.Topic)
	if err != nil {
		return false, err
	}
	// Get parent tree (level 2)
	level2Tree, err := ds.GetDailyTreeByID(sb.Topic, "2", parentID)
	if err != nil {
		return false, err
	}
	// Verify root hash of pool in level 2 tree
	cont := merkletree.StorageBucket{Content: tree.MerkleRoot}
	val, err = level2Tree.VerifyContent(cont)
	if err != nil {
		return false, err
	}
	if val == false {
		return false, nil
	}

	// Verify tree level 2 in tree level 1
	val, err = VerifyTree(level2Tree, "2", parentID)
	if err != nil {
		return false, err
	}
	if val == false {
		return false, nil
	}

	// Verify tree level 1 in tree level 0
	level1Tree, err := ds.GetDailyTreeByID("", "1", parentID)
	if err != nil {
		return false, err
	}
	val, err = VerifyTree(level1Tree, "1", parentID)
	if err != nil {
		return false, err
	}
	if val == false {
		return false, nil
	}

	return true, nil
}
