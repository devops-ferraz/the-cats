package reitsHelpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetLastDividend(ticker string) (float64, error) {
	baseURL := os.Getenv("INVESTIDOR10_BASE_URL")
	if baseURL == "" {
		return 0, fmt.Errorf("base URL is not configured")
	}

	url := fmt.Sprintf("%s/%s/%s/", baseURL, "fiis", ticker)
	tickerId, err := ExtractFiiId(url)
	if err != nil {
		return 0, err
	}

	history := os.Getenv("HISTORY_YIELDS")
	apiURL := fmt.Sprintf(history, strconv.Itoa(tickerId))
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data []struct {
		LastPay float64 `json:"price"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, fmt.Errorf("no dividend found")
	}

	lastDividend := data[len(data)-1].LastPay
	return lastDividend, nil
}
