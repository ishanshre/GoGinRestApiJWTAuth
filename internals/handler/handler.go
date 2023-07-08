package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ishanshre/GoRestApiExample/internals/helper"
	"github.com/ishanshre/GoRestApiExample/internals/models"
	"github.com/ishanshre/GoRestApiExample/internals/repository"
)

type VideoHandler interface {
	GetAllVideos(ctx *gin.Context)
	CreateVideo(ctx *gin.Context)
	GetVideoByID(ctx *gin.Context)
	DeleteVideoByID(ctx *gin.Context)
}

type handler struct {
	repo repository.DatabaseRepo
}

var validate *validator.Validate

func NewRepo(r repository.DatabaseRepo) VideoHandler {
	validate = validator.New()
	return &handler{
		repo: r,
	}
}

func (h *handler) GetAllVideos(ctx *gin.Context) {
	videos, err := h.repo.GetAllVideos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error in fetching all video informations",
		})
		return
	}

	for i := range videos {
		author, err := h.repo.GetAuhtorByID(videos[i].Author.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.Error{
				Message: "Error in fetching author detail",
			})
			return
		}

		videos[i].Author = *author
	}
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in fetching the videos information",
		Data:    videos,
	})
}

func (h *handler) CreateVideo(ctx *gin.Context) {
	var video *models.Video
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
	result, err := h.repo.CreateVideo(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error saving video information",
			Data:    err,
		})
		return
	}
	author, err := h.repo.GetAuhtorByID(video.Author.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: err.Error(),
		})
		return
	}
	result.Author = *author
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Video Info Successfull Saved",
		Data:    result,
	})
}

func (h *handler) GetVideoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in parsing the id",
		})
		return
	}
	video, err := h.repo.GetVideoByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: err.Error(),
		})
		return
	}
	author, err := h.repo.GetAuhtorByID(video.Author.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: err.Error(),
		})
		return
	}
	video.Author = *author
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in fetching the video",
		Data:    video,
	})

}

func (h *handler) DeleteVideoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in parsing the id",
		})
		return
	}
	if err := h.repo.DeleteVideoByID(id); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in deleting the video",
	})
}
