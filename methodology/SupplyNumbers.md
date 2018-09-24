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
A full node is run to monitor the blockchain directly.
We start at the genesis block and measure the current block height and mining reward.
A pre-determined mining reward is generated for every block.
Initially, the reward was 50 BTC and is halved every 210,000 blocks.
The current supply is determined using this formula: ![Image of Bitcoin reward formula](https://latex.codecogs.com/png.latex?\large&space;s=\sum_{n=0}^{b}{\frac{50}{2^{\lfloor\frac{n}{210000}\rfloor}}}).
We consider *b* the current block height and *s* the circulating supply.

### Ethereum
For Ethereum, we determine all mined Ether to be part of the circulating supply.
A full node is run to monitor the blockchain directly.
Each block reward is 3.0 ETH.
Additionally, uncle block rewards are considered and added to the result.
More details on uncle blocks and their influence on mining reward can be found [in the Ethereum documentation](https://github.com/ethereum/wiki/wiki/Mining#mining-rewards).

### Ripple
For Ripple circulating supply, DIA retrieves data from the API the Ripple foundation is running.
Ripple generated all coins in their genesis event, but is releasing them only slowly over time.
To get the latest amount of XRP a query to the [xrp_distribution endpoint](https://data.ripple.com/v2/network/xrp_distribution?limit=1&descending=true) is evaluated.

### EOS
For EOS circulating supply, a full node is run to monitor the blockchain directly.
Using the RPC API it is possible to retrieve the number of circulating tokens by calling [get_currency_stats](https://developers.eos.io/eosio-nodeos/reference#get_currency_stats).

### Bitcoin Cash
Bitcoin Cash is a hardfork of Bitcoin, thus the same methodology is used but based their view of the blockchain.
For Bitcoin Cash, we consider all mined coins to be part of the circulating supply.
A full node is run to monitor the blockchain directly.
We start at the genesis block and measure the current block height and mining reward.
A pre-determined mining reward is generated for every block.
Initially, the reward was 50 BTC and is halved every 210,000 blocks.
The current supply is determined using this formula: ![Image of Bitcoin-Cash reward formula](https://latex.codecogs.com/png.latex?\large&space;s=\sum_{n=0}^{b}{\frac{50}{2^{\lfloor\frac{n}{210000}\rfloor}}}).
We consider *b* the current block height and *s* the circulating supply.

### Stellar
Stellar created 100 Billion XLM (lumens) in their genesis event with a yearly inflation of 1 percent.
Because Stellar distinguishes between available and distributed coins, we determine the circulating amount of stellar by querying the official Stellar API.
[The relevant endpoint](https://dashboard.stellar.org/api/lumens) contains a field to query all distributed coins.

### Litecoin
Litecoin is very similar to Bitcoin.
One of the major differences is that the time between two blocks is 2.5 minutes instead of 10 in Bitcoin.
To have a coin mining reward schedule similar to Bitcoin, the halving period was set to 840000, i.e., four times the amount of Bitcoin.
DIA runs a full Litecoin node to monitor the blockchain and determine the amount of currently available LTC by this formula: ![Image of Litecoin reward formula](https://latex.codecogs.com/png.latex?\large&space;s=\sum_{n=0}^{b}{\frac{50}{2^{\lfloor\frac{n}{840000}\rfloor}}}).

### Tether
For Tether, we determine the circulating supply by retrieving the number of Tether from their official website <https://wallet.tether.to/transparency>.
While Tether is using a blockchain protocol, it is impossible to reliably determine the actually circulating supply because no official documentation is known to us that lists the addresses that are considered as reserve accounts.

### Cardano
Cardano is capped at [45 billion ADA](https://cardanodocs.com/cardano/monetary-policy/).
To determine the circulating supply, we run our own instance of the [Cardano Explorer](https://github.com/diadata-org/cardano-explorer-docker) and trace every minting event.
All coins created in these events are added up to determine the circulating supply.

### Monero
Monero is mined over inifinite time.
To determine the circulating supply, we run an instance of a Monero node and query this node periodically.
In its internal RPC API, the function [get_coinbase_tx_sum](https://getmonero.org/resources/developer-guides/daemon-rpc.html#get_coinbase_tx_sum) is used to retrieve the amount of mined coins.

### DASH
Dash is also forked from the Bitcoin network.
Thus, by running a node, the circulating supply can be determined.
The current supply is determined using this formula: ![Image of Bitcoin reward formula](https://latex.codecogs.com/png.latex?\large&space;s=\sum_{n=0}^{b}{\frac{50}{2^{\lfloor\frac{n}{210000}\rfloor}}}).

### IOTA
IOTA has a fixed number of tokens in its network, which is equal to the circulating supply.
[The initial crowdsale generated 2,779,530,283,277,761 IOTA tokens](https://docs.iota.org/introduction/iota-token/the-iota-token), which are all distributed and thus part of the circulating supply.

### Tron
The circulating supply of Tron is determined by running a Tron node and monitoring the blockchain.
However, funds that have been sent to address `TLsV52sRDL79HXGGm9yzwKibb6BeruhUzy` are considered burned (or [blackholed](https://github.com/tronprotocol/wiki/blob/master/docs/Technical_Overview_of_TRON.rst)).

### NEO

### Ethereum Classic
