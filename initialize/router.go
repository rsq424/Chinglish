package initialize

import (
	"Chinglish/global"
	"Chinglish/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	//Router.Use(middleware.LoadTls())

	ApiGroup := Router.Group("")
	router.InitBaseRouter(ApiGroup)
	global.GVA_LOG.Info("router register success")
	return Router
}