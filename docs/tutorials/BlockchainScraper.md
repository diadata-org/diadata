# Write your own blockchain scraper

## Add your own blockchain scraper
In order to add your own scraper for a blockchain supply source, you must adhere to our format.
We use Go modules for our scrapers, so that each data provider is living as an independent module.

## Directory structure

The working directory for blockchain scrapers is `diadata/blockchain-scrapers/`.
You must provide a Dockerfile for your blockchain client and for the Go wrapper that connects to our database.
An example dockerfile is provided for the Go bindings with a bitcoind client scraping the Bitcoin blockchain

```Docker
FROM golang:latest as build
WORKDIR $GOPATH/src/
COPY . .
WORKDIR $GOPATH/src/github.com/diadata-org/diadata/blockchain-scrapers/cmd
RUN go install
FROM gcr.io/distroless/base
COPY --from=build /go/bin/cmd /bin/blockchain-scrapers
CMD ["blockchain-scrapers"]
```

After cloning a container capable of executing Go, the files from the working directory are copied into the container.
Next, the go program located in `diadata/blockchain-scrapers/cmd` is built and installed and the finished binary is placed into a mininal distroless container.
From there it is executed using the statement in the last line.

## Blockchain Client

The bitcoind itself is initialized directly from the file `docker-compose.yml`.
Ideally, all blockchain clients are run from there directly

```
bitcoind:
	image:
		 kylemanna/bitcoind
	ports:
		 - "8332"
	volumes:
		 - /home/srv/bitcoind:/bitcoin
	command: btc_oneshot -prune=550 -rpcallowip=::/0 -disablewallet -rpcpassword=mysecretrpcdiapassword -rpcuser=mysecretrpcdiauser
	logging:
		 options:
				max-size: "50m"
	networks:
		 - scrapers-network
	deploy:
		 mode: global
		 placement:
				constraints:
					 - node.labels.bitcoind==true
		 restart_policy:
				delay: 2s
				window: 20s
```

Make sure to enable pruning so that only the latest mined blocks are stored in the container.
Add persistency in /home/srv for the daemon data, so it can restart without needing to resynch the blockchain.

## Go Wrapper

For your Go code, also add an entry into the docker-compose file so that your Dockerfile is automatically called.
An example entry should look something like this

```
btc: 
	build:
		 context: ../../../..
		 dockerfile: github.com/diadata-org/diadata/blockchain-scrapers/Dockerfile-btc
	image: blockchain-scrapers_btc 
	networks:
		 - scrapers-network
	logging:
		 options:
				max-size: "50m"
	secrets:
		 - api_diadata
```

Be sure to add your container to the `scrapers-network` virtual network.
Your credentials for the DIA API should be handled by the `secrets` directive in the section of your blockchain client.
To initialize the API from Go you can create

```go
var api *dia.Client
```

and then call

```go
var c dia.ConfigApi
configFile := "/run/secrets/api_diadata"
err := gonfig.GetConf(configFile, &c)
if err != nil {
	configFile = "../config/secrets/api_diadata.json"
	err = gonfig.GetConf(configFile, &c)
}
if err != nil {
	log.Println(err)
} else {
	log.Println("Loaded secret in", configFile)
}
api = dia.NewClient(&c)
```

For sending supply data to DIA, you can use `SendSupply` in the `ApiClient.go` file.

This is the basic structure of these files.
Run your Docker instance by calling

```shell
docker run <your instance>
```

from the main directory.
With `docker-compose up` you can test the entire system with your cryptocurrency VM and your Go wrapper in a virtual network.
