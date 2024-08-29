package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-example/config"
)

func main() {
	config.LoadEnv()
	appConfig := config.GetAppConfig()
	_, errDb := config.ConnectDB(&appConfig)
	if errDb != nil {
		return
	}
	fmt.Println("Hello, World!")

	route := gin.Default()

	route.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, "OK")
	})
	errRun := route.Run(`:3001`)
	if errRun != nil {
		return
	}
}
