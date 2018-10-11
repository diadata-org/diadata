FROM golang:latest as build

WORKDIR $GOPATH/src/

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/cmd/services/tradesBlockService

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/tradesBlockService /bin/tradesBlockService
COPY --from=build /go/src/github.com/diadata-org/diadata/config/ /config/

CMD ["tradesBlockService"]
