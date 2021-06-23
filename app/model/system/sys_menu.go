package system

import (
	"festival/app/model"
)

type SysMenu struct {
	model.BaseModel
	Name       string    `json:"name" gorm:"column:name;not null;type:varchar(20);comment:名称;"`
	ParentId   uint      `json:"parentId" gorm:"column:parent_id;type:bigint(20);comment:父ID;"`
	Sort       uint      `json:"sort" gorm:"column:sort;default:0;type:int(10);comment:排序;"`
	Router     string    `json:"router" gorm:"column:router;not null;type:varchar(100);comment:路由地址/链接/接口;"`
	Permission string    `json:"permission" gorm:"column:permission;default:'';not null;type:varchar(100);comment:权限标识;"`
	Type       uint      `json:"type" gorm:"column:type;default:0;type:tinyint(1);comment:类型(0-页面 1-接口);"`
	Method     string    `json:"method" gorm:"column:method;default:'GET';type:varchar(10);comment:类型:(GETPOST);"`
	Visible    uint      `json:"visible" gorm:"column:visible;default:1;type:tinyint(1);comment:是否显示(0否 1是);"`
	Icon       string    `json:"icon" gorm:"column:icon;default:'fa-cog';type:varchar(100);comment:图标;"`
	Extra      string    `json:"extra" gorm:"column:extra;default:null;type:varchar(100);comment:额外参数;"`
	Children   []SysMenu `json:"children" gorm:"-"`
}

func (SysMenu) TableName() string {
	return "s_menus"
}
