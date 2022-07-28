FROM us.icr.io/dia-registry/devops/build:latest as build

WORKDIR $GOPATH/src/

COPY ./cmd/influxMigration/indexReplay ./
RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/indexReplay /bin/indexReplay

CMD ["indexReplay"]