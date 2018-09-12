FROM golang:latest as build

WORKDIR $GOPATH

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/api-golang/blockchain-scrapers/blockchains/bitcoin/scrapers/

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/scrapers /bin/btc

CMD ["btc"]