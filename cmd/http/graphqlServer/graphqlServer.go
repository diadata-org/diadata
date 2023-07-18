package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/graphql/resolver"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"
)

// nolint: gas
func main() {

	ds, err := getSchema("./schema/quotation.graphql")
	if err != nil {
		panic(err)
	}

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	relStore, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}

	batchSizeString := utils.Getenv("BATCH_SIZE_INFLUX", "50")
	influxBatchSize, err := strconv.ParseInt(batchSizeString, 10, 64)
	if err != nil {
		log.Fatal("parse batch duration ", err)
	}

	diaSchema := graphql.MustParseSchema(ds, &resolver.DiaResolver{DS: *datastore, RelDB: *relStore, InfluxBatchSize: influxBatchSize}, graphql.UseFieldResolvers())

	mux := http.NewServeMux()
	urlFolderPrefix := utils.Getenv("URL_FOLDER_PREFIX", "/graphql")
	if !strings.HasPrefix(urlFolderPrefix, "/") {
		urlFolderPrefix = "/" + urlFolderPrefix
	}
	mux.Handle(urlFolderPrefix+"/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(page)
		if err != nil {
			return
		}
	}))

	mux.Handle(urlFolderPrefix+"/query", &relay.Handler{Schema: diaSchema})

	log.WithFields(log.Fields{"time": time.Now()}).Info("starting server")
	log.Fatal(http.ListenAndServe(utils.Getenv("LISTEN_PORT", ":1111"), logged(mux)))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/test/graphql/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// logging middleware
func logged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"path":    r.RequestURI,
			"IP":      r.RemoteAddr,
			"elapsed": time.Now().UTC().Sub(start),
		}).Info()
	})
}
