package flowhelper

import (
	"errors"
	"fmt"

	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

const (
	FlowAPI1       = "access-001.mainnet1.nodes.onflow.org:9000"
	FlowAPI2       = "access-001.mainnet2.nodes.onflow.org:9000"
	FlowAPI3       = "access-001.mainnet3.nodes.onflow.org:9000"
	FlowAPI4       = "access-001.mainnet4.nodes.onflow.org:9000"
	FlowAPI5       = "access-001.mainnet5.nodes.onflow.org:9000"
	FlowAPI6       = "access-001.mainnet6.nodes.onflow.org:9000"
	FlowAPI7       = "access-001.mainnet7.nodes.onflow.org:9000"
	FlowAPI8       = "access-001.mainnet8.nodes.onflow.org:9000"
	FlowAPI9       = "access-001.mainnet9.nodes.onflow.org:9000"
	FlowAPICurrent = "access.mainnet.nodes.onflow.org:9000"
	RequestLimit   = uint64(249)
)

var (
	RootHeight1       = uint64(7601063)
	RootHeight2       = uint64(8742959)
	RootHeight3       = uint64(9737133)
	RootHeight4       = uint64(9992020)
	RootHeight5       = uint64(12020337)
	RootHeight6       = uint64(12609237)
	RootHeight7       = uint64(13404174)
	RootHeight8       = uint64(13950742)
	RootHeight9       = uint64(14892104)
	RootHeightCurrent = uint64(15791891)
	RootHeights       = []uint64{RootHeight1, RootHeight2, RootHeight3, RootHeight4, RootHeight5, RootHeight6, RootHeight7, RootHeight8, RootHeight9, RootHeightCurrent}
)

// GetFlowClient returns a feasible client corresponding to the block's startheight.
func GetFlowClient(startheight uint64) (*client.Client, error) {
	if startheight >= RootHeightCurrent {
		fmt.Printf("make flow client at current level with: %s\n", FlowAPICurrent)
		return client.New(FlowAPICurrent, grpc.WithInsecure())
	} else if startheight >= RootHeight9 {
		return client.New(FlowAPI9, grpc.WithInsecure())
	} else if startheight >= RootHeight8 {
		return client.New(FlowAPI8, grpc.WithInsecure())
	} else if startheight >= RootHeight7 {
		return client.New(FlowAPI7, grpc.WithInsecure())
	} else if startheight >= RootHeight6 {
		return client.New(FlowAPI6, grpc.WithInsecure())
	} else if startheight >= RootHeight5 {
		return client.New(FlowAPI5, grpc.WithInsecure())
	} else if startheight >= RootHeight4 {
		return client.New(FlowAPI4, grpc.WithInsecure())
	} else if startheight >= RootHeight3 {
		return client.New(FlowAPI3, grpc.WithInsecure())
	} else if startheight >= RootHeight2 {
		return client.New(FlowAPI2, grpc.WithInsecure())
	} else if startheight >= RootHeight1 {
		return client.New(FlowAPI1, grpc.WithInsecure())
	}
	return nil, errors.New("startheight too small. No client available.")
}
