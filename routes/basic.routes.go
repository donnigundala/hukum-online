package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello Worldsss!"})
	})
}
