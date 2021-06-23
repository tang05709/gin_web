package migrates

import (
	"festival/app/common/db"
	"festival/app/model/module"
)

// 微信用户
// power by 7be.cn
func init() {
	db.DbList = append(db.DbList,
		module.ModMember{},
	)
}
