package routers

import (
	"festival/app/common/router"
	"festival/app/controller/index"
	system "festival/app/controller/system"
)

func init() {
	// 加载登陆路由
	g1 := router.New("admin", "/")
	g1.ANY("/", false, index.Index)
	g1.GET("/captcha", false, index.CaptchaImage)
	g1.GET("/500", false, system.Error)
	g1.GET("/404", false, system.NotFound)
	g1.GET("/403", false, system.Unauth)

	g2 := router.New("admin", "/admin")
	g2.GET("/login", false, index.LoginPage)
	g2.POST("/login", false, index.CheckLogin)
	g2.GET("/logout", false, index.Logout)
}
