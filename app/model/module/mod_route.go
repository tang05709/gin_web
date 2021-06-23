package module

import (
	"festival/app/model"
)

// 线路
// power by 7be.cn
type ModRoute struct {
	model.BaseModel
	Type uint `form:"type" json:"type" gorm:"column:type;type:tinyint(1);comment:类型:1建党,2建司;"`
	Cover string `form:"cover" json:"cover" gorm:"column:cover;type:varchar(280);comment:封面图;"`
	Name string `form:"name" json:"name" gorm:"column:name;type:varchar(280);comment:姓名;"`
	Step uint `form:"step" json:"step" gorm:"column:step;type:int(10);comment:需要步数;"`
	
}

func (ModRoute) TableName() string {
	return "mod_route"
}
