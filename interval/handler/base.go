package handler

import (
	"gin-blog-newest/interval/model"
	"gin-blog-newest/interval/service"
	"gin-blog-newest/pkg/logger"
	"gin-blog-newest/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BaseHandler[T any] interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type BaseHandlerImpl[T any] struct {
	BaseService service.BaseService[T]
	response    response.Response
	log         *logger.Logger
}

func NewBaseHandler[T any](baseService service.BaseService[T], logger *logger.Logger) BaseHandler[T] {
	return &BaseHandlerImpl[T]{
		BaseService: baseService,
		log:         logger,
	}
}

func (h *BaseHandlerImpl[T]) Create(ctx *gin.Context) {
	var data T
	h.log.Info().Msg("create user dsfjkshfjksha")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		h.response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	h.log.Info().Msg("create user")
	if err := h.BaseService.Create(&data); err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *BaseHandlerImpl[T]) Update(ctx *gin.Context) {
	var data model.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		h.response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
func (h *BaseHandlerImpl[T]) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := h.BaseService.Delete(id); err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *BaseHandlerImpl[T]) FindByID(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	data, err := h.BaseService.FindByID(id)
	if err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.response.SuccessResponse(ctx, data)
}

func (h *BaseHandlerImpl[T]) FindAll(ctx *gin.Context) {
	data, err := h.BaseService.FindAll()
	if err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.response.SuccessResponse(ctx, data)
}
