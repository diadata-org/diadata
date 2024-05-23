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

	withInflux, err := strconv.ParseBool(utils.Getenv("WITH_INFLUX", "false"))
	if err != nil {
		log.Fatal("parse WITH_INFLUX: ", err)
	}

	diaSchema := graphql.MustParseSchema(
		ds,
		&resolver.DiaResolver{DS: *datastore, RelDB: *relStore, InfluxBatchSize: influxBatchSize, WithInflux: withInflux},
		graphql.UseFieldResolvers(),
	)

	mux := http.NewServeMux()
	urlFolderPrefix := utils.Getenv("URL_FOLDER_PREFIX", "/graphql")
	if !strings.HasPrefix(urlFolderPrefix, "/") {
		urlFolderPrefix = "/" + urlFolderPrefix
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(page)
		if err != nil {
			return
		}
	})
	mux.Handle(urlFolderPrefix+"/", handler)
	mux.Handle(urlFolderPrefix+"/query", &relay.Handler{Schema: diaSchema})

	readTimeout, err := strconv.Atoi(utils.Getenv("READ_TIMEOUT", "10"))
	writeTimeout, err := strconv.Atoi(utils.Getenv("WRITE_TIMEOUT", "10"))
	idleTimeout, err := strconv.Atoi(utils.Getenv("IDLE_TIMEOUT", "10"))
	if err != nil {
		log.Error("parse timeouts: ", err)
	}

	log.Infof("readTimeout - writeTimeout - idleTimeout: %v - %v - %v", readTimeout, writeTimeout, idleTimeout)
	srv := &http.Server{
		Addr:         utils.Getenv("LISTEN_PORT", ":1111"),
		Handler:      logged(mux),
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	log.WithFields(log.Fields{"time": time.Now()}).Info("starting server")
	log.Fatal(srv.ListenAndServe())

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
				return fetch("/graphql/query", {
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

// responseRecorder is used to capture the status code for logging
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

// enhanced logging middleware
func logged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		log.Infof("Started %s %s", r.Method, r.URL.Path)
		// Log request details
		log.WithFields(log.Fields{
			"method":  r.Method,
			"url":     r.URL.Path,
			"headers": r.Header,
		}).Info("Request details")
		// Capture the response details
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)
		// Log response details
		log.WithFields(log.Fields{
			"path":    r.RequestURI,
			"IP":      r.RemoteAddr,
			"status":  rr.statusCode,
			"elapsed": time.Since(start),
		}).Info("Completed request")
	})
}
