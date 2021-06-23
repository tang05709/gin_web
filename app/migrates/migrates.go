package migrates

import (
	"festival/app/common/db"
	"festival/app/model/system"
)

func init() {
	db.DbList = append(db.DbList,
		system.SysMenu{},
		system.SysRole{},
		system.SysUser{},
	)
}
