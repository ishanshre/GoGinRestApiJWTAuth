package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ishanshre/GoRestApiExample/internals/helper"
	"github.com/ishanshre/GoRestApiExample/internals/models"
	"github.com/ishanshre/GoRestApiExample/internals/repository"
	"github.com/ishanshre/GoRestApiExample/internals/validators"
	"github.com/redis/go-redis/v9"
)

type VideoHandler interface {
	GetAllVideos(ctx *gin.Context)
	CreateVideo(ctx *gin.Context)
	GetVideoByID(ctx *gin.Context)
	DeleteVideoByID(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	GenerateAndSetToken(ctx context.Context, actual_user *models.User) (*models.LoginResponse, *helper.Token, error)
}

type handler struct {
	repo        repository.DatabaseRepo
	redisClient *redis.Client
}

var validate *validator.Validate

func NewRepo(r repository.DatabaseRepo, redisClient *redis.Client) VideoHandler {
	validate = validator.New()
	validate.RegisterValidation("upper", validators.UpperCase)
	validate.RegisterValidation("lower", validators.LowerCase)
	validate.RegisterValidation("number", validators.Number)
	return &handler{
		repo:        r,
		redisClient: redisClient,
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

func (h *handler) RegisterUser(ctx *gin.Context) {
	var user *models.CreateUser
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error in parsing json",
			Data:    err,
		})
		return
	}
	log.Println(user.Email)
	if err := validate.Struct(user); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusBadRequest, helper.Error{
				Message: "Input Validation Error",
				Data:    err.Error(),
			})
			return
		}

		// Handle validation errors
		for _, e := range validationErrors {
			ctx.JSON(http.StatusBadRequest, helper.Error{
				Message: fmt.Sprintf("Validation error for field '%s': %s\n", e.Field(), e.Tag()),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Input Validation Error",
			Data:    err.Error(),
		})
		return
	}
	exists, err := h.repo.UserExists(user.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in validating username",
			Data:    err,
		})
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Username already exists. Please take another username",
		})
		return
	}
	exists, err = h.repo.EmailExists(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in checking email exists",
			Data:    err,
		})
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Email already exists. Please take another username",
		})
		return
	}
	hashedPassword, err := helper.GeneratePassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error hashing password",
			Data:    err,
		})
		return
	}
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	result, err := h.repo.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error registering new user",
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in registering new user",
		Data:    result,
	})
}

func (h *handler) UserLogin(ctx *gin.Context) {
	user := &models.User{}

	// get username and password from request body
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in parsing json",
			Data:    err,
		})
		return
	}

	// check if username exists
	exists, err := h.repo.UserExists(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error in checking the username",
			Data:    err,
		})
		return
	}
	if !exists {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Username does not exists",
		})
		return
	}

	// fetch the user
	actual_user, err := h.repo.GetUserByUsername(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error occured on fetching the user",
			Data:    err,
		})
		return
	}

	// check the password
	if err := helper.ComparePasswordHash(actual_user.Password, user.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Invalid Username/Password",
			Data:    err,
		})
		return
	}
	token, tokenDetail, err := h.GenerateAndSetToken(ctx, actual_user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error in generating and setting token",
			Data:    err,
		})
		return
	}

	log.Println(tokenDetail.AccessToken.ExpiresAt)
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Login Success",
		Data:    token,
	})
}

func (h *handler) RefreshToken(ctx *gin.Context) {
	refreshToken := &models.RefreshToken{}
	if err := ctx.BindJSON(&refreshToken); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Cannot parse the json",
			Data:    err.Error(),
		})
		return
	}
	claims, err := helper.VerifyTokenWithClaims(refreshToken.RefreshToken, "refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in validating token",
			Data:    err.Error(),
		})
		return
	}
	exists, err := h.redisClient.Exists(ctx, claims.TokenID).Result()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Error in checking token in redis",
			Data:    err.Error(),
		})
		return
	}
	if exists == 1 {
		if err := h.redisClient.Del(ctx, claims.TokenID).Err(); err != nil {
			ctx.JSON(http.StatusBadRequest, helper.Error{
				Message: "Error in deleting token",
				Data:    err,
			})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, helper.Error{
			Message: "Token already revoked",
		})
		return
	}
	token, tokenDetail, err := h.GenerateAndSetToken(ctx, &models.User{ID: claims.UserID, Username: claims.Username})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.Error{
			Message: "Error in generating and setting token",
			Data:    err,
		})
		return
	}
	log.Println(tokenDetail.AccessToken.TokenID)
	ctx.JSON(http.StatusOK, helper.Success{
		Message: "Success in refreshing the token",
		Data:    token,
	})
}

func (h *handler) GenerateAndSetToken(ctx context.Context, actual_user *models.User) (*models.LoginResponse, *helper.Token, error) {
	token, tokenDetail, err := helper.GenerateLoginResponse(actual_user.ID, actual_user.Username)
	if err != nil {
		return nil, nil, err
	}
	tokenDetailJSON, err := json.Marshal(tokenDetail)
	if err != nil {
		return nil, nil, err
	}
	if errAccess := h.redisClient.Set(ctx, tokenDetail.AccessToken.TokenID, tokenDetailJSON, time.Until(tokenDetail.AccessToken.ExpiresAt)).Err(); errAccess != nil {
		return nil, nil, errAccess
	}
	return token, tokenDetail, nil
}
