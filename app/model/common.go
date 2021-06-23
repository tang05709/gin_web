package model

import (
	"festival/app/common/utils/page"
	"time"

	"gorm.io/gorm"
)

//SESSION前缀
const (
	USER_ID           = "uid"
	USER_SESSION_MARK = "user_info"
)

// 登陆用户的菜单列表缓存前缀
const MENU_CACHE = "menu_cache"
const API_CACHE = "api_cache"
const AUTH_CACHE = "auth_cache"

//响应结果
const (
	SUCCESS      = 0   // 成功
	ERROR        = 500 //错误
	UNAUTHORIZED = 403 //无权限
	FAIL         = -1  //失败
)

// 通用api响应
type CommonRes struct {
	Code int         `json:"code"` //响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg  string      `json:"msg"`  //消息
	Data interface{} `json:"data"` //数据内容
}

// 验证码响应
type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

//通用分页表格响应
type TableDataInfo struct {
	Page page.Paging
	Rows interface{} `json:"rows"` //数据
	Code int         `json:"code"` //响应编码 0 成功 500 错误 403 无权限
	Msg  string      `json:"msg"`  //消息
}

// 通用model
type BaseModel struct {
	Id        uint `form:"id" json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
