package auth

import (
	"festival/app/common/middleware"
	"festival/app/common/router"
	"festival/app/model"
	systemService "festival/app/service/system"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(c *gin.Context) {

	//判断是否登陆
	if systemService.IsLogin(c) {
		//根据url判断是否有权限

		url, ok, _ := middleware.GetUserPerm(c)
		need := router.FindPermission(url)
		if need && !ok {
			back(c, "/403", 403, "您没有操作权限")
			return
		}
		c.Next()
	} else {
		back(c, "/admin/login", 401, "未登陆")
	}
}

func back(c *gin.Context, path string, code int, msg string) {
	ajaxString := c.Request.Header.Get("X-Requested-With")
	if strings.EqualFold(ajaxString, "XMLHttpRequest") {
		c.JSON(http.StatusOK, model.CommonRes{
			Code: code,
			Msg:  msg,
		})
		c.Abort()
	} else {
		c.Redirect(http.StatusFound, path)
		c.Abort()
	}
}
