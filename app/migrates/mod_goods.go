package migrates

import (
	"festival/app/common/db"
	"festival/app/model/module"
)

// 礼品表
// power by 7be.cn
func init() {
	db.DbList = append(db.DbList,
		module.ModGoods{},
	)
}
