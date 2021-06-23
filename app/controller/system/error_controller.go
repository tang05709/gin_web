package system

import (
	"festival/app/common/response"

	"github.com/gin-gonic/gin"
)

func Unauth(c *gin.Context) {
	response.BuildTpl(c, "error/error").ToTpl(gin.H{"code": "403", "error": "您没有访问权限"})
}

func Error(c *gin.Context) {
	response.BuildTpl(c, "error/error").ToTpl(gin.H{"code": "500", "error": "系统错误"})
}

func NotFound(c *gin.Context) {
	response.BuildTpl(c, "error/error").ToTpl(gin.H{"code": "404", "error": "页面丢失"})
}
