package v1

import (
	"Chinglish/global"
	"Chinglish/global/response"
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	global.GVA_LOG.Info("status working")
	response.OkDetailed(gin.H{}, "运行正常", c)
}
