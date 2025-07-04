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

We use a **two-pipeline approach** for automated testing, building, and container image publishing that provides proper validation before production releases.

### Pipeline Overview

**Pipeline 1: PR Validation** (`pr-validation.yml`)
- **Trigger:** Pull request creation/updates with changes to `cmd/**/go.mod`
- **Purpose:** Validate changes before merge with real releases for Go module resolution
- **Actions:** Create release → Lint → Test → Build validation (no registry push)

**Pipeline 2: Production Build & Publish** (`production-deployment.yml`)  
- **Trigger:** Push to master branch with changes to `cmd/**/go.mod`
- **Purpose:** Build and publish validated container images to registry
- **Actions:** Verify release → Build → Push to IBM Cloud Container Registry

---

### Pipeline 1: PR Validation Workflow

1. **Trigger Detection**
   - **Trigger:** Pull request events (opened, updated, synchronized) - runs on all PRs
   - **Change detection:** Scan for modified `cmd/**/go.mod` files to determine if service validation is needed
   - **Action:** If no go.mod files changed, skip all validation steps; otherwise process changed services

2. **Service Detection & Version Extraction**
   - **Scan changed files:** Find all modified `cmd/**/go.mod` files
   - **Extract service info:** 
     - Service name from directory path (e.g., `cmd/http/graphqlServer` → `graphqlServer`)
     - Version from `github.com/diadata-org/diadata vX.X.X` dependency line
   - **Validate version format:** Ensure it matches `vX.X.X` pattern
   - **Output:** JSON array of services with name, path, and version

3. **Create Combined Release**
   - **Tag creation:** Creates git tag with version number (e.g., `v1.4.586`)
   - **GitHub release:** Single release covering all changed services for Go module resolution during linting
   - **Release body:** Includes all updated services, PR number, and validation status
   - **Purpose:** Enable Go module resolution and prepare for production deployment

4. **Lint Service (Per Service)**
   - **Setup:** Checkout code, install Go 1.22
   - **Dependencies:** Run `go mod tidy` and `go mod download`
   - **Note:** Real releases (not drafts) are required because Go module resolution during linting needs to resolve `github.com/diadata-org/diadata vX.X.X` dependencies, which requires actual Git tags to exist
   - **Linting:** Execute `golangci-lint` with project configuration
   - **Validation:** Ensure code quality standards before testing

5. **Test Service (Per Service)**
   - **Setup:** Go environment with 30s proxy indexing wait
   - **Dependency resolution:** Retry `go mod tidy` up to 8 times (4 min timeout)
   - **Test execution:**
     - If `*_test.go` files exist → Run `go test -v ./...`
     - If no tests → Run `go build -v .` to verify compilation
   - **Purpose:** Ensure functionality before build validation

6. **Build Validation (Per Service)**
   - **Dockerfile discovery:** Search for Dockerfile in order:
     - `{service_path}/Dockerfile`
     - `Dockerfile-{service_name}` (root)
     - `{service_path}/Dockerfile-{service_name}`
     - `build/Dockerfile-{service_name}` ✅ (most common)
   - **Build validation:** Build Docker image locally without pushing to registry
   - **Purpose:** Verify Docker build process works correctly

7. **PR Summary**
   - **Status report:** Comprehensive validation results for all services
   - **Pass criteria:** All lint, test, and build validations must succeed
   - **Output:** Ready/not ready for merge status with detailed feedback

---

### Pipeline 2: Production Build & Publish Workflow

1. **Trigger Detection**
   - **Trigger:** Push to master branch (after PR merge) with path filter for `cmd/**/go.mod` files
   - **Path filter:** Only triggers when `cmd/**/go.mod` files are actually modified in the push
   - **Action:** Detect services that changed in the merge

2. **Verify Release (Per Service)**
   - **Release verification:** Ensure the release created during PR validation exists
   - **Production update:** Update release notes with production deployment status
   - **Release body:** Production-ready release notes with container image info
   - **Purpose:** Confirm release readiness and mark production deployment

3. **Build & Push to Container Registry (Per Service)**
   - **Docker setup:** Configure Docker Buildx and authenticate with IBM Cloud Container Registry
   - **Image build:** Build production Docker images with version and latest tags
   - **Registry push:** Push to `us.icr.io/dia-registry/{service}:{version}` and `:latest`
   - **Container availability:** Images are now available for deployment by operations teams

4. **Build Summary**
   - **Audit trail:** Complete build report with actor, commit, timestamp
   - **Container tracking:** List all published container images with registry URLs
   - **Status reporting:** Success/failure status for container image publishing

---

### Key Benefits

- **🔒 Safety:** No production container images without PR validation
- **🏃 Speed:** Parallel processing of multiple services
- **📋 Visibility:** Comprehensive status reporting at each stage  
- **🔄 Reliability:** Retry logic and robust error handling
- **📦 Traceability:** Full audit trail from PR to container registry

## Support

- [Discord](support/discord.md)
- [Request a custom oracle](support/request-a-custom-oracle.md)

