# Steps to run an nft trade scraper locally
1. Navigate to the `deployments/local/exchange-scraper` directory of the project.
2. Run the required services using `docker-compose up -d`. This command will run and prepare the PostgreSQL database.
3. Set the required environment variables using the following commands:

```sh
export USE_ENV=true
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=password
export POSTGRES_HOST=localhost
export POSTGRES_DB=postgres
```

Or simply by sourcing the `local.env` inside the `deployments/local/exchange-scraper` directory.

4. Execute `main.go` from `cmd/services/blockchainservice` to populate the PostgreSQL database with different blockchain configs:
```sh
export DIA_CONFIG_DIR=./config
go run ./cmd/services/blockchainservice/main.go
```

5. Finally, run the scraping executable flagged as follows:

```sh
cd cmd/nftTradescrapers
export ETH_URI_REST=https://mainnet.infura.io/v3/xxxxx
go run main.go -nftclass MyNftTradeScraper
```

For an illustration of how to create an nft trade scraper, you can have a look at the `pkg/dia/nft/nftTrade-scrapers/opensea.go` file.

