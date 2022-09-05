package acalaswap

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// Adapted from: https://github.com/AcalaNetwork/acala.js/blob/master/packages/type-definitions/src/json/types.json
var tokenSymbolMap = map[string]int{
	"ACA":    0,
	"AUSD":   1,
	"DOT":    2,
	"LDOT":   3,
	"RENBTC": 20,
	"CASH":   21,
	"KAR":    128,
	"KUSD":   129,
	"KSM":    130,
	"LKSM":   131,
	"BNC":    168,
	"VSKSM":  169,
	"PHA":    170,
	"KINT":   171,
	"KBTC":   172,
}
var tokenSymbolMapReverse map[int]string

func init() {
	for k, v := range tokenSymbolMap {
		tokenSymbolMapReverse[v] = k
	}
}

type TokenSymbol struct {
	ACA    int
	AUSD   int
	DOT    int
	LDOT   int
	RENBTC int
	CASH   int
	KAR    int
	KUSD   int
	KSM    int
	LKSM   int
	BNC    int
	VSKSM  int
	PHA    int
	KINT   int
	KBTC   int
}
type EvmAddress types.H160

type DexShare struct {
	Token TokenSymbol
	Erc20 EvmAddress
}

type CurrencyID struct {
	Token    TokenSymbol
	DEXShare struct {
		DexShare0 DexShare
		DexShare1 DexShare
	}
	ERC20                EvmAddress
	StableAssetPoolToken types.U32
	LiquidCrowdloan      types.U32
	ForeignAsset         types.U32
}

type Balance types.U128
type EventDexSwapped struct {
	Phase            types.Phase
	Trader           types.AccountID
	Path             []CurrencyID
	LiquidityChanges []Balance
}

type Events struct {
	types.EventRecords
	Dex_Swapped []EventDexSwapped //nolint:stylecheck,golint

}
