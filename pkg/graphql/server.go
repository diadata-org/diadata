// package main

// import (
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	"github.com/diadata-org/diadata/pkg/graphql/resolver"
// 	models "github.com/diadata-org/diadata/pkg/model"
// 	graphql "github.com/graph-gophers/graphql-go"
// 	"github.com/graph-gophers/graphql-go/relay"
// 	log "github.com/sirupsen/logrus"
// )

// // nolint: gas
// func main() {

// 	ds, err := getSchema("./schema/quotation.graphql")
// 	if err != nil {
// 		panic(err)
// 	}

// 	datastore, err := models.NewDataStore()
// 	if err != nil {
// 		panic(err)
// 	}

// 	diaSchema := graphql.MustParseSchema(ds, &resolver.DiaResolver{DS: *datastore}, graphql.UseStringDescriptions())

// 	mux := http.NewServeMux()

// 	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write(page)
// 	}))

// 	mux.Handle("/query", &relay.Handler{Schema: diaSchema})

// 	log.WithFields(log.Fields{"time": time.Now()}).Info("starting server")
// 	log.Fatal(http.ListenAndServe("localhost:8080", logged(mux)))
// }

// var page = []byte(`
// <!DOCTYPE html>
// <html>
// 	<head>
// 		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
// 		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
// 		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
// 		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
// 		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
// 	</head>
// 	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
// 		<div id="graphiql" style="height: 100vh;">Loading...</div>
// 		<script>
// 			function graphQLFetcher(graphQLParams) {
// 				return fetch("/query", {
// 					method: "post",
// 					body: JSON.stringify(graphQLParams),
// 					credentials: "include",
// 				}).then(function (response) {
// 					return response.text();
// 				}).then(function (responseBody) {
// 					try {
// 						return JSON.parse(responseBody);
// 					} catch (error) {
// 						return responseBody;
// 					}
// 				});
// 			}

// 			ReactDOM.render(
// 				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
// 				document.getElementById("graphiql")
// 			);
// 		</script>
// 	</body>
// </html>
// `)

// func getSchema(path string) (string, error) {
// 	b, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(b), nil
// }

// // logging middleware
// func logged(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now().UTC()

// 		next.ServeHTTP(w, r)

// 		log.WithFields(log.Fields{
// 			"path":    r.RequestURI,
// 			"IP":      r.RemoteAddr,
// 			"elapsed": time.Now().UTC().Sub(start),
// 		}).Info()
// 	})
// }
