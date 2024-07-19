package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	reitsHelpers "github.com/devops-ferraz/the-cats/api/helpers/reits"
	"github.com/devops-ferraz/the-cats/api/structs/reits"

	"github.com/gin-gonic/gin"
)

type searchFilter struct {
}

func NewSearchFilterController() *searchFilter {
	return &searchFilter{}
}

func (t *searchFilter) SearchHandler(ctx *gin.Context) {
	baseURL := os.Getenv("BASE_URL")

	queryParams := url.Values{}
	for key, values := range ctx.Request.URL.Query() {
		for _, value := range values {
			queryParams.Add(key, value)
		}
	}

	finalURL := baseURL + "?" + queryParams.Encode()

	req, _ := http.NewRequest("GET", finalURL, nil)
	req.Header.Add("User-Agent", `insomnia/9.3.2`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	var response struct {
		SearchFilters []reits.SearchFilter `json:"list"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *searchFilter) UpdateDataReits(ctx *gin.Context) {
	filePath := ctx.Query("filePath")
	sheetName := ctx.Query("sheetName")

	if filePath == "" || sheetName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file path and sheet name are required",
		})
		return
	}

	response := reitsHelpers.UpdateDataReits(filePath, sheetName)

	ctx.JSON(http.StatusOK, response)
}
