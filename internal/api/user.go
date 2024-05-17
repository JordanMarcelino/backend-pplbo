package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/backend-pplbo/internal/models"
	"github.com/jordanmarcelino/backend-pplbo/internal/usecase"
	"github.com/jordanmarcelino/backend-pplbo/internal/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserHandler(log *logrus.Logger, useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{Log: log, UseCase: useCase}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	request := new(models.UserLogin)

	if err := ctx.ShouldBindJSON(request); err != nil {
		h.Log.Warnf("failed to bind request : %+v", err)
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("bad request"))
		return
	}

	err := h.UseCase.Login(ctx, request)
	if err != nil {
		h.Log.Warnf("failed to login : %+v", err)
		ctx.JSON(util.CheckError(err), models.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse[*string](nil, "success login"))
}

func (h *UserHandler) Register(ctx *gin.Context) {
	request := new(models.UserRegister)

	if err := ctx.ShouldBindJSON(request); err != nil {
		h.Log.Warnf("failed to bind request : %+v", err)
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("bad request"))
		return
	}

	response, err := h.UseCase.Create(ctx, request)
	if err != nil {
		h.Log.Warnf("failed to register : %+v", err)
		ctx.JSON(util.CheckError(err), models.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, models.NewSuccessResponse[models.UserResponse](*response, "success register"))
}

func (h *UserHandler) Get(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.Log.Warnf("failed to convert id : %+v", err)
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("invalid id"))
		return
	}

	response, err := h.UseCase.FindById(ctx, userID)
	if err != nil {
		h.Log.Warnf("failed to get user : %+v", err)
		ctx.JSON(util.CheckError(err), models.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse[models.UserResponse](*response, "success get user"))
}
