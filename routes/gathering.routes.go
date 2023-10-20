package routes

import (
	"go-crud/app/controllers"

	"github.com/gin-gonic/gin"
)

func GatheriingRoutes(engine *gin.Engine) {
	engine.GET("/gatherings", controllers.FindGatherings)
	engine.POST("/gatherings", controllers.CreateGathering)
	engine.GET("/gatherings/:id", controllers.FindGathering)
	engine.PATCH("/gatherings/:id", controllers.UpdateGathering)
	engine.DELETE("gatherings/:id", controllers.DeleteGathering)
}
