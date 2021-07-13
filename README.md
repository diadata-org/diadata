---
description: 'Our Mission: Wikipedia for Financial Data'
---

# Introduction

[![API Status](https://badgen.net/uptime-robot/status/m784441379-1bdbacd4cd81bf46c13bdb1f?label=API)](https://docs.diadata.org/api/docs/api) [![Pull Requests](https://badgen.net/github/prs/diadata-org/diadata?label=Pull%20Requests)](https://github.com/diadata-org/diadata/pulls)

![](.gitbook/assets/180926_dia_assets-57.png)

## What is DIA about?

Welcome to the [DIAdata.org](https://diadata.org/) platform. DIA is an ecosystem for open financial data in a financial smart contract ecosystem. The target of DIA is to bring together data analysts, data providers and data users. In general, DIA provides a reliable and verifiable bridge between off-chain data from various sources and on-chain smart contracts that can be used to build a variety of financial dApps.

## Get started right away

[Tutorials for contributors](https://docs.diadata.org/documentation/tutorials)

[API Documentation](https://docs.diadata.org/documentation/api-1)

If you want to dive into our API without further ado, feel free to take the below tables as a starting point. They present an excerpt of our API endpoints  a complete list of which can be found in our API Documentation. 

### Digital Assets

| Class | Description | API Example | API Documentation | Oracle | Number of Assets |
| :--- | :--- | :--- | :--- | :--- | :--- |
| CoinMarketCap Quotation | Actual price of the top 50 crypto assets as listed on CMC | [DIA Price](https://api.diadata.org/v1/foreignQuotation/CoinMarketCap/DIA) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#guest-quotation) for a detailed description | [CoinMarketCap Oracle](https://bscscan.com/address/0xbafee71d40babc12a3d0b2b8937ee62d3a070835) | 50 |
| Crypto Price | Actual price of a crypto asset such as Bitcoin \(BTC\) or Ether \(ETH\) | [BTC Price](https://api.diadata.org/v1/quotation/BTC) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#quotation) for a detailed description. | Oracle for [crypto assets](https://docs.diadata.org/documentation/oracle-documentation/crypto-assets) | ~800 |
| Supply | Actual circulating supply of a crypto asset | [BTC Supply](https://api.diadata.org/v1/supply/BTC) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#supply) for a detailed description | Oracle for [crypto assets](https://docs.diadata.org/documentation/oracle-documentation/crypto-assets) | ~150 |
| Exchange | Centralized and decentralized crypto exchanges | [List of exchanges](https://api.diadata.org/v1/exchanges) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#exchanges) for a detailed description |  | 28 |
| DeFi Interest Rate | Lending and borrowing rates on decentralized finance protocols | [ETH lending and borrowing rate](https://api.diadata.org/v1/defiLendingRate/DYDX/ETH) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#defi-interest-rate) for a detailed description | Oracle for [DeFi Lending](https://docs.diadata.org/documentation/oracle-documentation/defi-protocol-rates-and-states) | &gt;30 |
| DeFi Protocol Information | Total locked value and information on lending protocol | [DYDX](https://api.diadata.org/v1/defiLendingState/DYDX) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#defi-lending-protocol) for a detailed description | Oracle for [DeFi Protocol](https://docs.diadata.org/documentation/oracle-documentation/defi-protocol-rates-and-states) | 10 |
| Coingecko Quotation | Actual price of the top 100 crypto assets as listed on coingecko | [BTC Price](https://api.diadata.org/v1/foreignQuotation/Coingecko/BTC) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#guest-quotation) for a detailed description | [Coingecko Oracle](https://docs.diadata.org/documentation/oracle-documentation/guest-quotations) | 100 |
| Farming Pool | Rates and balances of farming pools | [YFI - WETH Pool](https://api.diadata.org/v1/FarmingPoolData/YFI/WETH) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#farming-pool-data) for a detailed description | [Farming Pool Oracle](https://docs.diadata.org/documentation/oracle-documentation/farming-pools) | 3000 |

### Traditional Assets

| Class | Description | API Example | API Documentation | Oracle | Number of Assets |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Interest Rate | Interest rate of overnight markets such as SOFR or SONIA incl. historical data | [SOFR](https://api.diadata.org/v1/interestrate/SOFR/2020-08-03) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#interest-rate) for a detailed description | _Coming soon_: Oracle for [interest rates](https://docs.diadata.org/documentation/methodology/traditional-assets/overnight-rates) | 3 |
| Fiat Prices | Price of fiat currencies vs USD | [List of currencies](https://api.diadata.org/v1/fiatQuotations) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#fiat-currency-exchange-rates) for a detailed description | Oracle for [fiat prices](https://docs.diadata.org/documentation/oracle-documentation/fiat-prices) | 10 |

## Applications

### Product use Cases

Based on the data we collect at DIA we build products an excerpt of which can be found below.

* [CVI](https://diadata.org/crypto-volatility-index/) - An interactive chart of the Crypto Volatility Index
* [Rates Calculator](https://diadata.org/compounded-rates-calculator/) - An interactive chart for the calculation of customized indices

### Partner Integrations

DIA's data feeds are being used by a range of market actors in the DeFi and CeFi space. Below is an excerpt of partners. A number of additional users are in testing phase. This list will be periodically updated.

* [Gather](https://gather.network/) - A platform that allows publishers to monetize without ads, provides businesses & developers to access cheap and reliable processing power.
* [CryptoLocally](https://cryptolocally.com/en) - A gateway for Fastest, cheapest, and easiest way to buy and sell crypto in your local currency
* [Polkastarter](https://www.polkastarter.com/) - A DEX built for cross-chain token pools and auctions, enabling projects to raise capital on a decentralized, permissionless and interoperable environment based on Polkadot.
* [ADD.XYZ](https://add.xyz/) - A full-stack aggregation and management platform der DeFi products.  DIA oracles supply off-chain & on-chain information.
* [Elrond](https://elrond.com/) - A blockchain architecture, designed to bring a 1000-fold cumulative improvement in throughput and execution speed. DIA oracles supply off-chain & cross-chain information.
* [ankr](https://www.ankr.com/) - Building an infrastructure platform and marketplace for Web3-stack deployment. DIA oracles supply off-chain & cross-chain information.
* [Hedget](https://www.hedget.com/) - Designed to be a DeFi option trading platform. DIA supplies price feeds to value Hedget’s derivative products.
* [SIX/F10 PSD II Sandbox](https://f10-sandbox-portal.apps.ndgit.com/#/apis) - PSD II and pricing data for FinTechs
* [Open-Source Calculation Agent](https://blockstate.com/decentral-calculation-agent/)

## **Feedback**

Your feedback helps improve DIA in order for us to find data and applications that are the most useful for you and the broader developer community.

Link: [Feedback Form](https://docs.google.com/forms/d/e/1FAIpQLSePxDwbEURjes4nw8GUzaT-XfYttRw_6F2xAR607FKACsn7ew/viewform)

