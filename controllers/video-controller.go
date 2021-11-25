package controller

import (
	"golang-gin/entities"
	"golang-gin/services"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entities.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entities.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
