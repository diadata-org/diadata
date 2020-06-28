package models

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"time"

	"github.com/cbergoon/merkletree"
)

// Bucket implements Content interface from merkle_tree package.
// @Content is a byte slice of fixed size per @Type
// @ID is a unique identification string
// @Type designates the category of data (for instance interest rate or trade)
// @HashRate sets the frequency of hashing the data of the corresponding type
// @Timestamp is the (Unix-?)time, the container is hashed
type Bucket struct {
	Content *bytes.Buffer
	// properties of the bucket
	size     int
	Topic    string
	HashRate time.Duration
	// values possibly assigned to the bucket
	used      bool
	ID        string
	Timestamp time.Time
}

// BucketPool implements a leaky pool of Buckets in the form of a bounded channel.
type BucketPool struct {
	c     chan Bucket
	width int
	Topic string
}

// Forest is a collection of merkletrees. It represents the collection
// of merkle trees in one category during a day.
type Forest struct {
	Trees []*merkletree.MerkleTree
	Topic string
	Day   time.Time
}

// NewBucket creates a new bucket of size @size.
// TO DO: Extend to Type of Bucket (and pool below)
func NewBucket(size int, topic string) (b *Bucket) {
	return &Bucket{
		Content: bytes.NewBuffer(make([]byte, 0, size)),
		size:    size,
		Topic:   topic,
	}
}

// NewBucketPool creates a new BucketPool bounded to the given maxSize.
// It is initialized with empty Buckets sized based on width.
func NewBucketPool(maxNum int, size int, topic string) (bp *BucketPool) {
	bp = &BucketPool{
		c:     make(chan Bucket, maxNum),
		width: size,
		Topic: topic,
	}
	// Fill channel with empty buckets
	for i := 0; i < maxNum; i++ {
		bucket := NewBucket(size, topic)
		bp.c <- *bucket
	}
	return
}

// Size returns the size of a bucket
func (b *Bucket) Size() int {
	return b.size
}

// Used returns true if the bucket was written to
func (b *Bucket) Used() bool {
	return b.used
}

// Len returns the numbers of elements in the bucket pool
func (bp *BucketPool) Len() int {
	return len(bp.c)
}

// Get gets a Bucket from the BucketPool, or creates a new one if none are
// available in the pool.
func (bp *BucketPool) Get() (b Bucket, err error) {
	select {
	case b = <-bp.c:
		// Get bucket from pool

		if b.used {
			// In this case, all buckets from the pool have been used and a new pool
			// should be created
			bp.c <- b
			return *NewBucket(bp.width, bp.Topic), errors.New("size error. pool is exhausted")
		}
	default:
		return *NewBucket(bp.width, bp.Topic), errors.New("size error. pool is exhausted")
		// fmt.Println("make new bucket")
		// b = *NewBucket(bp.width)
	}
	return
}

// Put returns the given Bucket to the BucketPool.
func (bp *BucketPool) Put(b Bucket) bool {
	// Check whether Bucket is of the right kind. If not, reject.
	if bp.Topic != b.Topic {
		fmt.Println("error with topics")
		return false
	}

	select {
	case bp.c <- b:
		// Bucket went back into pool.
		return true
	default:
		// Bucket didn't go back into pool, just discard.
		return false
	}
}

// WriteContent appends a byte array to a bucket if there is enough
// space. Does not write and returns false if there isn't.
func (b *Bucket) WriteContent(bs []byte) bool {
	if b.Content.Len()+len(bs) > b.Size() {
		return false
	}
	b.Content.Write(bs)
	b.used = true
	return true
}

// AddTree adds a tree to a forest
// TO DO: Check for the right type?
func (f Forest) AddTree(t *merkletree.MerkleTree) Forest {
	f.Trees = append(f.Trees, t)
	return f
}

// Size returns the number of trees in the forest @f.
func (f Forest) Size() int {
	return len(f.Trees)
}

// CalculateHash calculates the hash of a bucket. Is needed for a bucket in
// order to implement Content from merkle_tree.
func (b Bucket) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(b.Content.Bytes()); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Equals is true if buckets are identical. Is needed for a bucket in
// order to implement Content from merkle_tree.
func (b Bucket) Equals(other merkletree.Content) (bool, error) {
	// Extend for other fields, but which? Do we need all fields?
	if !EqualBytes(b.Content.Bytes(), other.(Bucket).Content.Bytes()) {
		return false, nil
	}
	if b.size != other.(Bucket).size {
		return false, nil
	}
	if b.ID != other.(Bucket).ID {
		return false, nil
	}
	if b.Topic != other.(Bucket).Topic {
		return false, nil
	}
	if b.HashRate != other.(Bucket).HashRate {
		return false, nil
	}
	return true, nil
}

// func (b Bucket) Equals(other Bucket) (bool, error) {
// 	// Extend for other fields, but which? Do we need all fields?
// 	if !EqualBytes(b.Content.Bytes(), (other.Content).Bytes()) {
// 		return false, nil
// 	}
// 	if b.size != other.size {
// 		return false, nil
// 	}
// 	if b.ID != other.ID {
// 		return false, nil
// 	}
// 	if b.Type != other.Type {
// 		return false, nil
// 	}
// 	if b.HashRate != other.HashRate {
// 		return false, nil
// 	}
// 	return true, nil
// }

// MakeTree returns a Merkle tree built from the Buckets in the pool @bp
func MakeTree(bp *BucketPool) (*merkletree.MerkleTree, error) {
	leafs := []merkletree.Content{}
	L := bp.Len()
	for i := 0; i < L; i++ {
		leafs = append(leafs, <-bp.c)
	}
	t, err := merkletree.NewTree(leafs)
	return t, err
}

// EqualBytes compares two byte slices. Should be put into some helper package.
func EqualBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// -------------------------------------------------------------------------
// Additional merkle tree methods/functions which are not in the package yet
// -------------------------------------------------------------------------

// NumNodes computes the number of nodes in the tree given by the root node @node.
// Leafs are not counted.
func NumNodes(node *merkletree.Node) int {
	count := 1
	if node.Left.C == nil {
		count += NumNodes(node.Left)
	}
	if node.Right.C == nil {
		count += NumNodes(node.Right)
	}
	return count
}
