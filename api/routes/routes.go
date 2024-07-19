package routes

import (
	"github.com/devops-ferraz/the-cats/api/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	exampleController := controllers.NewExampleController()
	reitSearchFilterController := controllers.NewSearchFilterController()
	v1 := router.Group("/v1")
	{
		//EXAMPLE
		v1.POST("examples", exampleController.CreateExample)
		v1.GET("examples", exampleController.FindAllExamples)
		v1.GET("examples/current/:ticker", exampleController.GetCurrentValue)
		v1.GET("examples/ticker-id/:type/:ticker", exampleController.GetTickerId)
		v1.GET("examples/get-average-pvp", exampleController.GetAveragePVP)
		v1.POST("examples/calculator", exampleController.ReitCalculator)
		v1.GET("examples/properties/:ticker", exampleController.FindReitProperty)

		//REITS
		v1.GET("reits/search", reitSearchFilterController.SearchHandler)
		v1.GET("reits/update-data-reits", reitSearchFilterController.UpdateDataReits)
	}

	return v1
}
