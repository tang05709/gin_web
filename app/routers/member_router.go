package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/module"
)

// 微信用户
// power by 7be.cn
func init() {
	g2 := router.New("admin", "/admin/module", auth.Auth)
	g2.GET("/members", true, module.ModMemberList)
	g2.GET("/member/edit", true, module.ModMemberEdit)
	g2.POST("/member/save", true, module.ModMemberSave)
	g2.POST("/member/del", true, module.ModMemberDel)
}
