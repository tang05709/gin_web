package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/module"
)

// 线路
// power by 7be.cn
func init() {
	g2 := router.New("admin", "/admin/module", auth.Auth)
	g2.GET("/routes", true, module.ModRouteList)
	g2.GET("/route/edit", true, module.ModRouteEdit)
	g2.POST("/route/save", true, module.ModRouteSave)
	g2.POST("/route/del", true, module.ModRouteDel)
}
