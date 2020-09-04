package merklehashing

import (
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
