---
description: Overview over deployed oracle contracts on our supported chains
---

# Deployed Contracts

DIA operates oracles on multiple blockchains. Here, we provide an overview over the deployed oracle contracts on each supported chain.

DIA Development Oracle contracts are smart contracts that provide a selected range of asset prices on various blockchains for live testing on a respective Mainnet and Testnet.\
The contracts are upgraded and exchanged on a rolling basis and are not maintained indefinitely.\
DIA Development Oracle contracts are not intended to be integrated in a dApp. DIA deploys dedicated contracts for dApps. Please request a dedicated oracle by contacting the team on [Discord](https://discord.com/invite/zFmXtPFgQj) or the [DIA DAO Forum](https://dao.diadata.org).

### Ethereum

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://etherscan.io/address/0xa93546947f3015c986695750b8bbea8e26d65856">0xa935...5856</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}

{% tab title="Ropsten Testnet" %}
| Published Assets                     | Smart Contract Address                                                                   | Oracle Type                                                                                                                   |
| ------------------------------------ | ---------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x8147...2d9a](https://etherscan.io/address/0x814712cc9fa606a4b372b87cd27775959e052d9a) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract) |
{% endtab %}

{% tab title="Kovan Testnet" %}
| Published Assets                                              | Smart Contract Address                                                                         | Oracle Type                                                                                                                     |
| ------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md)                          | [0xb73d...A660](https://kovan.etherscan.io/address/0xb73db1A6a85219742fbd0fC7cc275c62209aA660) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract)   |
| [Coingecko Symbols](guest-quotations/coingecko-quotations.md) | [0x50e0...eF1C](https://kovan.etherscan.io/address/0x50e087d98a33ceb1ced159ad9255d6f228f2ef1c) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

### Binance Smart Chain

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                                                      | Smart Contract Address                                                                  | Oracle Type                                                                                                                     |
| --------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md)                                  | [0xf35b...693c](https://bscscan.com/address/0xf35bee4b6727d2d1c9167c5fb4d51855d6bb693c) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract)   |
| [Coinmarketcap Symbols](guest-quotations/coinmarketcap-quotations.md) | [0xbAFE...0835](https://bscscan.com/address/0xbafee71d40babc12a3d0b2b8937ee62d3a070835) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
| [Coingecko Symbols](guest-quotations/coingecko-quotations.md)         | [0x4814...06d8](https://bscscan.com/address/0x48140d0116964f05c97f08e0b3271d78b12506d8) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}

{% tab title="Testnet" %}
| Published Assets                                                     | Smart Contract Address                                                                           | Oracle Type                                                                                                                     |
| -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md)                                 | [0xf35b...B693c](https://testnet.bscscan.com/address/0xf35bee4b6727d2d1c9167c5fb4d51855d6bb693c) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract)   |
| [Coinmarketcap Oracle](guest-quotations/coinmarketcap-quotations.md) | [0x42D4...10a8](https://testnet.bscscan.com/address/0x42d44f1c45349d47d34976ce3a2ff0c3dd3210a8)  | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

### Polygon

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                                                   | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xf44b3c104f39209cd8420a1d3ca4338818aa72ab](https://polygonscan.com/address/0xf44b3c104f39209cd8420a1d3ca4338818aa72ab) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |


{% endtab %}

{% tab title="Mumbai Testnet" %}
__

| Published Assets                     | Smart Contract Address                                                                                       | Oracle Type                                                                                                                   |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | __[0xA3d2...d50a](https://explorer-mumbai.maticvigil.com/address/0xA3d2127F85041729fec05Ca483b302ddb806d50a) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract) |
{% endtab %}
{% endtabs %}

### xDaiChain

{% tabs %}
{% tab title="Mainnet" %}


| Published Assets                     | Smart Contract Address                                                                              | Oracle Type                                                                                                                   |
| ------------------------------------ | --------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xcF23...7569](https://blockscout.com/poa/xdai/address/0xcF2374824C2Ff84F07Ff4adcA074dfeDDa5C7569) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract) |
{% endtab %}

{% tab title="Sokol Testnet" %}
| Published Assets                     | Smart Contract Address                                                                               | Oracle Type                                                                                                                   |
| ------------------------------------ | ---------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x45D4...5f78](https://blockscout.com/poa/sokol/address/0x45D4B75228Ed3ee068A64dD1d5b53094A5015f78) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract) |
{% endtab %}
{% endtabs %}

### Astar/Shiden

{% tabs %}
{% tab title="Astar" %}
| Published Assets                     | Smart Contract Address                                                                           | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xd793...Ec75](https://blockscout.com/astar/address/0xd79357ebb0cd724e391f2b49a8De0E31688fEc75) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}

{% tab title="Shiden Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                                         | Oracle Type                                                                                                                     |
| ------------------------------------ | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xCe78...815d](https://blockscout.com/shiden/address/0xCe784F99f87dBa11E0906e2fE954b08a8cc9815d/transactions) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}

{% tab title="Shibuya Testnet" %}
| Published Assets                     | Smart Contract Address                                                                         | Oracle Type                                                                                                                     |
| ------------------------------------ | ---------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x1232...733E](https://shibuya.subscan.io/account/0x1232acd632dd75f874e357c77295da3f5cd7733e) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

### Avalanche

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://snowtrace.io/address/0x1fe94dfcb35a020ca05ab94bfd6e60f14eecfa31">0x1fe94dfcb35a020ca05ab94bfd6e60f14eecfa31</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}

{% tab title="Fuji Testnet" %}
| Published Assets                     | Smart Contract Address                                                                                        | Oracle Type                                                                                                                     |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x1cDF...F4f4](https://cchain.explorer.avax-test.network/address/0x1cDFEfC93D97E1B09e040a1f2d04b170eb60F4f4) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

### Aurora

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://explorer.mainnet.aurora.dev/address/0xf4e9C0697c6B35fbDe5a17DB93196Afd7aDFe84f/">0xf4e9...e84f</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract">Key/Value Oracle</a></td><td></td></tr></tbody></table>
{% endtab %}

{% tab title="Testnet" %}
| Published Assets                     | Smart Contract Address                                                                                  |                                                                                                                                 |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xf4e9...e84f](https://explorer.testnet.aurora.dev/address/0xf4e9C0697c6B35fbDe5a17DB93196Afd7aDFe84f) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

### Metis

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://andromeda-explorer.metis.io/address/0x6E6E633320Ca9f2c8a8722c5f4a993D9a093462E">0x6E6E...462E</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract">Key/Value Oracle</a></td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

### [Moonbeam](https://docs.moonbeam.network/learn/dapps-list/oracles/dia/)

{% tabs %}
{% tab title="Moonbeam" %}
| Published Assets                     | Smart Contract Address                                                                                  | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x1f1B...728f](https://blockscout.moonbeam.network/address/0x1f1BAe8D7a2957CeF5ffA0d957cfEDd6828D728f) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}

{% tab title="Moonriver" %}
| Published Assets                     | Smart Contract Address                                                                                                         | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x11f74b94afb5968119c98ea277a2b73208bb39ab](https://moonriver.moonscan.io/address/0x11f74b94afb5968119c98ea277a2b73208bb39ab) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}

{% tab title="Moonbeam Testnet" %}
| Published Assets                     | Smart Contract Address                                                                                                                                | Oracle Type                                                                                                                     |
| ------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x5425F5d4ba2B7dcb277C369cCbCb5f0E7185FB41](https://moonbase-blockscout.testnet.moonbeam.network/address/0x5425F5d4ba2B7dcb277C369cCbCb5f0E7185FB41) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

### Arbitrum

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://arbiscan.io/address/0xd041478644048d9281f88558e6088e9da97df624">0xd041478644048d9281f88558e6088e9da97df624</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

### Fantom

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://ftmscan.com/address/0xc5ca9c52d3d8d7f9bb17beeb85c2c3d119ab504f">0xC5cA...504f</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract">Key/Value Oracle</a></td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

### Celo

| Contract Name | Contract Address                                                                                                           |
| ------------- | -------------------------------------------------------------------------------------------------------------------------- |
| DIA Oracle    | [0xCd8E18890E416Aa7ab09aa793b406C187747C687](https://explorer.celo.org/address/0xCd8E18890E416Aa7ab09aa793b406C187747C687) |

### NEAR Testnet

| Contract Name      | Contract Address                                                                                   |
| ------------------ | -------------------------------------------------------------------------------------------------- |
| DIA Request Oracle | [@contract.dia-test.testnet](https://explorer.testnet.near.org/accounts/contract.dia-test.testnet) |
