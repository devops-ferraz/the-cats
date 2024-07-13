package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/devops-ferraz/the-cats/api/helpers/reits"
	"github.com/devops-ferraz/the-cats/api/structs/examples"
	"github.com/gin-gonic/gin"
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
	ticker := ctx.Query("ticker")
	if ticker == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ticker parameter is required"})
		return
	}
	currentValue, err := reits.GetCurrentValue(ticker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"currentValue": currentValue})
}
