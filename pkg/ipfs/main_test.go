package ipfs_test

import (
	"github.com/diadata-org/diadata/pkg/ipfs"
	"testing"
)

func TestIPFSIntegration(t *testing.T) {
	client := ipfs.NewIPFSClient()

	t.Run("Test Case 1: Verify that the library can connect to an IPFS node and retrieve content.", func(t *testing.T) {
		content, err := client.GetContent("QmZTR5bcpQD7cFgTorqxZDYaew1Wqgfbd2ud9QqGPAkK2V") // This is the CID of a well-known IPFS test file.
		if err != nil {
			t.Errorf("Failed to connect to IPFS node and retrieve content: %v", err)
		}
		if content == "" {
			t.Error("Content is empty")
		}
	})

	t.Run("Test Case 2: Verify that the library can add content to IPFS and retrieve the corresponding content address.", func(t *testing.T) {
		testContent := "Hello, IPFS!"
		cid, err := client.AddContent(testContent)
		if err != nil {
			t.Errorf("Failed to add content to IPFS: %v", err)
		}
		if cid == "" {
			t.Error("Content address is empty")
		}
	})

	t.Run("Test Case 3: Verify that the library can publish content to IPFS and retrieve the corresponding IPFS hash.", func(t *testing.T) {
		testContent := "Hello, IPFS again!"
		ipfsHash, err := client.PublishContent(testContent)
		if err != nil {
			t.Errorf("Failed to publish content to IPFS: %v", err)
		}
		if ipfsHash == "" {
			t.Error("IPFS hash is empty")
		}
	})
}
