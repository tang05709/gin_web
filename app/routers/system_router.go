package routers

import (
	"festival/app/common/middleware/auth"
	"festival/app/common/router"

	"festival/app/controller/system"
)

func init() {
	g2 := router.New("admin", "/admin/system", auth.Auth)
	g2.GET("/users", true, system.UserList)
	g2.GET("/user/edit", true, system.UserEdit)
	g2.POST("/user/save", true, system.UserSave)
	g2.POST("/user/del", true, system.UserDelete)
	g2.GET("/menus", true, system.MenuList)
	g2.GET("/menu/edit", true, system.MenuEdit)
	g2.POST("/menu/save", true, system.MenuSave)
	g2.POST("/menu/del", true, system.MenuDel)
	g2.GET("/roles", true, system.RoleList)
	g2.GET("/role/edit", true, system.RoleEdit)
	g2.POST("/role/save", true, system.RoleSave)
	g2.POST("/role/del", true, system.RoleDelete)
}
