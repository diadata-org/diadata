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
	Type     string
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
}

// NewBucket creates a new bucket of size @size.
func NewBucket(size int) (b *Bucket) {
	return &Bucket{
		Content: bytes.NewBuffer(make([]byte, 0, size)),
		size:    size,
	}
}

// NewBucketPool creates a new BytePool bounded to the given maxSize.
// It is initialized with empty Buckets sized based on width.
func NewBucketPool(maxNum int, size int) (bp *BucketPool) {
	bp = &BucketPool{
		c:     make(chan Bucket, maxNum),
		width: size,
	}
	// Fill channel with empty buckets
	for i := 0; i < maxNum; i++ {
		bucket := NewBucket(size)
		bp.c <- *bucket
	}
	return
}

// Size returns the size of a bucket
func (b *Bucket) Size() int {
	return b.size
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
		fmt.Println("get bucket from pool, if not used yet")
		if b.used {
			// In this case, all buckets from the pool have been used and a new pool
			// should be created
			bp.c <- b
			return *NewBucket(bp.width), errors.New("size error. pool is exhausted")
		}
	default:
		fmt.Println("make new bucket")
		b = *NewBucket(bp.width)
	}
	return
}

// Put returns the given Bucket to the BucketPool.
func (bp *BucketPool) Put(b Bucket) bool {
	// Check whether Bucket is of the right kind. If not, reject.

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
// func (b Bucket) Equals(other merkletree.Content) (bool, error) {
// 	// Extend for other fields, but which? Do we need all fields?
// 	if !EqualBytes(b.Content.Bytes(), other.(Bucket).Content.Bytes()) {
// 		return false, nil
// 	}
// 	if b.size != other.(Bucket).size {
// 		return false, nil
// 	}
// 	if b.ID != other.(Bucket).ID {
// 		return false, nil
// 	}
// 	if b.Type != other.(Bucket).Type {
// 		return false, nil
// 	}
// 	if b.HashRate != other.(Bucket).HashRate {
// 		return false, nil
// 	}
// 	return true, nil
// }
func (b Bucket) Equals(other Bucket) (bool, error) {
	// Extend for other fields, but which? Do we need all fields?
	if !EqualBytes(b.Content.Bytes(), (other.Content).Bytes()) {
		return false, nil
	}
	if b.size != other.size {
		return false, nil
	}
	if b.ID != other.ID {
		return false, nil
	}
	if b.Type != other.Type {
		return false, nil
	}
	if b.HashRate != other.HashRate {
		return false, nil
	}
	return true, nil
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

// numNodes computes the number of nodes in the tree given by the root node @node.
func numNodes(node *merkletree.Node) int {
	count := 1
	// fmt.Println("content of left node: ", node.Left.C)
	if node.Left.C == nil {
		count += numNodes(node.Left)
	}
	if node.Right.C == nil {
		count += numNodes(node.Right)
	}
	return count
}
