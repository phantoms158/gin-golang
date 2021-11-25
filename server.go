package main

import (
	controller "golang-gin/controllers"
	"golang-gin/middlewares"
	"golang-gin/repository"
	"golang-gin/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    services.VideoService      = services.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		// middlewares.BasicAuth(),
		// gindump.Dump(),
	)

	repository.ConnectDataBase()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"messsage": "Video input is valid!!!"})
		}
	})

	server.Run(":8080")
}
