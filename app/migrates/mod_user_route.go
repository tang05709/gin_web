package migrates

import (
	"festival/app/common/db"
	"festival/app/model/module"
)

// 点亮线路表
// power by 7be.cn
func init() {
	db.DbList = append(db.DbList,
		module.ModUserRoute{},
	)
}
