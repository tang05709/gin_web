package module

import (
	"festival/app/model"
	"festival/app/model/system"
)

// 点亮线路表
// power by 7be.cn
type ModUserRoute struct {
	model.BaseModel
	UserId uint `form:"userId" json:"userId" gorm:"column:userId;type:int(11);comment:用户ID;"`
	RouteId uint `form:"routeId" json:"routeId" gorm:"column:routeId;type:int(11);comment:线路ID;"`

	User system.SysUser `json:"user" gorm:"foreignKey:UserId;references:Id;comment:用户;"`
	Route ModRoute `json:"route" gorm:"foreignKey:RouteId;references:Id;comment:线路;"`
}

func (ModUserRoute) TableName() string {
	return "mod_user_route"
}
