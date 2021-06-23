package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"
	"festival/app/controller/index"
	"festival/app/controller/system"
)

func init() {
	g2 := router.New("admin", "/admin", auth.Auth)
	g2.GET("/index", false, index.AdminIndex)
	g2.GET("/profile", false, system.UserProfile)
	g2.POST("/saveProfile", false, system.UserSaveProfile)
}
