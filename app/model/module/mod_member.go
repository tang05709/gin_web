package module

import (
	"festival/app/model"
)

// 微信用户
// power by 7be.cn
type ModMember struct {
	model.BaseModel
	Openid  string `form:"openid" json:"openid" gorm:"column:openid;type:varchar(280);comment:微信标识;"`
	Route   uint   `form:"route" json:"route" gorm:"column:route;type:tinyint(1);comment:路线:0未选择,1建党,2建司;"`
	Name    string `form:"name" json:"name" gorm:"column:name;type:varchar(280);comment:姓名;"`
	Phone   string `form:"phone" json:"phone" gorm:"column:phone;type:varchar(11);comment:手机号;"`
	Agency  string `form:"agency" json:"agency" gorm:"column:agency;type:varchar(280);comment:投保机构;"`
	Avg     uint   `form:"avg" json:"avg" gorm:"column:avg;type:int(10);comment:平均步数;"`
	Total   uint   `form:"total" json:"total" gorm:"column:total;type:int(10);comment:累积步数;"`
	Address string `form:"address" json:"address" gorm:"column:address;type:varchar(280);comment:地址;"`
}

func (ModMember) TableName() string {
	return "mod_member"
}
