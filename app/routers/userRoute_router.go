package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/module"
)

// 点亮线路表
// power by 7be.cn
func init() {
	g2 := router.New("admin", "/admin/module", auth.Auth)
	g2.GET("/userRoutes", true, module.ModUserRouteList)
	g2.GET("/userRoute/edit", true, module.ModUserRouteEdit)
	g2.POST("/userRoute/save", true, module.ModUserRouteSave)
	g2.POST("/userRoute/del", true, module.ModUserRouteDel)
}
