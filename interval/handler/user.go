package handler

import (
	"gin-blog-newest/interval/model"
	"gin-blog-newest/interval/service"
	"gin-blog-newest/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	BaseHandler[model.User]
	GetByUsername(ctx *gin.Context)
	GetByEmail(ctx *gin.Context)
	GetByPhone(ctx *gin.Context)
}

type UserHandlerImpl struct {
	BaseHandlerImpl[model.User]
	us service.UserService
}

func NewUserHandler(us service.UserService, logger *logger.Logger) UserHandler {
	return &UserHandlerImpl{
		us: us,
		BaseHandlerImpl: BaseHandlerImpl[model.User]{
			BaseService: us,
			log:         logger,
		},
	}
}

func (h *UserHandlerImpl) GetByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := h.us.FindByUsername(username)
	h.log.Info().Msg("get user by username")
	if err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.response.SuccessResponse(ctx, user)
}
func (h *UserHandlerImpl) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := h.us.FindByEmail(email)
	if err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.response.SuccessResponse(ctx, user)
}
func (h *UserHandlerImpl) GetByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")
	user, err := h.us.FindByPhone(phone)
	if err != nil {
		h.response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.response.SuccessResponse(ctx, user)

}
