# 🏃 Get Started

## What is DIA?&#x20;

**DIA (Decentralised Information Asset) is a multi-chain, end-to-end, open-source data and oracle platform for Web3.**

The DIA platform enables the sourcing, validation and sharing of transparent and verified data feeds for traditional and digital financial applications. DIA’s institutional-grade data feeds cover asset prices, metaverse data, lending rates and more.

DIA’s data is directly sourced from a broad array of on-chain and off-chain sources at individual a trade-level. This allows DIA feeds to be fully customised with regards to the mix of sources and methodologies, resulting in tailor-made, high resilience feeds, setting a new paradigm for oracles.

Links: [Homepage](https://diadata.org/) | [Medium Blog](https://medium.com/dia-insights) | [Github](https://github.com/diadata-org) | [Twitter](https://twitter.com/DIAdata\_org) | [Discord](https://discord.gg/zFmXtPFgQj) | [Telegram](https://t.me/DIAdata\_org)

[![API Status](https://badgen.net/uptime-robot/status/m784441379-1bdbacd4cd81bf46c13bdb1f?label=API)](https://docs.diadata.org/api/docs/api) [![Pull Requests](https://badgen.net/github/prs/diadata-org/diadata?label=Pull%20Requests)](https://github.com/diadata-org/diadata/pulls)

## The DIA Ecosystem

DIA is an open-source ecosystem for financial data. The aim of DIA is to make transparent, accurate and verifiable data accessible for Web3, by bringing together data providers, users and DAO community members.

![architecture](https://user-images.githubusercontent.com/103409771/231812128-bb1d2fd0-5946-418b-a82c-641bf586a3d5.png)

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

![data offering](https://user-images.githubusercontent.com/103409771/231812400-71aafd2a-2ded-478c-970e-f4724fec656a.png)

### <mark style="color:blue;">Data sources</mark>

By leveraging a community of developers and a distributed network of nodes, DIA scrapes, collects and aggregates data directly from multiple on-chain sources or Centralised Exchange APIs at individual trade levels.&#x20;

Here is a high-level overview of data sources DIA leverages:

![to seel all data sources, visit the Data Sources section](https://user-images.githubusercontent.com/103409771/231812534-60fe69b3-2d9d-4f12-bab8-d14043f485d8.png)

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

## Contribute

DIA through its DAO operations is working towards building the biggest open-source data platform of Web3. To reach this goal, DIA leverages the power of communities and decentralisation on many levels to scale its product development and growth.&#x20;

Contribute to the DAO through the two verticals below:

- [Product development](contribute/product-development.md)
- [Ecosystem growth](contribute/ecosystem-growth.md)

## CI/CD Pipeline

For automated tests, build, deployments we use the following pipeline.
This is the full workflow:

1. **Trigger Detection**
   - Trigger: Push to master (or feature/automated-service-deployment for testing)
   - Path filter: Only triggers on changes to cmd/**/go.mod files
   - Action: Scan for modified go.mod files in the cmd/ directory

2. **Service Detection & Version Extraction**
   - Scan changed files: Find all modified cmd/**/go.mod files
   - Extract service info for each file:
     - Service name from directory path (e.g., cmd/http/graphqlServer → graphqlServer)
     - Version from github.com/diadata-org/diadata vX.X.X dependency line
   - Validate version format: Ensure it matches vX.X.X pattern
   - Output: JSON array of services with name, path, and version

3. **Create GitHub Release (Per Service)**
   - Tag creation: Creates git tag with version number (e.g., v1.4.586)
   - Release creation: GitHub release named with just the version (e.g., v1.4.586)
   - Release body: Includes service info and auto-generated content
   - Purpose: Ensures the version exists before testing/building

4. **Test Service (Per Service)**
   - Setup: Checkout code, install Go 1.22
   - Dependency resolution: Run go mod tidy to update go.sum
   - Test execution:
     - If *_test.go files exist → Run go test -v ./...
     - If no tests → Run go build -v . to verify compilation
   - Validation: Ensures code quality before proceeding

5. **Build & Push Docker Image (Per Service)**
   - Dockerfile discovery: Search for Dockerfile in order:
     - {service_path}/Dockerfile
     - Dockerfile-{service_name} (root)
     - {service_path}/Dockerfile-{service_name}
     - build/Dockerfile-{service_name} ✅ (most common)
   - Docker setup: Configure Docker Buildx
   - Registry login: Authenticate with IBM Cloud Container Registry
   - Image build:
     - Convert service name to lowercase for Docker compatibility
     - Build with tags: us.icr.io/dia-registry/{service}:{version} and latest
   - Image push: Push both version and latest tags to registry

6. **Pipeline Summary**
   - Dependency: Waits for all previous jobs to complete
   - Status reporting:
     - If changes detected: List processed services, success/failure status
     - If no changes: Report that pipeline was skipped
   - Output: Comprehensive summary of all pipeline activities

## Support

- [Discord](support/discord.md)
- [Request a custom oracle](support/request-a-custom-oracle.md)

