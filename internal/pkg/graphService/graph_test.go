package graph

import (
	"github.com/ethereum/go-ethereum/log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	RAND_SEED      = 543785798
	WEEK_NUMPOINTS = 30 * 24 * 7
	TEST_DIR       = "./testData/"
	CREATED        = "test.png"
	EXPECTED       = "test_golden.png"
)

func TestPriceGraph(t *testing.T) {
	numPoints := WEEK_NUMPOINTS
	prices := make([]float64, numPoints)
	times := make([]int64, numPoints)

	// generate data
	rand.Seed(RAND_SEED)
	var x = 0.0
	for i := 0; i < len(prices); i++ {
		prices[i] = math.Sin(float64(rand.Float64()))
		x += 0.1
	}

	weekAgo := time.Now().Add(-time.Hour * 7 * 24)
	for i := range times {
		times[i] = weekAgo.Add(time.Minute * time.Duration(2*i)).Unix()
	}

	// create new graph
	err := PriceGraph(prices, times, TEST_DIR+CREATED)
	if err != nil {
		t.Fatal("Failed to create graph:", err)
	}

	// check if images are present
	_, err = os.Stat(filepath.FromSlash(TEST_DIR + EXPECTED))
	if err != nil {
		t.Fatal("Failed to access golden image:", err)
	}

	_, err = os.Stat(filepath.FromSlash(TEST_DIR + CREATED))
	if err != nil {
		t.Fatal("Failed to access created image:", err)
	}

	log.Info("Compare images '" + TEST_DIR + EXPECTED + "' (expected) and '" + TEST_DIR + CREATED + "' (created) manually")
}
