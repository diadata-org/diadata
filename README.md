---
description: 'Our Mission: Wikipedia for Financial Data'
---

# Introduction

[![API Status](https://badgen.net/uptime-robot/status/m784441379-1bdbacd4cd81bf46c13bdb1f?label=API)](https://docs.diadata.org/api/docs/api) [![Pull Requests](https://badgen.net/github/prs/diadata-org/diadata?label=Pull%20Requests)](https://github.com/diadata-org/diadata/pulls) [![Netlify Status](https://api.netlify.com/api/v1/badges/4be89751-9655-472f-9bfe-c8e57b9528b2/deploy-status)](https://coinhub.diadata.org)

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
| Crypto Price | Actual price of a crypto asset such as Bitcoin \(BTC\) or Ether \(ETH\) | [BTC Price](https://api.diadata.org/v1/quotation/BTC) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#quotation) for a detailed description. | Oracle for [crypto assets](https://docs.diadata.org/documentation/oracle-documentation/crypto-assets) | 414 |
| Supply | Actual circulating supply of a crypto asset | [BTC Supply](https://api.diadata.org/v1/supply/BTC) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#supply) for a detailed description | Oracle for [crypto assets](https://docs.diadata.org/documentation/oracle-documentation/crypto-assets) | 414 |
| Exchange | Centralized and decentralized crypto exchanges | [List of exchanges](https://api.diadata.org/v1/exchanges) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#exchanges) for a detailed description |  | 15 |
| DeFi Interest Rate | Lending and borrowing rates on decentralized finance protocols | [ETH lending and borrowing rate](https://api.diadata.org/v1/defiLendingRate/DYDX/ETH) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#defi-interest-rate) for a detailed description | Oracle for [DeFi Lending](https://docs.diadata.org/documentation/oracle-documentation/defi-protocol-rates-and-states) | &gt;30 |
| DeFi Protocol Information | Total locked value and information on lending protocol | [DYDX](https://api.diadata.org/v1/defiLendingState/DYDX) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#defi-lending-state) for a detailed description | Oracle for [DeFi Protocol](https://docs.diadata.org/documentation/oracle-documentation/defi-protocol-rates-and-states) | 6 |

### Traditional Assets

| Class | Description | API Example | API Documentation | Oracle | Number of Assets |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Interest Rate | Interest rate of overnight markets such as SOFR or SONIA incl. historical data | [SOFR](https://api.diadata.org/v1/interestrate/SOFR/2020-08-03) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#interest-rate) for a detailed description | _Coming soon_: Oracle for [interest rates](https://docs.diadata.org/documentation/methodology/traditional-assets/overnight-rates) | 3 |
| Fiat Prices | Price of fiat currencies vs USD | [List of currencies](https://api.diadata.org/v1/fiatQuotations) | Click [here](https://docs.diadata.org/documentation/api-1/api-endpoints#fiat-currency-exchange-rates) for a detailed description | Oracle for [fiat prices](https://docs.diadata.org/documentation/oracle-documentation/fiat-prices) | 10 |

## Financial Products

Based on the data we collect at DIA we build products an excerpt of which can be found below.

* [CoinHub](http://coinhub.diadata.org) - A transparent, open-source alternative to CM
* [CVI](https://diadata.org/crypto-volatility-index/) - An interactive chart of the Crypto Volatility Index
* [Rates Calculator](https://diadata.org/compounded-rates-calculator/) - An interactive chart for the calculation of customized indices
* [SIX/F10 PSD II Sandbox](https://f10-sandbox-portal.apps.ndgit.com/#/apis) - PSD II and pricing data for FinTechs

## **Feedback**

Your feedback helps improve DIA in order for us to find data and applications that are the most useful for you and the broader developer community.

Link: [Feedback Form](https://docs.google.com/forms/d/e/1FAIpQLSePxDwbEURjes4nw8GUzaT-XfYttRw_6F2xAR607FKACsn7ew/viewform)

