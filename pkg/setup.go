package pkg

import (
	"gin-blog-newest/config"
	"gin-blog-newest/interval/database"
	"gin-blog-newest/interval/middleware"
	"gin-blog-newest/interval/router"
	"gin-blog-newest/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetUp() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.ApiLogger())
	serverConf := config.NewConfig()
	db := database.InitMysql(serverConf)
	logger.Info().Msg("server is running")
	router.InitRouter(r, db, logger.Get())
	r.Run("localhost:8080")
	//err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", serverConf.Server.Port), r)
	//if err != nil {
	//	return
	//}
}
