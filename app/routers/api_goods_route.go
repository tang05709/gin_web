package routers

import (
	"festival/app/common/router"
	"festival/app/controller/api"
)

// 礼品
func init() {
	g1 := router.New("api", "/api")
	g1.POST("/goods", false, api.ApiGoodsList)
	g1.POST("/user-goods", false, api.UserGoodsList)
	g1.POST("/user-goods/save", false, api.UserGoodsSave)
	g1.GET("/user-goods/:id", false, api.GoodsDetail)
}
