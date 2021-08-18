FROM golang:1.14 as build

WORKDIR $GOPATH/src/

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/cmd/stock-scrapers
RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/stock-scrapers /bin/stock-scrapers
COPY --from=build /go/src/github.com/diadata-org/diadata/config /config/

CMD ["stockScraper"]