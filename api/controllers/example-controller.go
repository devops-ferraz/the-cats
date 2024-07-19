package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	reitsHelpers "github.com/devops-ferraz/the-cats/api/helpers/reits"
	examples "github.com/devops-ferraz/the-cats/api/structs/examples"
	reits "github.com/devops-ferraz/the-cats/api/structs/reits"
	gin "github.com/gin-gonic/gin"
)

type example struct {
	examples []examples.Example
}

func NewExampleController() *example {
	return &example{}
}
func (ctrl *example) CreateExample(c *gin.Context) {
	tickerList := os.Getenv("THE_CAT_TICKER_LIST")
	if tickerList == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The environment variable THE_CAT_TICKER_LIST is not set"})
		return
	}
	tickerSymbols := strings.Split(tickerList, ",")
	if len(tickerSymbols) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ticker list is empty"})
		return
	}
	newExample := examples.NewExample()
	if err := c.BindJSON(&newExample); err != nil {
		return
	}
	ctrl.examples = append(ctrl.examples, *newExample)

	c.JSON(http.StatusOK, newExample)
}
func (t *example) FindAllExamples(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.examples)
}
func (t *example) GetCurrentValue(ctx *gin.Context) {
	ticker := ctx.Param("ticker")
	if ticker == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ticker parameter is required"})
		return
	}
	currentValue, err := reitsHelpers.GetCurrentValue(ticker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"currentValue": currentValue})
}
func (t *example) GetTickerId(ctx *gin.Context) {
	assetType := ctx.Param("type")
	ticker := ctx.Param("ticker")
	if assetType == "" || ticker == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "type and ticker parameters are required"})
		return
	}

	baseURL := os.Getenv("INVESTIDOR10_BASE_URL")
	if baseURL == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Base URL is not configured"})
		return
	}

	url := fmt.Sprintf("%s/%s/%s/", baseURL, assetType, ticker)
	tickerId, err := reitsHelpers.ExtractFiiId(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tickerId": tickerId})
}
func (t *example) GetAveragePVP(ctx *gin.Context) {
	tickers := os.Getenv("THE_CAT_TICKER_LIST")
	tickerSymbols := strings.Split(tickers, ",")

	results := make(map[string]interface{})

	for _, tickerSymbol := range tickerSymbols {
		result, err := reitsHelpers.CalculatePVP(tickerSymbol)
		if err != nil {
			results[tickerSymbol] = gin.H{"error": err.Error()}
		} else {
			results[tickerSymbol] = result
		}
	}

	ctx.JSON(http.StatusOK, results)
}
func (t *example) ReitCalculator(ctx *gin.Context) {
	var fund reits.RealEstateFund

	log.Println("Starting ReitCalculator")

	if err := ctx.ShouldBindJSON(&fund); err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	evolution, err := reitsHelpers.CalculateEvolution(fund)
	if err != nil {
		log.Printf("Error calculating evolution: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"evolution": evolution,
	})

	log.Println("Finished ReitCalculator")
}
func (t *example) FindReitProperty(ctx *gin.Context) {
	baseURL := os.Getenv("REIT_PROPERTIE_URL")
	ticker := ctx.Param("ticker")
	finalURL := fmt.Sprintf(baseURL, ticker)

	resp, err := http.Get(finalURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	var response reits.ReitProperty
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
