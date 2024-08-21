package hydrationhelper

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type AssetId struct {
	Key   string
	Value uint32
}

type PalletAssetRegistryAssetType int

const (
	Token PalletAssetRegistryAssetType = iota
	XYK
	StableSwap
	Bond
	External
)

const (
	DefaultRefreshDelay              = 400 // millisec
	DefaultSleepBetweenContractCalls = 300 // millisec
	DefaultEventsLimit               = 100
	DefaultSwapContractsLimit        = 100
)

func (a PalletAssetRegistryAssetType) String() string {
	switch a {
	case Token:
		return "Token"
	case XYK:
		return "XYK"
	case StableSwap:
		return "StableSwap"
	case Bond:
		return "Bond"
	case External:
		return "External"
	default:
		return "Unknown"
	}
}

type PalletAssetRegistryAssetDetails struct {
	Name               types.OptionBytes
	AssetType          types.U8
	ExistentialDeposit types.U128
	Symbol             types.OptionBytes
	Decimals           types.OptionU8
}
