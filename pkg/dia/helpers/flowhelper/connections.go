package flowhelper

import (
	"errors"
	"fmt"

	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	FlowAPI10      = "access-001.mainnet10.nodes.onflow.org:9000"
	FlowAPI11      = "access-001.mainnet11.nodes.onflow.org:9000"
	FlowAPI12      = "access-001.mainnet12.nodes.onflow.org:9000"
	FlowAPI13      = "access-001.mainnet13.nodes.onflow.org:9000"
	FlowAPI14      = "access-001.mainnet14.nodes.onflow.org:9000"
	FlowAPI15      = "access-001.mainnet15.nodes.onflow.org:9000"
	FlowAPI16      = "access-001.mainnet16.nodes.onflow.org:9000"
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
	RootHeight10      = uint64(15791891)
	RootHeight11      = uint64(16755602)
	RootHeight12      = uint64(17544523)
	RootHeight13      = uint64(18587478)
	RootHeight14      = uint64(19050753)
	RootHeight15      = uint64(21291692)
	RootHeight16      = uint64(23830813)
	RootHeightCurrent = uint64(27341470)
	RootHeights       = []uint64{
		RootHeight1,
		RootHeight2,
		RootHeight3,
		RootHeight4,
		RootHeight5,
		RootHeight6,
		RootHeight7,
		RootHeight8,
		RootHeight9,
		RootHeight10,
		RootHeight11,
		RootHeight12,
		RootHeight13,
		RootHeight14,
		RootHeight15,
		RootHeight16,
		RootHeightCurrent,
	}
)

// GetFlowClient returns a feasible client corresponding to the block's startheight.
// https://docs.onflow.org/node-operation/past-sporks/
func GetFlowClient(startheight uint64) (*client.Client, error) {
	if startheight >= RootHeightCurrent {
		fmt.Printf("make flow client at current level with: %s\n", FlowAPICurrent)
		return client.New(FlowAPICurrent, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight16 {
		return client.New(FlowAPI16, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight15 {
		return client.New(FlowAPI15, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight14 {
		return client.New(FlowAPI14, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight13 {
		return client.New(FlowAPI13, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight12 {
		return client.New(FlowAPI12, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight11 {
		return client.New(FlowAPI11, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight10 {
		return client.New(FlowAPI10, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight9 {
		return client.New(FlowAPI9, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight8 {
		return client.New(FlowAPI8, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight7 {
		return client.New(FlowAPI7, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight6 {
		return client.New(FlowAPI6, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight5 {
		return client.New(FlowAPI5, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight4 {
		return client.New(FlowAPI4, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight3 {
		return client.New(FlowAPI3, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight2 {
		return client.New(FlowAPI2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else if startheight >= RootHeight1 {
		return client.New(FlowAPI1, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return nil, errors.New("startheight too small. No client available.")
}
