# Oracles in DIA

## Immutable Oracle Feeds

The Coinhub oracles are located in a [single smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef). This smart contract holds mappings from an index to an asset's name, its price, and supply data. By using `getCoinInfo(coin_name)` it is possible to retrieve this data. Along with the actual data there is a metadata field for the timestamp of the last update. The return value consists of the price in US Cents, the circulating supply, the timestamp of the update and the short handle of the asset.

We also feed the current USDT price traded at the decentraliced Exchange Bancor, this value can be retrieved by querying `Bancor` .

Updates are supplied by the DIA Oracle service that periodically supplies updates into the smart contract. In the event view, the latest updates can be seen. The event fields show the values in the following locations:

* Price in Cent is in the second field \(display as `number`\)
* Supply is in the third field \(display as `number`\)
* UNIX Timestamp of last update is displayed in the fourth field \(display as `number`\)
* Name is displayed in the sixth and last field \(display as `text`\)

## Example DApp; ECB/FED FX Rates for Cryptos

As an example application, we implemented and deployed a [simple smart contract that converts crypto asset prices from our oracle from USD to EUR](https://etherscan.io/address/0xccb30bf12177705d41ac208802a6066482a76eaa). By calling `getAssetEurRate()` \(use the "Read Contract" tab for easy web access\) the requested crypto asset price is retrieved from our oracle and multiplied by the ECB exchange rate for EUR-USD, which we also publish in that oracle contract. The rate is displayed in a fix comma format with five decimal digits. The source code of that example contract can be found at the "Code" tab. This shows how easy it is to integrate our oracle into any application: compile your application against the oracle interface and set the address, where the oracle is deployed, afterwards. Ideally, this is done in a way that allows updates should the oracle address change at any time in the future.

