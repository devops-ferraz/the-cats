package reitsHelpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetCurrentValue(name string) (float64, error) {
	quote := os.Getenv("CURRENT_QUOTE")
	apiURL := fmt.Sprintf(quote, name)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	content := data["content"].([]interface{})
	html := content[0].(string)

	re := regexp.MustCompile(`class="price "><span>R\$ ([\d,]+)<`)
	match := re.FindStringSubmatch(html)
	if match == nil {
		return 0, fmt.Errorf("value not found")
	}

	valueStr := strings.ReplaceAll(match[1], ",", ".")
	currentValue, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, err
	}

	return currentValue, nil
}
