package main

import (
	"Gin_Studies/Service"
	"Gin_Studies/controller"
	"Gin_Studies/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var (
	videoService    Service.VideoService       = Service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLogOutput() { //Creating Log file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setUpLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Video Page",
			})
		})
	}

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
