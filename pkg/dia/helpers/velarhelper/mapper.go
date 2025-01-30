package velarhelper

import (
	"encoding/hex"
	"fmt"
	"strings"

	stacks "github.com/diadata-org/diadata/pkg/dia/helpers/stackshelper"
)

func DecodeSwapEvents(tx stacks.Transaction) ([]SwapEvent, error) {
	result := make([]SwapEvent, 0, len(tx.Events))
	lastEventIndex := -1

	for i, e := range tx.Events {
		if !isSwapEvent(e) {
			continue
		}
		bytes, err := hex.DecodeString(e.ContractLog.Value.Hex[2:])
		if err != nil {
			return nil, err
		}
		event, err := stacks.DeserializeCVTuple(bytes)
		if err != nil {
			return nil, err
		}

		info := SwapEvent{TxID: tx.TxID, Timestamp: tx.BlockTime}
		info.AmountIn, err = stacks.DeserializeCVUint(event["amt-in"])
		if err != nil {
			return nil, err
		}
		info.AmountOut, err = stacks.DeserializeCVUint(event["amt-out"])
		if err != nil {
			return nil, err
		}

		poolAssets, err := decodePoolAssets(event["pool"])
		if err != nil {
			return nil, err
		}
		info.TickerID = poolAssets[0] + "_" + poolAssets[1]

		if val, ok := event["token-in"]; ok {
			info.TokenIn, err = stacks.DeserializeCVPrincipal(val)
			if err != nil {
				return nil, err
			}
		} else {
			sender, err := stacks.DeserializeCVPrincipal(event["user"])
			if err != nil {
				return nil, err
			}
		outer:
			for j := lastEventIndex + 1; j < i; j++ {
				ev := &tx.Events[j]
				if ev.Asset.EventType != "transfer" ||
					ev.Asset.Sender != sender ||
					ev.Asset.Recipient != e.ContractLog.ContractID {
					continue
				}

				switch ev.Type {
				case "fungible_token_asset":
					info.TokenIn = strings.Split(ev.Asset.ID, "::")[0]
					break outer
				case "stx_asset":
					for _, a := range poolAssets {
						if a == DeployerAddress+".wstx" {
							info.TokenIn = a
							break outer
						}
					}
				}
			}
		}

		switch info.TokenIn {
		case poolAssets[0]:
			info.TokenOut = poolAssets[1]
		case poolAssets[1]:
			info.TokenOut = poolAssets[0]
		default:
			err = fmt.Errorf("unable to decode input token address: no matches for %s", info.TokenIn)
			return nil, err
		}

		result = append(result, info)
		lastEventIndex = i
	}

	return result, nil
}

func isSwapEvent(ev stacks.Event) bool {
	log := ev.ContractLog
	return ev.Type == "smart_contract_log" &&
		log.Topic == "print" &&
		(strings.HasPrefix(log.ContractID, DeployerAddress) || strings.HasPrefix(log.ContractID, DeployerAddressV2)) &&
		strings.Contains(log.Value.Repr, "(op \"swap\")")
}

func decodePoolAssets(src []byte) ([2]string, error) {
	result := [2]string{}

	pool, err := stacks.DeserializeCVTuple(src)
	if err != nil {
		return result, err
	}

	result[0], err = stacks.DeserializeCVPrincipal(pool["token0"])
	if err != nil {
		return result, err
	}
	result[1], err = stacks.DeserializeCVPrincipal(pool["token1"])
	if err != nil {
		return result, err
	}

	return result, nil
}
