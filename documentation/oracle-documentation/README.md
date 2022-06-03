# ⛓ Oracle Documentation

Oracles are smart contract interfaces that bring data from an external source into a smart contract. Smart contracts run in an isolated environment in a virtual machine and can not "measure" any outside information by themselves. The only way to get external data into a smart contract is by executing a transaction with the external data as a payload. With that data, a smart contract can perform calculations and operations that depend on this data, such as calculating an interest based on the reference interest rates published by a central bank.

DIA is capable of publishing financial data with such an oracle so that any smart contract can read and use it from within the Ethereum virtual machine. By generating verifiable oracle data, any user can use financial data in smart contracts. For each of our asset classes, there is an example data set available in the oracle.

To learn how to interact with DIA oracle smart contracts

{% content-ref url="access-the-oracle.md" %}
[access-the-oracle.md](access-the-oracle.md)
{% endcontent-ref %}

## Multi-chain delivery

DIA operates oracles on 25+ blockchains. To enable developers to test DIA's oracle services, DIA provides Development Oracles for each supported chain. These Development Oracle contracts are smart contracts that provide a selected range of asset prices on various blockchains for live testing on a respective Mainnet and Testnet.

Get an overview on the following link:

{% content-ref url="deployed-contracts.md" %}
[deployed-contracts.md](deployed-contracts.md)
{% endcontent-ref %}

## Data Categories

### Price feeds

{% content-ref url="crypto-assets.md" %}
[crypto-assets.md](crypto-assets.md)
{% endcontent-ref %}

{% content-ref url="broken-reference" %}
[Broken link](broken-reference)
{% endcontent-ref %}

{% content-ref url="fiat-prices.md" %}
[fiat-prices.md](fiat-prices.md)
{% endcontent-ref %}

### Randomness

{% content-ref url="randomness-oracle.md" %}
[randomness-oracle.md](randomness-oracle.md)
{% endcontent-ref %}

### Guest quotations

DIA provides oracles for several guest quotations.

{% content-ref url="guest-quotations/coingecko-quotations.md" %}
[coingecko-quotations.md](guest-quotations/coingecko-quotations.md)
{% endcontent-ref %}

{% content-ref url="guest-quotations/coinmarketcap-quotations.md" %}
[coinmarketcap-quotations.md](guest-quotations/coinmarketcap-quotations.md)
{% endcontent-ref %}

### Indexes

{% content-ref url="broken-reference" %}
[Broken link](broken-reference)
{% endcontent-ref %}

### Rates

{% content-ref url="broken-reference" %}
[Broken link](broken-reference)
{% endcontent-ref %}

{% content-ref url="broken-reference" %}
[Broken link](broken-reference)
{% endcontent-ref %}

## Other

{% content-ref url="polkadot-offchain-worker.md" %}
[polkadot-offchain-worker.md](polkadot-offchain-worker.md)
{% endcontent-ref %}

{% content-ref url="polkadot-medianizer.md" %}
[polkadot-medianizer.md](polkadot-medianizer.md)
{% endcontent-ref %}

{% content-ref url="solana-oracle.md" %}
[solana-oracle.md](solana-oracle.md)
{% endcontent-ref %}

{% content-ref url="near-request-oracle.md" %}
[near-request-oracle.md](near-request-oracle.md)
{% endcontent-ref %}

## Immutable Oracle Feeds

The DIA oracle data can be accessed in an Ethereum [smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef). This smart contract holds mappings from an index to an asset's name, its price value, and – if applicable – supply data.

By using the read function `getCoinInfo(coin_name)` it is possible to retrieve this data. Along with the actual data there is a metadata field for the timestamp of the last update.

Updates are supplied by the DIA Oracle service that periodically supplies updates into the smart contract. Each update also generates an event, so that the latest updates can be seen in the event view on Etherscan.

## Example DApp: ECB FX Rates for Cryptos

As an example application, we implemented and deployed a [simple smart contract that converts crypto asset prices from our oracle from USD to EUR](https://etherscan.io/address/0xccb30bf12177705d41ac208802a6066482a76eaa).&#x20;

Call `getAssetEurRate()` with a crypto asset name as argument in order to request the current crypto asset price from our oracle which is then converted to EUR by using the ECB exchange rate for EUR-USD, also published in the oracle contract.  In the above link, you can use the "Read Contract" tab for easy web access. The rate is displayed in a fix comma format with five decimal digits. The source code of that example contract can be found at the "Code" tab. This shows how easy it is to integrate our oracle into any application: compile your application against the oracle interface and set the address where the oracle is deployed afterwards. Ideally, this should be done in a way that allows updates if the oracle address changes at any time in the future.

## Gas Estimation

All DIA oracles use the go-ethereum gas estimation to estimate the gas costs for an update contract call. To ensure timely delivery of the transaction, the gas estimate is scaled to 110% of the original estimate.
