package routes

import (
	"go-crud/app/controllers"

	"github.com/gin-gonic/gin"
)

func MemberRoutes(engine *gin.Engine) {
	engine.GET("/members", controllers.FindMembers)
	engine.POST("/members", controllers.CreateMember)
	engine.GET("/members/:id", controllers.FindMember)
	engine.PATCH("/members/:id", controllers.UpdateMember)
	engine.DELETE("members/:id", controllers.DeleteMember)
}
