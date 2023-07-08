package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func NewRepo(r repository.VideoService) VideoHandler {
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
	}
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in fetching the videos information",
		Data:    videos,
	})
}

func (h *handler) Save(ctx *gin.Context) {
	var video models.Video
	ctx.BindJSON(&video)
	video, err := h.repo.Save(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error saving video information",
		})
	}
	ctx.JSON(200, helper.Success{
		Message: "Video Info Successfull Saved",
		Data:    video,
	})
}
