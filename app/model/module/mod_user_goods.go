package module

import (
	"festival/app/model"
	"festival/app/model/system"
)

// 兑换礼品表
// power by 7be.cn
type ModUserGoods struct {
	model.BaseModel
	UserId uint `form:"userId" json:"userId" gorm:"column:userId;type:int(11);comment:用户ID;"`
	GoodsId uint `form:"goodsId" json:"goodsId" gorm:"column:goodsId;type:int(11);comment:礼品ID;"`
	Status int `form:"status" json:"status" gorm:"column:status;type:tinyint(1);comment:状态：0审核中1已发货-1驳回;"`
	Name string `form:"name" json:"name" gorm:"column:name;type:varchar(280);comment:用户名;"`
	Phone string `form:"phone" json:"phone" gorm:"column:phone;type:varchar(11);comment:电话号;"`
	Address string `form:"address" json:"address" gorm:"column:address;type:varchar(300);comment:联系地址;"`

	User system.SysUser `json:"user" gorm:"foreignKey:UserId;references:Id;comment:用户;"`
	Goods ModGoods `json:"goods" gorm:"foreignKey:GoodsId;references:Id;comment:礼品;"`
}

func (ModUserGoods) TableName() string {
	return "mod_user_goods"
}
