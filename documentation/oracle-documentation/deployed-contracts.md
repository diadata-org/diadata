---
description: Overview over deployed oracle contracts on our supported chains
---

# Deployed Contracts

DIA operates oracles on multiple blockchains. Here, we provide an overview over the deployed oracle contracts on each supported chain.

DIA Development Oracle contracts are smart contracts that provide a selected range of asset prices on various blockchains for live testing on a respective Mainnet and Testnet.\
The contracts are upgraded and exchanged on a rolling basis and are not maintained indefinitely.\
DIA Development Oracle contracts are not intended to be integrated in a dApp. DIA deploys dedicated contracts for dApps. Please request a dedicated oracle by contacting the team on [Discord](https://discord.com/invite/zFmXtPFgQj) or the [DIA DAO Forum](https://dao.diadata.org/).

The audit of our oracle smart contract can be found here:

{% file src="../../.gitbook/assets/02_Smart Contract Audit_DIA_Oracle_v2.pdf" %}

### Ethereum

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://etherscan.io/address/0xa93546947f3015c986695750b8bbea8e26d65856">0xa935...5856</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}

{% tab title="Ropsten Testnet" %}
| Published Assets                     | Smart Contract Address                                                                           | Oracle Type                                                                                                                   |
| ------------------------------------ | ------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x1E27...d5eC](https://ropsten.etherscan.io/address/0x1e27d6b118e2e618e1b902e85428a27f49edd5ec) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract) |
{% endtab %}

{% tab title="Kovan Testnet" %}
| Published Assets                                              | Smart Contract Address                                                                         | Oracle Type                                                                                                                     |
| ------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md)                          | [0xb73d...A660](https://kovan.etherscan.io/address/0xb73db1A6a85219742fbd0fC7cc275c62209aA660) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract)   |
| [Coingecko Symbols](guest-quotations/coingecko-quotations.md) | [0x50e0...eF1C](https://kovan.etherscan.io/address/0x50e087d98a33ceb1ced159ad9255d6f228f2ef1c) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Ethereum chain oracle</summary>

1. BTC / USD
2. DIA / USD
3. ETH / USD
4. USDC / USD
5. FTM / USD
6. SDN / USD
7. KSM / USD
8. ASTR / USD

</details>

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

<details>

<summary>See all available assets on the BSC chain oracle</summary>

1. BNB / USD
2. BTC / USD
3. DIA / USD
4. USDC / USD
5. ETH / USD
6. USDT / USD
7. XRP / USD

</details>

### Polygon

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://polygonscan.com/address/0xf44b3c104f39209cd8420a1d3ca4338818aa72ab">0xf44b...72ab</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>


{% endtab %}

{% tab title="Mumbai Testnet" %}
__

| Published Assets                     | Smart Contract Address                                                                                       | Oracle Type                                                                                                                   |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | __[0xA3d2...d50a](https://explorer-mumbai.maticvigil.com/address/0xA3d2127F85041729fec05Ca483b302ddb806d50a) | [CoinInfo Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-coininfo-oracle-contract) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Polygon chain oracle</summary>

1. BTC / USD
2. DIA / USD
3. ETH / USD
4. SDN / USD
5. KSM / USD
6. ASTR / USD
7. FTM / USD

</details>

### Gnosis Chain

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

<details>

<summary>See all available assets on the Gnosis chain oracle</summary>

1. DIA / USD
2. USDC / USD
3. XRP / USD
4. USDT / USD
5. ETH / USD
6. BNB / USD
7. BTC / USD

</details>

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
| Published Assets                     | Smart Contract Address                                                                             | Oracle Type                                                                                                                     |
| ------------------------------------ | -------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x1232...733E](https://blockscout.com/shibuya/address/0x1232AcD632Dd75f874E357c77295Da3f5Cd7733E) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Astar/Shiden chains oracles</summary>

1. SDN / USD
2. DIA / USD
3. FTM / USD
4. Metis / USD
5. KSM / USD
6. ASTR / USD
7. ETH / USD
8. BTC / USD

</details>

### Avalanche

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://snowtrace.io/address/0x1fe94dfcb35a020ca05ab94bfd6e60f14eecfa31">0x1fe9...fa31</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}

{% tab title="Fuji Testnet" %}
| Published Assets                     | Smart Contract Address                                                                                        | Oracle Type                                                                                                                     |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x1cDF...F4f4](https://cchain.explorer.avax-test.network/address/0x1cDFEfC93D97E1B09e040a1f2d04b170eb60F4f4) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Avalanche chain oracle</summary>

1. SDN / USD
2. DIA / USD
3. ASTR / USD
4. FTM / USD
5. ETH / USD
6. KSM / USD
7. BTC / USD

</details>

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

<details>

<summary>See all available assets on the Aurora chain oracle</summary>

1. ASTR / USD
2. SDN / USD
3. DIA / USD
4. KSM / USD
5. FTM / USD
6. ETH / USD
7. BTC / USD

</details>

### Metis

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://andromeda-explorer.metis.io/address/0x6E6E633320Ca9f2c8a8722c5f4a993D9a093462E">0x6E6E...462E</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract">Key/Value Oracle</a></td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Metis chain oracle</summary>

1. SDN / USD
2. DIA / USD
3. ASTR / USD
4. FTM / USD
5. ETH / USD
6. KSM / USD
7. BTC / USD

</details>

### Moonbeam

{% tabs %}
{% tab title="Moonbeam" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://blockscout.moonbeam.network/address/0x1f1BAe8D7a2957CeF5ffA0d957cfEDd6828D728f">0x1f1B...728f</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}

{% tab title="Moonriver" %}
| Published Assets                     | Smart Contract Address                                                                            | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x11f7...39ab](https://moonriver.moonscan.io/address/0x11f74b94afb5968119c98ea277a2b73208bb39ab) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}

{% tab title="Moonbeam Testnet" %}
| Published Assets                     | Smart Contract Address                                                                                                   | Oracle Type                                                                                                                     |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x5425...FB41](https://moonbase-blockscout.testnet.moonbeam.network/address/0x5425F5d4ba2B7dcb277C369cCbCb5f0E7185FB41) | [Key/Value Oracle](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Moonbeam chain oracle</summary>

1. FTM / USD
2. SDN / USD
3. USDC / USD
4. DIA / USD
5. ETH / USD
6. BTC / USD
7. ASTR / USD
8. KSM / USD
9. MOVR / USD

</details>

### Nervos

{% tabs %}
{% tab title="Godwoken Testnet" %}
| Published Assets                     | Smart Contract Address                                                                           | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x96c5...cb99](https://v1.aggron.gwscan.com/account/0x96c5f1e50c4393efa890699cd9aecf3fb58dcb99) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Nervos chain oracle</summary>

1. FTM / USD
2. SDN / USD
3. USDC / USD
4. DIA / USD
5. ETH / USD
6. BTC / USD
7. ASTR / USD
8. KSM / USD
9. MOVR / USD

</details>

### Arbitrum

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://arbiscan.io/address/0xd041478644048d9281f88558e6088e9da97df624">0xd041...f624</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2">Key/Value Oracle V2</a></td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Arbitrum chain oracle</summary>

1. ASTR / USD
2. KSM / USD
3. FTM / USD
4. SDN / USD
5. DIA / USD
6. ETH / USD
7. BTC / USD

</details>

### Fantom

{% tabs %}
{% tab title="Mainnet" %}
<table><thead><tr><th>Published Assets</th><th>Smart Contract Address</th><th>Oracle Type</th><th data-type="files">Audit</th></tr></thead><tbody><tr><td><a href="crypto-assets.md">DIA Asset Prices</a></td><td><a href="https://ftmscan.com/address/0xc5ca9c52d3d8d7f9bb17beeb85c2c3d119ab504f">0xC5cA...504f</a></td><td><a href="https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract">Key/Value Oracle</a></td><td></td></tr></tbody></table>
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Fantom chain oracle</summary>

1. DIA / USD
2. ASTR / USD
3. SDN / USD
4. ETH / USD
5. FTM / USD
6. KSM / USD
7. MOVR/ USD
8. USDC / USD

</details>

### Evmos

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                    | Oracle Type                                                                                                                           |
| ------------------------------------ | ----------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x5d60...E6ed](https://evm.evmos.org/address/0x5d60C36A600391C3dFc5d76ad18959163613E6ed) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}

{% tab title="Testnet" %}
| Published Assets                     | Smart Contract Address                                                                    | Oracle Type                                                                                                                           |
| ------------------------------------ | ----------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x70eA...cD07](https://evm.evmos.dev/address/0x70eA682A109695D465DB975B40E8fDbBFda5cD07) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Evmos chain oracle</summary>

1. FTM / USD
2. SDN / USD
3. USDC / USD
4. DIA / USD
5. ETH / USD
6. BTC / USD
7. ASTR / USD
8. KSM / USD
9. MOVR / USD

</details>

### Fuse

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                       | Oracle Type                                                                                                                           |
| ------------------------------------ | -------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x61a5...320e](https://explorer.fuse.io/address/0x61a598Cd6340B8edcb4faE7Eabcd117Ff371320e) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Fuse chain oracle</summary>

1. ASTR / USD
2. SDN / USD
3. DIA / USD
4. FTM / USD
5. BTC / USD
6. KSM / USD
7. ETH / USD

</details>

### Telos

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                      | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x7512...a5ef](https://www.teloscan.io/address/0x7512fb605c45cedff8552eaca2a020c13a04a5ef) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}

{% tab title="Testnet" %}
| Published Assets                     | Smart Contract Address                                                                          | Oracle Type                                                                                                                           |
| ------------------------------------ | ----------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x261c...12dd](https://testnet.teloscan.io/address/0x261cf410d0a83193d647e47c35178288d99e12dd) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Telos chain oracle</summary>

1. ASTR / USD
2. KSM / USD
3. FTM / USD
4. SDN / USD
5. USDC / USD
6. DIA / USD
7. ETH / USD
8. BTC / USD

</details>

### Velas

{% tabs %}
{% tab title="EVM Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                            | Oracle Type                                                                                                                           |
| ------------------------------------ | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0x0A7d...8856](https://evmexplorer.velas.com/address/0x0A7dC648C44e31636252be2267B86e0d9E1F8856) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Velas chain oracle</summary>

1. ASTR / USD
2. KSM / USD
3. FTM / USD
4. SDN / USD
5. USDC / USD
6. DIA/ USD
7. ETH / USD
8. BTC / USD

</details>

### Clover

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                  | Oracle Type                                                                                                                           |
| ------------------------------------ | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xCd2C...17a5](https://clvscan.com/address/0xcd2ca164a2aec86b03474d1a76d25a2aa0a517a5) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Clover chain oracle</summary>

1. ASTR / USD
2. KSM / USD
3. FTM / USD
4. SDN / USD
5. USDC / USD
6. DIA / USD
7. ETH / USD
8. BTC / USD

</details>

### Celo

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                     | Smart Contract Address                                                                        | Oracle Type                                                                                                                           |
| ------------------------------------ | --------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [DIA Asset Prices](crypto-assets.md) | [0xCd8E...C687](https://explorer.celo.org/address/0xCd8E18890E416Aa7ab09aa793b406C187747C687) | [Key/Value Oracle V2](https://docs.diadata.org/documentation/oracle-documentation/access-the-oracle#dia-key-value-oracle-contract-v2) |
{% endtab %}
{% endtabs %}

<details>

<summary>See all available assets on the Celo chain oracle</summary>

1. DIA / USD
2. USDC / USD
3. XRP / USD
4. BTC / USD
5. ETH / USD
6. BNB / USD

</details>

### NEAR

{% tabs %}
{% tab title="Mainnet" %}
| Published Assets                                                                   | Oracle Address                                                                             | Oracle Type                                   |
| ---------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ | --------------------------------------------- |
| [DIA Assets](https://docs.diadata.org/documentation/api-1/api-endpoints#quotation) | [@contract.diadata.near](https://explorer.mainnet.near.org/accounts/contract.diadata.near) | [NEAR Request Oracle](near-request-oracle.md) |
{% endtab %}

{% tab title="Testnet" %}
| Published Assets                                                                   | Oracle Address                                                                                     | Oracle Type                                   |
| ---------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | --------------------------------------------- |
| [DIA Assets](https://docs.diadata.org/documentation/api-1/api-endpoints#quotation) | [@contract.dia-test.testnet](https://explorer.testnet.near.org/accounts/contract.dia-test.testnet) | [NEAR Request Oracle](near-request-oracle.md) |
{% endtab %}
{% endtabs %}

{% hint style="info" %}
Asset price updates in NEAR blockchain are provided on request basis
{% endhint %}

<details>

<summary>See all available assets on the Near chain oracle</summary>

1. ASTR / USD
2. KSM / USD
3. FTM / USD
4. SDN / USD
5. USDC / USD
6. DIA / USD
7. ETH / USD
8. BTC / USD

</details>
