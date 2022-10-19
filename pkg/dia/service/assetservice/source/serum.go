package source

import (
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gagliardetto/solana-go/programs/serum"
	solanav2rpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/mr-tron/base58"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/rpc"
)

type SerumPair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     string
}

const (
	// Public Solana clients.
	// rpcEndpointSolana         = "https://nd-646-378-408.p2pify.com/b2b39b407964cd4688015cd3c9027199"
	rpcEndpointSolana         = "https://solana-api.projectserum.com"
	dexProgramAddress         = "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin" // refer - https://github.com/project-serum/serum-dex
	nameServiceProgramAddress = "namesLPneVptA9Z5rqUDD9tMTWEJwofgaYwp8cawRkX"
	dotTokenTLD               = "6NSu2tci4apRKQtt257bAVcvqYjB3zV2H1dWo56vgpa6"
	marketDataSize            = 388
)

type SerumAssetSource struct {
	solanaRpcClient   *rpc.Client
	rpcv2Client       *solanav2rpc.Client
	tokenNameRegistry map[string]tokenMeta
	assetChannel      chan dia.Asset
	doneChannel       chan bool
	blockchain        string
}

func NewSerumAssetSource(exchange dia.Exchange) *SerumAssetSource {

	var assetChannel = make(chan dia.Asset)
	var doneChannel = make(chan bool)
	var sas *SerumAssetSource

	exchangeFactoryContractAddress = ""

	sas = &SerumAssetSource{
		solanaRpcClient: rpc.NewClient(utils.Getenv("SOLANA_URI_REST", rpcEndpointSolana)),
		rpcv2Client:     solanav2rpc.New(utils.Getenv("SOLANA_URI_REST", rpcEndpointSolana)),
		assetChannel:    assetChannel,
		doneChannel:     doneChannel,
		blockchain:      dia.SOLANA,
	}

	go func() {
		sas.fetchAssets()
	}()
	return sas

}

func (sas *SerumAssetSource) Asset() chan dia.Asset {
	return sas.assetChannel
}

func (sas *SerumAssetSource) Done() chan bool {
	return sas.doneChannel
}

func (sas *SerumAssetSource) fetchAssets() {
	pairs, err := sas.getPairs()
	if err != nil {
		log.Error("get pairs: ", err)
		return
	}
	log.Info("Found ", len(pairs), " pairs")
	checkMap := make(map[string]struct{})
	sas.tokenNameRegistry, err = sas.getTokenNames()
	if err != nil {
		log.Error(err)
		return
	}
	for _, pair := range pairs {
		if tokenInfo, ok := sas.tokenNameRegistry[pair.BaseMint.String()]; ok {
			if _, ok := checkMap[tokenInfo.symbol]; !ok {
				checkMap[tokenInfo.symbol] = struct{}{}
				sas.assetChannel <- dia.Asset{
					Symbol:     tokenInfo.symbol,
					Name:       tokenInfo.name,
					Address:    tokenInfo.mint,
					Decimals:   tokenInfo.decimals,
					Blockchain: sas.blockchain,
				}
			}
		}
		if tokenInfo, ok := sas.tokenNameRegistry[pair.QuoteMint.String()]; ok {
			if _, ok := checkMap[tokenInfo.symbol]; !ok {
				checkMap[tokenInfo.symbol] = struct{}{}
				sas.assetChannel <- dia.Asset{
					Symbol:     tokenInfo.symbol,
					Name:       tokenInfo.name,
					Address:    tokenInfo.mint,
					Decimals:   tokenInfo.decimals,
					Blockchain: sas.blockchain,
				}
			}
		}
	}
	sas.doneChannel <- true
}

func (sas *SerumAssetSource) getPairs() ([]*serum.MarketV2, error) {

	// var dataslice solanav2rpc.DataSlice
	// offset := uint64(100 * 1e6)
	// length := uint64(1024)
	// dataslice.Offset = &offset
	// dataslice.Length = &length

	// resp, e := sas.rpcv2Client.GetProgramAccountsWithOpts(
	// 	context.Background(),
	// 	solanav2.MustPublicKeyFromBase58(dexProgramAddress),
	// 	&solanav2rpc.GetProgramAccountsOpts{
	// 		Commitment: solanav2rpc.CommitmentFinalized,
	// 		Encoding:   solanav2.EncodingBase64,
	// 		DataSlice:  &dataslice,
	// 		Filters: []solanav2rpc.RPCFilter{
	// 			{
	// 				DataSize: marketDataSize,
	// 			},
	// 		},
	// 	},
	// )
	// if e != nil {
	// 	log.Error("get program accounts v2: ", e)
	// }
	// out := make([]*serumv2.MarketV2, 0)
	// log.Info("here")
	// for _, keyedAcct := range resp {
	// 	acct := keyedAcct.Account
	// 	marketV2 := &serumv2.MarketV2{}
	// 	if err := marketV2.Decode(acct.Data.GetBinary()); err != nil {
	// 		return nil, fmt.Errorf("decoding market v2: %w", err)
	// 	}
	// 	out = append(out, marketV2)
	// }

	resp, err := sas.solanaRpcClient.GetProgramAccounts(
		solana.MustPublicKeyFromBase58(dexProgramAddress),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{},
		},
	)
	if err != nil {
		log.Error("get program accounts: ", err)
		return nil, err
	}
	out := make([]*serum.MarketV2, 0)
	for _, keyedAcct := range resp {
		acct := keyedAcct.Account
		marketV2 := &serum.MarketV2{}
		if err := marketV2.Decode(acct.Data); err != nil {
			return nil, fmt.Errorf("decoding market v2: %w", err)
		}
		out = append(out, marketV2)
	}
	return out, nil
}

func (sas *SerumAssetSource) getTokenNames() (map[string]tokenMeta, error) {
	names := make(map[string]tokenMeta)
	tldPublicKey := solana.MustPublicKeyFromBase58(dotTokenTLD)
	resp, err := sas.solanaRpcClient.GetProgramAccounts(
		solana.MustPublicKeyFromBase58(nameServiceProgramAddress),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					Memcmp: &rpc.RPCFilterMemcmp{
						Bytes: tldPublicKey[:],
					},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, keyedAcct := range resp {
		if t, ok := extractTokenMetaFromData(keyedAcct.Account.Data[96:]); ok {
			names[t.mint] = t
		}
	}
	return names, nil
}

type tokenMeta struct {
	name     string
	symbol   string
	mint     string
	decimals uint8
}

func extractTokenMetaFromData(data []byte) (tokenMeta, bool) {
	var t tokenMeta
	if len(data) > 0 {
		nameSize := int(data[0])
		nameStart := 4
		nameEnd := nameStart + nameSize
		if len(data) > nameEnd {
			t.name = string(data[nameStart:nameEnd])
			symbolSize := int(data[nameEnd])
			symbolStart := 4 + nameEnd
			symbolEnd := symbolStart + symbolSize
			if len(data) > symbolEnd {
				t.symbol = string(data[symbolStart:symbolEnd])
				mintSize := 32
				mintStart := symbolEnd
				mintEnd := mintStart + mintSize
				if len(data) > mintEnd {
					t.mint = base58.Encode(data[mintStart:mintEnd])
					t.decimals = data[mintEnd]
					return t, true
				}
			}
		}
	}
	return tokenMeta{}, false
}
