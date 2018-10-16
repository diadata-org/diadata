FROM golang:latest as build

WORKDIR $GOPATH

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/bitcoin/scrapers/btc

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/btc /bin/btc

CMD ["btc"]