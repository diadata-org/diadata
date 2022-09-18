package acalaswap

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

// Adapted from: https://github.com/AcalaNetwork/acala.js/blob/master/packages/type-definitions/src/json/types.json,
// https://github.com/AcalaNetwork/Acala/blob/master/primitives/src/currency.rs
var TokenSymbolMap = map[string]types.U8{
	// 0 - 19: Acala & Polkadot native tokens
	"ACA":  0,
	"AUSD": 1,
	"DOT":  2,
	"LDOT": 3,
	"TAP":  4,
	// 20 - 39: External tokens (e.g. bridged)
	"RENBTC": 20,
	"CASH":   21,
	// 40 - 127: Polkadot parachain tokens

	// 128 - 147: Karura & Kusama native tokens
	"KAR":  128,
	"KUSD": 129,
	"KSM":  130,
	"LKSM": 131,
	"TAI":  132,
	// 148 - 167: External tokens (e.g. bridged)
	// 149: Reserved for renBTC
	// 150: Reserved for CASH
	// 168 - 255: Kusama parachain tokens
	"BNC":   168,
	"VSKSM": 169,
	"PHA":   170,
	"KINT":  171,
	"KBTC":  172,
}
var TokenSymbolMapReverse map[int]string

func init() {
	TokenSymbolMapReverse = make(map[int]string)
	for k, v := range TokenSymbolMap {
		TokenSymbolMapReverse[int(v)] = k
	}
}

type TokenSymbol struct {
	Token types.U8
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
