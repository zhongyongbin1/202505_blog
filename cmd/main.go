package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-blog-newest/config"
	"gin-blog-newest/interval/database"
	"gin-blog-newest/interval/middleware"
	"gin-blog-newest/interval/router"
	"gin-blog-newest/pkg/logger"

	"github.com/gin-gonic/gin"
)

// @title           Gin Blog API
// @version         1.0
// @description     这是一个基于Gin框架的博客API服务
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	logger.InitLogger()
	log := logger.Get()
	log.Info().Msg("server is running")
	conf := config.NewConfig()
	if conf.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化数据库
	db := database.InitMysql(conf)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.ApiLogger())
	router.InitRouter(r, db, log)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Server.Port),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("listen and serve error")
		}
	}()
	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("server is shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("server shutdown error")
	}
	log.Info().Msg("server is exiting")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}
	log.Info().Msg("server exiting")
}
