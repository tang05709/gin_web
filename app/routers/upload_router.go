package routers

import (
	"festival/app/common/router"
	"festival/app/controller/system"
)

func init() {
	g1 := router.New("admin", "/admin")
	g1.ANY("/upload", false, system.UploadFile)
}