FROM golang:latest as build

WORKDIR $GOPATH/src/

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/cmd/services/pairDiscoveryService

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/pairDiscoveryService /bin/pairDiscoveryService
COPY --from=build /go/src/github.com/diadata-org/diadata/config/ /config/

CMD ["pairDiscoveryService"]
