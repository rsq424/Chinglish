package router

import (
	"Chinglish/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("status", v1.Status)
	}
	return BaseRouter
}