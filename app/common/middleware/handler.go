package middleware

import (
	"net/http"
	"strings"

	systemModel "festival/app/model/system"
	systemService "festival/app/service/system"

	"festival/app/model"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	. "github.com/gin-gonic/gin"
)

func InjectParams(c *Context) pongo2.Context {
	// 全局参数注入
	ajaxString := c.Request.Header.Get("X-Requested-With")
	if !strings.EqualFold(ajaxString, "XMLHttpRequest") {
		menus, err1 := systemService.GetUserMenus(c)
		_, ok, router := GetUserPerm(c)
		response := make(map[string]interface{}, 0)
		data, _ := c.Get("data")
		var isMap bool
		response, isMap = data.(map[string]interface{})
		if !isMap {
			response = make(map[string]interface{}, 0)
		}
		if ok {
			response["router"] = router
		}
		response["user"] = systemService.GetProfile(c)
		if err1 == nil {
			response["menus"] = menus
		}
		// pjax := strings.ToLower(c.Request.Header.Get("X-PJAX"))
		// if pjax == "true" {
		// 	response["pjax"] = true
		// } else {
		// 	response["pjax"] = false
		// }
		return response
	}
	return nil
}

// 获取用户权限菜单
func GetUserPerm(c *gin.Context) (url string, ok bool, permission systemModel.SysMenu) {
	url = c.Request.URL.Path
	strEnd := url[len(url)-1 : len(url)]
	if strings.EqualFold(strEnd, "/") {
		url = strings.TrimRight(url, "/")
	}
	method := strings.ToUpper(c.Request.Method)
	tmp, err := systemService.GetUserAuth(c)
	if err != nil {
		return url, false, systemModel.SysMenu{}
	}
	tk := method + "-" + url
	permission, ok = tmp[tk]
	return url, ok, permission
}

func HandleNotFound(c *gin.Context) {
	ajaxString := c.Request.Header.Get("X-Requested-With")
	if strings.EqualFold(ajaxString, "XMLHttpRequest") {
		c.JSON(http.StatusOK, model.CommonRes{
			Code: 404,
			Msg:  "方法不存在",
		})
		c.Abort()
	} else {
		c.Redirect(http.StatusFound, "/404")
		c.Abort()
	}
	return
	// return func(c *gin.Context) {
	// 	ajaxString := c.Request.Header.Get("X-Requested-With")
	// 	if strings.EqualFold(ajaxString, "XMLHttpRequest") {
	// 		c.JSON(http.StatusOK, model.CommonRes{
	// 			Code: 404,
	// 			Msg:  "方法不存在",
	// 		})
	// 		c.Abort()
	// 	} else {
	// 		c.Redirect(http.StatusFound, "/404")
	// 		c.Abort()
	// 	}
	// }
}
