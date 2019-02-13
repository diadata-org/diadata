FROM golang:latest as build

WORKDIR $GOPATH

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/eos/scrapers/eos

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/eos /bin/eos

CMD ["eos"]
