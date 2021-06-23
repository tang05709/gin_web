package routers

import (
	"festival/app/common/router"
	"festival/app/controller/api"
)

// 点亮线路表
// power by 7be.cn
func init() {
	g1 := router.New("api", "/api")
	g1.POST("/routes", false, api.ApiRouteList)
	g1.POST("/user-routes", false, api.ApiUserRouteList)
	g1.POST("/user-routes/save", false, api.APiUserRouteSave)
	g1.POST("/user-routes/delete", false, api.APiUserRouteDelete)
}
