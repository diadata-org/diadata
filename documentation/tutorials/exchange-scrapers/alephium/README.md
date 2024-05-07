https://github.com/alephium/go-sdk/tree/master

https://explorer.alephium.org/

https://github.com/alephium/token-list

### openapi

https://node.mainnet.alephium.org/docs

https://backend.mainnet.alephium.org/docs/#/

Each token pairs has its own contract. 

You can do something like this to get all token pairs:

```
curl https://backend.mainnet.alephium.org/contracts/vyrkJHG49TXss6pGAz2dVxq5o7mBXNNXAV18nAeqVT1R/sub-contracts?limit=100
```

we do have endpoints for supporting fetching events. 

For AYIN, each token pair emits its own swap event, 

for example for ALPH <-> USDT pair (contract address: 2A5R8KZQ3rhKYrW7bAS4JTjY9FCFLJg6HjQpqSFZBqACX), 

following endpoints are useful:

1. 
```
curl https://node.mainnet.alephium.org/events/contract/2A5R8KZQ3rhKYrW7bAS4JTjY9FCFLJg6HjQpqSFZBqACX?start=0&limit=10

```

This returns the first 10 events for the ALPH <-> USDT token pair, the eventIndex of the swap events is 2. e.g.
```
{
      "blockHash": "00000000000bd8c4cc2df9caad18afd86fa6d9b09cb9ed61ef5ff34427515080",
      "txId": "0cd81aec1ee6744dc119db6f86328b4c846af86ce885e02d60582121012344f0",
      "eventIndex": 2,
      "fields": [
        {
          "type": "Address",
          "value": "1BjwWZEFaWYD9HLyNASTiVZD1ct4Z75Fpv8oGW7C9BWjz"
        },
        {
          "type": "U256",
          "value": "500000000000000000000"
        },
        {
          "type": "U256",
          "value": "0"
        },
        {
          "type": "U256",
          "value": "0"
        },
        {
          "type": "U256",
          "value": "117658826"
        },
        {
          "type": "Address",
          "value": "1BjwWZEFaWYD9HLyNASTiVZD1ct4Z75Fpv8oGW7C9BWjz"
        }
      ]
    }
```    

The schema of the events is event

```
    Swap(sender: Address, amount0In: U256, amount1In: U256, amount0Out: U256, amount1Out: U256, to: Address)
```

The response of this endpoint also contains a field nextStart, 
which is useful to continue fetching the next batch of events until it reaches to the end.

2.
```curl https://node.mainnet.alephium.org/events/contract/2A5R8KZQ3rhKYrW7bAS4JTjY9FCFLJg6HjQpqSFZBqACX/current-count```

This endpoint returns the current event count for this contract, 
we can use this number of as the value to fetch events from now going forward

The SDK has more dev friendly support for event subscription but it requires having the artifacts of the contracts, 
otherwise we can use the above endpoints to poll the events. 
We have written a general description of the event here (https://docs.alephium.org/dapps/events/), 
which might be helpful as well.

each token pairs has its own contract. You can do something like this to get all token pairs:
```
curl https://backend.mainnet.alephium.org/contracts/vyrkJHG49TXss6pGAz2dVxq5o7mBXNNXAV18nAeqVT1R/sub-contracts?limit=100
```

## way to get token info

Other than the token list, Fungible token contracts also implement a sets of standard interfaces to get name, symbol, decimals, etc. We can call these functions separately or in a single HTTP multi-call, e.g.

```

> curl https://node.mainnet.alephium.org/contracts/multicall-contract -d '{"calls": [{"group":0, "address": "zSRgc7goAYUgYsEBYdAzogyyeKv3ne3uvWb3VDtxnaEK", "methodIndex": 0}, {"group":0, "address": "zSRgc7goAYUgYsEBYdAzogyyeKv3ne3uvWb3VDtxnaEK", "methodIndex": 1}, {"group":0, "address": "zSRgc7goAYUgYsEBYdAzogyyeKv3ne3uvWb3VDtxnaEK", "methodIndex": 2}, {"group":0, "address": "zSRgc7goAYUgYsEBYdAzogyyeKv3ne3uvWb3VDtxnaEK", "methodIndex": 3}]}'  | jq '.results | .[] | .returns'

[
  {
    "type": "ByteVec",
    "value": "55534454"
  }
]
[
  {
    "type": "ByteVec",
    "value": "546574686572205553442028416c706842726964676529"
  }
]
[
  {
    "type": "U256",
    "value": "6"
  }
]
[
  {
    "type": "U256",
    "value": "1786596516040"
  }
]
```

Where zSRgc7goAYUgYsEBYdAzogyyeKv3ne3uvWb3VDtxnaEK is the USDT contract on Alephium, 

55534454 is the hex reprentation of symbol (USDT), 

546574686572205553442028416c706842726964676529 is the hex represenation of name (Tether USD (AlphBridge)), 

6 is decimal, 

1786596516040 is total supply

This doc about fungible tokens might be helpful as well (https://docs.alephium.org/tokens/fungible-tokens)

## how to interpret swap event

Given the token pair address 2A5R8KZQ3rhKYrW7bAS4JTjY9FCFLJg6HjQpqSFZBqACX, we can do 

```
➜  Blockchain curl https://node.mainnet.alephium.org/contracts/call-contract -d '{"address": "2A5R8KZQ3rhKYrW7bAS4JTjY9FCFLJg6HjQpqSFZBqACX", "methodIndex": 7, "group": 0}' jq .returns
[
  {
    "type": "ByteVec",
    "value": "0000000000000000000000000000000000000000000000000000000000000000"
  },
  {
    "type": "ByteVec",
    "value": "556d9582463fe44fbd108aedc9f409f69086dc78d994b88ea6c9e65f8bf98e00"
  }
]
```
to understand that in this token pair, the 1st token is ALPH (0000000000000000000000000000000000000000000000000000000000000000 is ALPH's special token id), and 2nd token is USDT. 

In the swap event (which is mapped to fields in the response)
```
event Swap(sender: Address, amount0In: U256, amount1In: U256, amount0Out: U256, amount1Out: U256, to: Address)
```

if we are swapping from ALPH to USDT, then amount0In ((fields[1]) will be the amount for ALPH and amount1Out (fields[4]) will be the amount for USDT. 

If we are swapping from USDT to ALPH, then amount1In ((fields[2]) will be the amount for USDT and amount0Out (fields[3]) will be the amount for ALPH.
