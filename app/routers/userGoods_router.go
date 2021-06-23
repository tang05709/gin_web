package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/module"
)

// 兑换礼品表
// power by 7be.cn
func init() {
	g2 := router.New("admin", "/admin/module", auth.Auth)
	g2.GET("/userGoodss", true, module.ModUserGoodsList)
	g2.GET("/userGoods/edit", true, module.ModUserGoodsEdit)
	g2.POST("/userGoods/save", true, module.ModUserGoodsSave)
	g2.POST("/userGoods/del", true, module.ModUserGoodsDel)
	g2.POST("/userGoods/status", true, module.ModUserGoodsStatus)
}
