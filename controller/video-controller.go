package controller

import (
	"Gin_Studies/Service"
	"Gin_Studies/entity"
	"Gin_Studies/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
	ShowAll(c *gin.Context)
}

type controller struct {
	service Service.VideoService
}

var validate *validator.Validate

func New(service Service.VideoService) VideoController {
	validate = validator.New()
	err := validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	if err != nil {
		return nil
	}
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

//func (c *controller) Save(gc *gin.Context) entity.Video {
//	var video entity.Video
//	err := gc.BindJSON(&video)
//	if err != nil {
//		return entity.Video{}
//	}
//	c.service.Save(video)
//	return video
//}

func (c *controller) Save(gc *gin.Context) error {
	var video entity.Video
	err := gc.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(gc *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	gc.HTML(http.StatusOK, "index.html", data)
}
