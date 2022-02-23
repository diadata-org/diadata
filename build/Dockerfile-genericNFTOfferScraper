ARG protocol
FROM us.icr.io/dia-registry/devops/build-117:latest as build

WORKDIR $GOPATH/src/

COPY ./cmd/nftOfferscraper ./
RUN go mod tidy && go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/nftOfferscraper /bin/nftOfferscraper

CMD ["nftOfferscraper"]