package helpers

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// SymbolIsBlackListed return true if the symbol is blacklisted
// Symbols are blacklisted when they are duplicated and there is no
// approach to resolve the conflict
func SymbolIsBlackListed(symbol string) bool {
	switch symbol {
	case "UNI-V2":
		return true
	case "VET":
		return true
	case "ACC":
		return true
	case "ALT":
		return true
	case "APE":
		return true
	case "ARB":
		return true
	case "AT":
		return true
	case "ATX":
		return true
	case "BBK":
		return true
	case "BET":
		return true
	case "BIT":
		return true
	case "BITS":
		return true
	case "BLZ":
		return true
	case "BTM":
		return true
	case "CAN":
		return true
	case "CAT":
		return true
	case "CBC":
		return true
	case "CEN":
		return true
	case "CMCT":
		return true
	case "CMS":
		return true
	case "CMT":
		return true
	case "CPC":
		return true
	case "CRC":
		return true
	case "DFT":
		return true
	case "EDR":
		return true
	case "ENT":
		return true
	case "ETT":
		return true
	case "EVN":
		return true
	case "EXC":
		return true
	case "FAIR":
		return true
	case "FT":
		return true
	case "GCC":
		return true
	case "GENE":
		return true
	case "GET":
		return true
	case "GOT":
		return true
	case "HC":
		return true
	case "HERO":
		return true
	case "HMC":
		return true
	case "HNC":
		return true
	case "HOT":
		return true
	case "ICN":
		return true
	case "IQ":
		return true
	case "KEY":
		return true
	case "KNC":
		return true
	case "KNT":
		return true
	case "LBTC":
		return true
	case "LNC":
		return true
	case "MAG":
		return true
	case "MORE":
		return true
	case "MTC":
		return true
	case "NET":
		return true
	case "NTK":
		return true
	case "ONG":
		return true
	case "ORS":
		return true
	case "PAI":
		return true
	case "PASS":
		return true
	case "PLC":
		return true
	case "PLY":
		return true
	case "PUT":
		return true
	case "PXC":
		return true
	case "QBT":
		return true
	case "RCN":
		return true
	case "RED":
		return true
	case "SCC":
		return true
	case "SLT":
		return true
	case "SPD":
		return true
	case "TTC":
		return true
	case "WEB":
		return true
	case "XIN":
		return true
	case "XRA":
		return true
	default:
		return false
	}
}

// AddressIsBlacklisted returns true if a token address is blacklisted
func AddressIsBlacklisted(address common.Address) bool {
	switch strings.ToLower(address.Hex()) {
	case "0x3e191a6ef96f87092fe8dce0d3f01977b08d6acf":
		return true
	case "0x8870f11b5d16f1fd4ce26aff514566621dca4828":
		return true
	case "0x5a7a0ad8d92fbeee4a9a68d35cd29fe248cad790":
		return true
	case "0xa93f73b5723f1ac86736a9e4310a39707868ccf2":
		return true
	case "0xe4d247b7cebd5e3957ee41a247074457a1e7402d":
		return true
	case "0x96d32a2035ea5491017543d4e55c17f58fbf1c57":
		return true
	case "0x70a72833d6bf7f508c8224ce59ea1ef3d0ea3a38":
		return true
	case "0x343373daeea3a8a3b0465b94bc706258767dfab5":
		return true
	case "0x3258ed6f9b939558e39a3e6f27023233b43e3ee8":
		return true
	case "0xca45cd9eb7e995b97a4f062fe2ff4f196369e0d9":
		return true
	case "0x50b7f2b98bf473848656d6b986725cfa72b7fd6b":
		return true
	case "0xfa524f57bbf98a2dc9bbf33b9f62f2202bfd7c68":
		return true
	case "0xf111820c5216d8fea8cae0fbcb87f9f188fc1887":
		return true
	case "0x6a16e1144a9129577823c4751e23142370d67d14":
		return true
	case "0xcb9e106e86b3c2349ea7c6dda63cd8b8267135aa":
		return true
	case "0x216867dd2bc12753bf59638950b275a094c23358":
		return true
	case "0x40c190fd64888e55a4206ee10f9d39d744237fe7":
		return true
	case "0x582ba13e41034212b050fa22026fc74425a9c941":
		return true
	case "0x0707fcd0f4c7f875a2bb744904b4cc008453f046":
		return true
	case "0x84633d3453a1f3582f7ab33b8116a1346308a084":
		return true
	default:
		return false
	}
}
