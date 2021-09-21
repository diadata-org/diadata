FROM us.icr.io/dia-registry/devops/build:latest as build

WORKDIR $GOPATH

WORKDIR $GOPATH/src/
COPY ./cmd/blockchain/ethereum/oracleService ./

RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/oracleService /bin/oracleService
COPY --from=build /config/ /config/

CMD ["oracleService"]