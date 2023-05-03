package ipfs

import (
	"github.com/diadata-org/diadata/pkg/utils"
	icore "github.com/ipfs/boxo/coreiface"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"strings"
	"time"
)

type IPFSClient struct {
	coreApi        icore.CoreAPI
	node           *interface{}
	ctx
	DefaultTimeout time.Duration
}

func NewIPFSClient() *IPFSClient {
	ipfsApi, NodeApi := SpawnNode()
	timeout, err := strconv.Atoi(utils.Getenv("IPFS_DEFAULT_TIMEOUT", "600"))
	if err != nil {
		logrus.Fatalln("Failed creating IPFSClient ", err)
		return nil
	}
	return &IPFSClient{
		sh:             shell.NewShell(utils.Getenv("IPFS_ADDRESS", "127.0.0.1:5001")),
		DefaultTimeout: time.Duration(timeout),
	}
}

func (client *IPFSClient) GetContent(cid string) (string, error) {
	rc, err := client.sh.Cat(cid)
	if err != nil {
		return "", err
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {

		}
	}(rc)

	data, errRead := io.ReadAll(rc)
	if errRead != nil {
		return "", errRead
	}
	return string(data), nil
}

func (client *IPFSClient) AddContent(content string) (string, error) {
	cid, err := client.sh.Add(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	return cid, nil
}

func (client *IPFSClient) PublishContent(content string) (string, error) {
	cid, err := client.AddContent(content)
	if err != nil {
		return "", err
	}

	// implement additional logic for publishing the content using IPNS and pub-sub functionalities.

	return cid, nil
}
