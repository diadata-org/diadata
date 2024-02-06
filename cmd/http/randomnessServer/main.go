package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/diadata-org/diadata/pkg/utils"
)

var ransomnessSigner *utils.RandomnessSigner

type DrandApiResponse struct {
	Randomness        string `json:"randomness"`
	Round             int    `json:"round"`
	Signature         string `json:"signature"`
	PreviousSignature string `json:"previous_signature"`
}

type ApiResponse struct {
	ResultFrom   string `json:"result_from"`
	DiaSignature string `json:"dia_signature"`

	DrandApiResponse
}

func callAPI(url string, wg *sync.WaitGroup, resultChan chan<- *ApiResponse, errChan chan<- error) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errChan <- fmt.Errorf("non-OK HTTP status from %s: %d", url, resp.StatusCode)
		return
	}

	var drandResult DrandApiResponse
	var result ApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&drandResult); err != nil {
		errChan <- err
		log.Println("err", err)
		return
	}
	result.PreviousSignature = drandResult.PreviousSignature
	result.Signature = drandResult.Signature
	result.Randomness = drandResult.Randomness
	result.Round = drandResult.Round

	result.ResultFrom = url

	signedData, err := ransomnessSigner.Sign(result.Randomness, result.Round)
	if err != nil {
		errChan <- err
		return
	}

	result.DiaSignature = signedData

	resultChan <- &result
}

func fetchFirstResponse(baseURLs []string, number int) (*ApiResponse, error) {
	var wg sync.WaitGroup
	resultChan := make(chan *ApiResponse, 1)
	errChan := make(chan error, len(baseURLs))

	for _, baseURL := range baseURLs {
		url := fmt.Sprintf("%s/public/%d", strings.TrimSpace(baseURL), number)
		wg.Add(1)
		go callAPI(url, &wg, resultChan, errChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	for {
		select {
		case result := <-resultChan:
			return result, nil
		case err := <-errChan:
			if len(errChan) == 0 {
				return nil, err
			}
		}
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	baseUrlsStr := os.Getenv("BASE_URL")
	if baseUrlsStr == "" {
		baseUrlsStr = "https://api.drand.sh,https://api2.drand.sh,https://drand.cloudflare.com"
	}
	baseURLs := strings.Split(baseUrlsStr, ",")

	numberStr := r.URL.Query().Get("round")
	if numberStr == "" {
		http.Error(w, "Round parameter is missing", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Round parameter must be an integer", http.StatusBadRequest)
		return
	}

	result, err := fetchFirstResponse(baseURLs, number)
	if err != nil {
		log.Println("error", err)
		http.Error(w, "Failed to fetch data from APIs", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func main() {
	signerKey := os.Getenv("SIGNER_KEY")
	ransomnessSigner = utils.NewRandomnessSigner(signerKey)
	http.HandleFunc("/randomness", apiHandler)
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
