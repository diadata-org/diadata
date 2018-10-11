FROM golang:latest as build

WORKDIR $GOPATH/src/

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/cmd/exchange-scrapers/ecb

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/ecb /bin/ecb
COPY --from=build /go/src/github.com/diadata-org/diadata/config/ /config/

CMD ["ecb"]
