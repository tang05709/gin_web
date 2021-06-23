package system

import (
	"festival/app/model"
)

type SysUser struct {
	model.BaseModel
	Account  string  `form:"account" json:"account" binding:"required,gte=6" gorm:"not null;unique;primary_key;column:account;type:varchar(20);comment:账号;"`
	UserName string  `form:"username" json:"username" binding:"required" gorm:"column:user_name;type:varchar(20);comment:昵称;"`
	Email    string  `form:"email" json:"email" gorm:"column:email;default:null;type:varchar(30);comment:邮箱;"`
	Phone    string  `form:"phone" json:"phone" gorm:"column:phone;default:null;type:varchar(20);comment:电话;"`
	Gender   uint    `form:"gender" json:"gender" gorm:"column:gender;default:0;type:tinyint(1);comment:性别(0未知 1男 2女);"`
	Avatar   string  `form:"avatar" json:"avatar" gorm:"column:avatar;default:http://f.rn4.cn/logo.png;type:varchar(200);comment:头像;"`
	Password string  `form:"password" json:"password" binding:"required,gte=6" gorm:"column:password;not null;type:varchar(200);comment:密码;"`
	Salt     string  `json:"salt" gorm:"column:salt;default:'';type:varchar(20);comment:盐值;"`
	Status   uint    `json:"status" gorm:"column:status;default:1;type:tinyint(1);comment:状态(0异常 1正常);"`
	RoleId   uint    `form:"roleId" json:"roleId"  gorm:"column:role_id;comment:角色id"`
	Role     SysRole `json:"role"`
}



func (SysUser) TableName() string {
	return "s_users"
}
