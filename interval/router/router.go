package router

import (
	"gin-blog-newest/interval/handler"
	"gin-blog-newest/interval/repository"
	"gin-blog-newest/interval/service"
	"gin-blog-newest/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB, log *logger.Logger) {
	api := r.Group("/api/v1")
	{
		// 初始化user路由
		InitUserRouter(api, db, log)
	}

}
func InitUserRouter(api *gin.RouterGroup, db *gorm.DB, log *logger.Logger) {
	userRepo := repository.NewUserRepository(db, log)
	userService := service.NewUserService(userRepo, log)
	userHandler := handler.NewUserHandler(userService, log)
	api.POST("/users/register", userHandler.Create)
}
