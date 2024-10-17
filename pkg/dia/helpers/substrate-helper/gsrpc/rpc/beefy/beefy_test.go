package beefy

import (
	"os"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/client"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/config"
)

var testBeefy Beefy

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testBeefy = NewBeefy(cl)
	os.Exit(m.Run())
}
