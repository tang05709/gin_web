package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/module"
)

// 礼品表
// power by 7be.cn
func init() {
	g2 := router.New("admin", "/admin/module", auth.Auth)
	g2.GET("/goodss", true, module.ModGoodsList)
	g2.GET("/goods/edit", true, module.ModGoodsEdit)
	g2.POST("/goods/save", true, module.ModGoodsSave)
	g2.POST("/goods/del", true, module.ModGoodsDel)
}
