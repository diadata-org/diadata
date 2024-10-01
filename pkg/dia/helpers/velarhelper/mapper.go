package velarhelper

import (
	"fmt"
	"strings"
)


func ExtractSwapInfo(response string) SwapInfo {
	swapInfo := SwapInfo{}

	swapInfo.AmountIn = extractUint(response, "amt-in u")
	swapInfo.AmountOut = extractUint(response, "amt-out u")
	swapInfo.TokenIn = extractString(response, "token-in '")
	swapInfo.TokenOut = extractString(response, "token-out '")
	swapInfo.Symbol = extractString(response, "symbol \"")
	swapInfo.LpToken = extractString(response, "lp-token '")
	swapInfo.Token0 = extractString(response, "token0 '")
	swapInfo.Token1 = extractString(response, "token1 '")

	swapInfo.ProtocolFee = extractFee(response, "protocol-fee (tuple (den u", " (num u")
	swapInfo.ShareFee = extractFee(response, "share-fee (tuple (den u", " (num u")
	swapInfo.SwapFee = extractFee(response, "swap-fee (tuple (den u", " (num u")

	return swapInfo
}

func extractUint(response, prefix string) uint64 {
	start := strings.Index(response, prefix)
	if start == -1 {
		return 0
	}
	start += len(prefix)

	end := strings.IndexAny(response[start:], " )")
	if end == -1 {
		return 0
	}

	var result uint64
	fmt.Sscanf(response[start:start+end], "%d", &result)
	return result
}

func extractString(response, prefix string) string {
	start := strings.Index(response, prefix)
	if start == -1 {
		return ""
	}
	start += len(prefix)

	end := strings.IndexAny(response[start:], " )")
	if end == -1 {
		return ""
	}

	return response[start : start+end]
}

func extractFee(response, denPrefix, numPrefix string) Fee {
	fee := Fee{}
	fee.Denominator = extractUint(response, denPrefix)
	fee.Numerator = extractUint(response, numPrefix)

	return fee
}