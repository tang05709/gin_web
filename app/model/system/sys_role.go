package system

import (
	"festival/app/model"
)

type SysRole struct {
	model.BaseModel
	Name  string    `form:"name" json:"name" gorm:"not null;unique;column:name;type:varchar(20);comment:名称;"`
	Desc  string    `form:"desc" json:"desc" gorm:"column:desc;;type:varchar(200);comment:描述;"`
	Menus []SysMenu `form:"menus" json:"menus" gorm:"many2many:s_role_menus;"`
}

func (SysRole) TableName() string {
	return "s_roles"
}
