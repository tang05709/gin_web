package module

import (
	"festival/app/model"
)

// 礼品表
// power by 7be.cn
type ModGoods struct {
	model.BaseModel
	Title string `form:"title" json:"title" gorm:"column:title;type:varchar(200);comment:标题;"`
	Cover string `form:"cover" json:"cover" gorm:"column:cover;type:varchar(300);comment:封面;"`
	Num int `form:"num" json:"num" gorm:"column:num;type:int(11);comment:库存量;"`
	
}

func (ModGoods) TableName() string {
	return "mod_goods"
}
