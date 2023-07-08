package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ishanshre/GoRestApiExample/internals/helper"
	"github.com/ishanshre/GoRestApiExample/internals/models"
	"github.com/ishanshre/GoRestApiExample/internals/repository"
)

type VideoHandler interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type handler struct {
	repo repository.VideoService
}

var validate *validator.Validate

func NewRepo(r repository.VideoService) VideoHandler {
	validate = validator.New()
	return &handler{
		repo: r,
	}
}

func (h *handler) FindAll(ctx *gin.Context) {
	videos, err := h.repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error in fetching all video informations",
		})
		return
	}
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in fetching the videos information",
		Data:    videos,
	})
}

func (h *handler) Save(ctx *gin.Context) {
	var video models.Video
	if err := ctx.BindJSON(&video); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: err.Error(),
		})
		return
	}
	if err := validate.Struct(video); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: err.Error(),
		})
	}
	video, err := h.repo.Save(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error saving video information",
		})
		return
	}
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Video Info Successfull Saved",
		Data:    video,
	})
}
