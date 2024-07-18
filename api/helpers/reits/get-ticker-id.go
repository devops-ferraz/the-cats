package reitsHelpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ExtractFiiId(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch the page: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	fiiID := extractFiiIdFromBody(string(body))
	if fiiID == 0 {
		return 0, fmt.Errorf("fii_id not found in the page content")
	}

	return fiiID, nil
}

func extractFiiIdFromBody(html string) int {
	const prefix = "fii_id: "
	index := strings.Index(html, prefix)
	if index == -1 {
		return 0
	}

	start := index + len(prefix)
	end := start
	for end < len(html) && html[end] >= '0' && html[end] <= '9' {
		end++
	}

	fiiIDStr := html[start:end]
	fiiID, err := strconv.Atoi(fiiIDStr)
	if err != nil {
		return 0
	}

	return fiiID
}
