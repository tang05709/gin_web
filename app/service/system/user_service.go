package system

import (
	"festival/app/common/cache"
	"festival/app/common/db"
	"festival/app/common/middleware/sessions"
	"festival/app/common/utils/gconv"
	"festival/app/common/utils/gmd5"
	"festival/app/common/utils/page"
	"festival/app/common/utils/random"
	"festival/app/model"
	"festival/app/model/system"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

//用户session列表
var SessionList sync.Map

//保存用户信息到session
func SaveUserToSession(user *system.SysUser, c *gin.Context) string {
	session := sessions.Default(c)
	session.Set(model.USER_ID, user.Id)
	tmp, _ := json.Marshal(user)
	session.Set(model.USER_SESSION_MARK, string(tmp))
	session.Save()
	sessionId := session.SessionId()
	SessionList.Store(sessionId, c)
	return sessionId
}

// 判断用户是否已经登录
func IsLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	uid := session.Get(model.USER_ID)
	if uid != nil {
		return true
	}
	return false
}

// 增加用户
func Add(user *system.SysUser) (err error) {
	newSalt := random.GenerateSubId(6)
	newToken := user.Account + user.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	user.Salt = newSalt
	user.Password = newToken
	err = db.Global().Create(&user).Error
	return err
}

// 删除用户
func DeleteUser(user system.SysUser) (err error) {
	err = db.Global().Where("`id` = ?", user.Id).Delete(&system.SysUser{}).Error
	return err
}

// 更新用户
func Update(user *system.SysUser) (u *system.SysUser, err error) {
	if user.Password != "" {
		newSalt := random.GenerateSubId(6)
		newToken := user.Account + user.Password + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		user.Salt = newSalt
		user.Password = newToken
	}
	err = db.Global().Omit("account").Updates(&user).Error
	return user, err
}

// 获取用户列表
func UserList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&system.SysUser{})
	users := make([]system.SysUser, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Preload("Role").Order("created_at desc").Find(&users).Error
	return err, users, total
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func Login(loginnName, password string, c *gin.Context) (sessionId string, err error) {
	var user system.SysUser
	//查询用户信息
	user = system.SysUser{Account: loginnName}
	err = db.Global().Where("account = ?", loginnName).First(&user).Error
	if err != nil {
		return "", err
	}
	//校验密码
	token := user.Account + password + user.Salt
	token = gmd5.MustEncryptString(token)
	fmt.Println(token)
	if user.Password != token {
		return "", errors.New("账号或密码错误")
	}
	sessionId = SaveUserToSession(&user, c)
	return sessionId, err
}

// 用户注销
func LoginOut(c *gin.Context) error {
	user := GetProfile(c)
	if user != nil {
		ClearMenuCache(user)
	}
	session := sessions.Default(c)
	sessionId := session.SessionId()
	SessionList.Delete(sessionId)
	session.Delete(model.USER_ID)
	session.Delete(model.USER_SESSION_MARK)
	return session.Save()
}

// 获得用户信息详情
func GetProfile(c *gin.Context) *system.SysUser {
	session := sessions.Default(c)
	tmp := session.Get(model.USER_SESSION_MARK)
	if tmp == nil {
		return nil
	}
	s := tmp.(string)
	var u system.SysUser
	err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		return nil
	}
	return &u
}

func GetUser(id uint) (u system.SysUser, err error) {
	var user system.SysUser
	user.Id = uint(id)
	err = db.Global().First(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

//清空用户菜单缓存
func ClearMenuCache(user *system.SysUser) {
	cache.Instance().Delete(model.MENU_CACHE + gconv.String(user.Id))
}

//强退用户
func ForceLogout(sessionId string) error {
	var tmp interface{}
	if v, ok := SessionList.Load(sessionId); ok {
		tmp = v
	}

	if tmp != nil {
		c := tmp.(*gin.Context)
		if c != nil {
			LoginOut(c)
		}
	}

	return nil
}
