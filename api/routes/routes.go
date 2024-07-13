package routes

import (
	"github.com/devops-ferraz/the-cats/api/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	exampleController := controllers.NewExampleController()
	v1 := router.Group("/v1")
	{
		//EXAMPLE
		v1.POST("examples", exampleController.CreateExample)
		v1.GET("examples", exampleController.FindAllExamples)
		v1.GET("examples/current", exampleController.GetCurrentValue)
	}

	return v1
}
