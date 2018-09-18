FROM golang:latest as build

WORKDIR $GOPATH/src/

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/api-golang/blockchain-scrapers/eos/scrapers

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/eos /bin/eos

CMD ["eos"]