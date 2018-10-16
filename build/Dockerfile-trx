FROM golang:latest as build

RUN go get github.com/PaddyMc/go-tron/calculations

WORKDIR $GOPATH

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/tron/scrapers/trx

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/trx /bin/trx

CMD ["trx"]