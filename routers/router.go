package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/knight0zh/demo_config/config"
	"github.com/knight0zh/demo_server/middlewares"
	"github.com/knight0zh/demo_server/routers/api/demo"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(middlewares.Logger(config.AccessLogger), middlewares.Auth())
	r.Use(gin.Recovery())

	r.GET("/hello/world", demo.HelloWorld)
	return r
}
