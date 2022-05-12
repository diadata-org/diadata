---
cover: .gitbook/assets/1_dZqpEo-3_IlZsrZ_J9X7Xw.png
coverY: 0
---

# 🤔 What is DIA?

## What is DIA?&#x20;

**DIA (Decentralised Information Asset) is a multi-chain, end-to-end, open-source data and oracle platform for Web3.**

The DIA platform enables the sourcing, validation and sharing of transparent and verified data feeds for traditional and digital financial applications. DIA’s institutional-grade data feeds cover asset prices, metaverse data, lending rates and more.

DIA’s data is directly sourced from a broad array of on-chain and off-chain sources at individual a trade-level. This allows DIA feeds to be fully customised with regards to the mix of sources and methodologies, resulting in tailor-made, high resilience feeds, setting a new paradigm for oracles.

## The DIA Ecosystem

DIA is an open-source ecosystem for financial data. The aim of DIA is to make transparent, accurate and verifiable data accessible for Web3, by bringing together data providers, users and DAO community members.

![](.gitbook/assets/Architecture.svg)

<mark style="color:blue;">**Data sources:**</mark> every second, thousands of trade level data points from multiple on-chain and off-chain sources (CEXs, DEXs, DeFi, NFT, Metaverse, etc) are aggregated into the DIA Platform.

<mark style="color:blue;">**DIA Open-Source Platform:**</mark>

* **Open Contributors**: to collect the data, an open community of developers build scrapers to connect new data sources with the DIA open-source platform.
* **Governors/validators:** a community of token holders discuss and vote to improve the platform; crowd-approve and validate features and govern the DIA Association.
* **Off-chain computation:** the ingested raw data is computed and sanitised based on fully transparent methodologies to avoid temporary off-data and outliers.&#x20;

<mark style="color:blue;">**Blockchains**</mark><mark style="color:blue;">:</mark> new data points are shipped via API or dedicated oracle smart contracts and made available to developers in 25+ layer 1 and layer 2 blockchain networks.

<mark style="color:blue;">**dApps:**</mark> DIA oracles are easy to integrate into smart contracts to power a broad range of DeFi, TradFi and Metaverse use cases, such as lending, lending, staking or stablecoins.

## DIA's data offering

### <mark style="color:blue;">Data categories</mark>

DIA's crowd-sourcing approach to data sourcing uniquely positions DIA to utilise the broadest possible set of sources and ensure maximum coverage of asset price data. Any data feed that is publicly accessible can be sourced, regardless of whether it is listed on exchanges and what its trading volume is - there is no dependency on third-party data providers.&#x20;

Here is a high-level overview of DIA's data offering:

![](<.gitbook/assets/Data categories.svg>)

### <mark style="color:blue;">Data sources</mark>

By leveraging a community of developers and a distributed network of nodes, DIA scrapes, collects and aggregates data directly from multiple on-chain sources or Centralised Exchange APIs at individual trade levels.&#x20;

Here is a high-level overview of data sources DIA leverages:

![to seel all data sources, visit the Data Sources section](<.gitbook/assets/Data sources.svg>)

### <mark style="color:blue;">Delivery methods</mark>

DIA enables users to select how they ingest data using multiple delivery methods present on a constantly growing L1 / L2 ecosystem, powered by a network of strong industry partners.

#### **On-chain delivery | DIA xNode**

Oracles can be customised to trigger updates based on:

* **Request**: updates are triggered via a dedicated smart contract or API call.
* **Time**: updates are triggered in predefined time intervals
* **Deviation**: updates are triggered by a deviation from the last reported value

#### **Off-chain delivery  | DIA xBase**

&#x20;Data can be delivered off-chain via

* **Rest API**: delivers data feeds in a predefined, standardised format
  * **Update frequency**: 120s
  * **Methodology**: [MAIR](documentation/methodology/digital-assets/exchangeprices/mair-moving-average-with-interquartile-range-filter.md)
  * **Sources**: [all available](documentation/data-sources/natively-sourced-data.md)
* **GraphQL**: enables more flexibility and direct adjustments of feed attributions
  * **Update frequency**: custom
  * **Methodology**: custom
  * **Sources**: custom

## Support

[![image-text](https://content.diadata.org/wp-content/uploads/2022/02/SP-technical.svg)](https://discord.gg/zFmXtPFgQj)[![image-text](https://content.diadata.org/wp-content/uploads/2022/02/SP-Custom-Oracle.svg)](https://medium.com/dia-insights/how-to-submit-a-dia-custom-delivery-request-cdr-in-5-minutes-6f88b0a4ca56)

## Contribute

[![image-text](https://content.diadata.org/wp-content/uploads/2022/02/CB-Product.svg)](https://gitcoin.co/diadata-org)[![image-text](https://content.diadata.org/wp-content/uploads/2022/02/CB-Ecosystem.svg)](https://app.dework.xyz/dia-dao)
