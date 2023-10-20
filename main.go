package main

import (
	"go-crud/app/models"
	"go-crud/configs"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin engine
	engine := gin.Default()

	// run db setup
	initDb(engine)

	// run registered routes
	registerRoutes(engine)

	// run gin
	engine.Run()
}

func initDb(engine *gin.Engine) {
	db := configs.Db()
	db.AutoMigrate(&models.Member{})
	db.AutoMigrate(&models.GatheringType{})
	db.AutoMigrate(&models.Gathering{})
	db.AutoMigrate(&models.InvitationStatus{})
	db.AutoMigrate(&models.Invitation{})

	engine.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
}

func registerRoutes(gin *gin.Engine) {
	routes.BasicRoutes(gin)
	routes.MemberRoutes(gin)
	routes.GatheriingTypeRoutes(gin)
	routes.GatheriingRoutes(gin)
}
