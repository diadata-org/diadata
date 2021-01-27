FROM golang:1.14 as build

WORKDIR $GOPATH/src/

COPY . .

WORKDIR $GOPATH/src/github.com/diadata-org/diadata/cmd/services/githubScraper
RUN go install

FROM gcr.io/distroless/base

COPY --from=build /go/bin/githubScraper /bin/githubScraper
COPY --from=build /go/src/github.com/diadata-org/diadata/config /config/

CMD ["githubScraper"]