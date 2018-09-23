# Circulating Supply Numbers

Circulating Supply is a metric to determine the actually tradeable volume of an asset.
In case of cryptocurrencies, this number is determined in a different way for every coin.
Circulating supply is always measured in a unitless dimension, i.e., the number of coins (or tokens).
In contrast to traditional assets like stocks, numbers of circulating supply are often volatile in crypto markets, because of mechanics like mining that create more coins all the time.
In general, supply data is determined as close to the actual source as possible.
In most cases, monitoring the blockchain and tracking mining/minting/burning events is necessary as part of the methodology.

## API Access
Our crypto circulating supply data can be retrieved from our API using <https://api.diadata.org/v1/supply/TLA>, with TLA being the short name of a currency.
As an example, the current Bitcoin circulating supply is stored at <https://api.diadata.org/v1/supply/BTC>.

## Measurement Methodology
For each coin we measure, we have different methodology to determine the circulating supply.
### Bitcoin
For Bitcoin, we consider all mined coins to be part of the circulating supply.
We start at the genesis block and measure the current block height and mining reward.
For each block, the reward is added to the supply from the previous block.

### Ethereum
For Ethereum, we determine all mined Ether to be part of the circulating supply.
Additionally, uncle block rewards are considered and added to the result.

### Ripple
For Ripple, we take the amount of XRP generated in the genesis event and subtract the number of XRP held by the Ripple foundation.

### IOTA
IOTA has a fixed number of tokens in its network, which is equal to the circulating supply.

### Tether
For Tether, we determine the circulating supply by retrieving the number of Tether from their official website <https://wallet.tether.to/transparency>.
