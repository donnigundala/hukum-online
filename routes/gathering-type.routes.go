package routes

import (
	"go-crud/app/controllers"

	"github.com/gin-gonic/gin"
)

func GatheriingTypeRoutes(engine *gin.Engine) {
	engine.GET("/gathering-types", controllers.FindGatheringTypes)
	engine.POST("/gathering-types", controllers.CreateGatheringType)
	engine.GET("/gathering-types/:id", controllers.FindGatheringType)
	engine.PATCH("/gathering-types/:id", controllers.UpdateGatheringType)
	engine.DELETE("gathering-types/:id", controllers.DeleteGatheringType)
}
