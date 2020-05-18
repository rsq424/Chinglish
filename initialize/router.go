package initialize

import (
	"Chinglish/global"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	global.GVA_LOG.Info("router register success")
	return Router
}