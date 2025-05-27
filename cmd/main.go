package main

import (
	"gin-blog-newest/pkg"
	"gin-blog-newest/pkg/logger"
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
	logger.Info().Msg("Starting server...")
	pkg.SetUp()
}
