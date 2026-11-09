package diaApi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/gin-gonic/gin"
)

type stubStore struct {
	models.Datastore
}

func (stubStore) GetForeignQuotationInflux(symbol, source string, timestamp time.Time) (models.ForeignQuotation, error) {
	// Mirrors the datastore behavior for times before data exists: empty result, nil error.
	return models.ForeignQuotation{}, nil
}

func doRequest(t *testing.T, route, target string, h gin.HandlerFunc) (int, map[string]interface{}) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET(route, h)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, target, nil))
	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %q", w.Body.String())
	}
	return w.Code, body
}

func TestHistoricQueriesReturn4xxJSON(t *testing.T) {
	env := &Env{DataStore: stubStore{}}
	cases := []struct {
		name, route, target string
		h                   gin.HandlerFunc
		want                int
	}{
		{"foreignQuotation before history", "/foreignQuotation/:source/:symbol", "/foreignQuotation/Coingecko/BTC?time=1647349656", env.GetForeignQuotation, 404},
		{"foreignQuotation bad time", "/foreignQuotation/:source/:symbol", "/foreignQuotation/Coingecko/BTC?time=abc", env.GetForeignQuotation, 400},
		{"assetUpdates bad starttime", "/assetUpdates/:blockchain/:address/:deviation/:frequencySeconds", "/assetUpdates/Ethereum/0x0000000000000000000000000000000000000000/20/120?starttime=x", env.GetAssetUpdates, 400},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			code, body := doRequest(t, c.route, c.target, c.h)
			if code != c.want {
				t.Fatalf("status = %d, want %d (%v)", code, c.want, body)
			}
			if _, ok := body["errormessage"]; !ok {
				t.Fatalf("missing JSON error message: %v", body)
			}
		})
	}
}
