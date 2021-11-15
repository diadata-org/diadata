FROM us.icr.io/dia-registry/devops/build:latest as build

WORKDIR $GOPATH/src/

COPY ./cmd/influxMigration/migration ./
RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/influxMigration /bin/influxMigration

CMD ["influxMigration"]