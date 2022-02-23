FROM us.icr.io/dia-registry/devops/build:latest as build

WORKDIR $GOPATH/src/

COPY ./cmd/benchmarkedIndex ./
RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/benchmarkedIndex /bin/benchmarkedIndex

CMD ["benchmarkedIndex"]