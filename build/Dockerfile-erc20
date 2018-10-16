FROM golang:latest as build

WORKDIR $GOPATH

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/scrapers/erc20

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/erc20 /bin/erc20

ENTRYPOINT ["erc20"]
