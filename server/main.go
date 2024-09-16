package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pfedgarmauricio/go-caloric-calc/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntry)

	router.PUT("/entry/:id/update", routes.UpdateEntry)
	router.DELETE("/entry/:id", routes.DeleteEntry)

	router.Run(":" + port)
}
