package diaApi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/gin-gonic/gin"
)

type stubStore struct {
	models.Datastore
	fq    models.ForeignQuotation
	fqErr error
	aq    []models.AssetQuotation
	aqErr error
}

func (s stubStore) GetForeignQuotationInflux(symbol, source string, timestamp time.Time) (models.ForeignQuotation, error) {
	// Zero value with nil error mirrors the datastore behavior for times before data exists.
	return s.fq, s.fqErr
}

func (s stubStore) GetAssetQuotations(asset dia.Asset, starttime time.Time, endtime time.Time) ([]models.AssetQuotation, error) {
	return s.aq, s.aqErr
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

const (
	fqRoute = "/foreignQuotation/:source/:symbol"
	auRoute = "/assetUpdates/:blockchain/:address/:deviation/:frequencySeconds"
	auURL   = "/assetUpdates/Ethereum/0x0000000000000000000000000000000000000000/20/120"
)

func TestForeignQuotationStatusCodes(t *testing.T) {
	backendErr := errors.New("influx unavailable")
	cases := []struct {
		name, target string
		store        stubStore
		want         int
	}{
		{"empty data before history -> 404", "/foreignQuotation/Coingecko/BTC?time=1647349656", stubStore{}, 404},
		{"backend error -> 500", "/foreignQuotation/Coingecko/BTC?time=1647349656", stubStore{fqErr: backendErr}, 500},
		{"yahoo reversed backend error -> 500", "/foreignQuotation/YahooFinance/EUR-USD?time=1647349656", stubStore{fqErr: backendErr}, 500},
		{"yahoo reversed no history keeps 200", "/foreignQuotation/YahooFinance/EUR-USD?time=1647349656", stubStore{}, 200},
		{"bad time -> 400", "/foreignQuotation/Coingecko/BTC?time=abc", stubStore{}, 400},
		{"special chars in query -> 400", "/foreignQuotation/Coingecko/BTC?time=1%27", stubStore{}, 400},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			env := &Env{DataStore: c.store}
			code, body := doRequest(t, fqRoute, c.target, env.GetForeignQuotation)
			if code != c.want {
				t.Fatalf("status = %d, want %d (%v)", code, c.want, body)
			}
			if c.want >= 400 {
				if _, ok := body["errormessage"]; !ok {
					t.Fatalf("missing JSON error message: %v", body)
				}
				if int(body["errorcode"].(float64)) != c.want {
					t.Fatalf("errorcode = %v, want %d", body["errorcode"], c.want)
				}
			}
		})
	}
}

// assetUpdates needs a RelDB; GetAsset is served from the redis cache, so a throwaway redis
// (REDISURL) is enough. Skipped when none is reachable.
func newCachedRelDB(t *testing.T) models.RelDB {
	t.Helper()
	if os.Getenv("REDISURL") == "" {
		t.Skip("REDISURL not set; skipping assetUpdates tests")
	}
	rdb, err := models.NewRelDataStoreWithOptions(false, true)
	if err != nil {
		t.Skipf("no redis: %v", err)
	}
	asset := dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: "Ethereum", Symbol: "TST", Name: "Test"}
	if err := rdb.SetAssetCache(asset); err != nil {
		t.Skipf("no redis: %v", err)
	}
	return *rdb
}

func TestAssetUpdatesStatusCodes(t *testing.T) {
	rdb := newCachedRelDB(t)
	cases := []struct {
		name, target string
		store        stubStore
		want         int
	}{
		{"empty quotations -> 404", auURL, stubStore{}, 404},
		{"backend error -> 500", auURL, stubStore{aqErr: errors.New("influx unavailable")}, 500},
		{"bad starttime -> 400", auURL + "?starttime=x", stubStore{}, 400},
		{"bad deviation -> 400", "/assetUpdates/Ethereum/0x0000000000000000000000000000000000000000/x/120", stubStore{}, 400},
		{"special chars -> 400", auURL + "?starttime=%27", stubStore{}, 400},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			env := &Env{DataStore: c.store, RelDB: rdb}
			code, body := doRequest(t, auRoute, c.target, env.GetAssetUpdates)
			if code != c.want {
				t.Fatalf("status = %d, want %d (%v)", code, c.want, body)
			}
			if _, ok := body["errormessage"]; !ok {
				t.Fatalf("missing JSON error message: %v", body)
			}
		})
	}
}
