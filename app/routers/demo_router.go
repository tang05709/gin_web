package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/demo"
)

func init() {
	g2 := router.New("admin", "/admin/demo", auth.Auth)
	g2.GET("/chart", false, demo.Chart)
	g2.GET("/buttons", false, demo.Buttons)
	g2.GET("/editors", false, demo.Editors)
	g2.GET("/genernal", false, demo.Genernal)
	g2.GET("/form", false, demo.Form)
	g2.GET("/modals", false, demo.Modals)
}
