package system

import (
	"festival/app/common/response"
	"festival/app/common/utils/page"
	"festival/app/model"
	systemModel "festival/app/model/system"
	"festival/app/request"
	systemService "festival/app/service/system"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func LoginPage(c *gin.Context) {
	if systemService.IsLogin(c) {
		c.Redirect(http.StatusFound, "/admin/index")
		return
	}
	response.BuildTpl(c, "backend/login.html").ToTpl()
}

func UserProfile(c *gin.Context) {
	response.BuildTpl(c, "backend/system/user/profile").ToTpl()
}

func UserEdit(c *gin.Context) {
	var req *request.IdReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p := gin.H{}
	if req != nil {
		user, err := systemService.GetUser(req.Id)
		if err != nil {
			c.Redirect(http.StatusFound, "/500")
			return
		}
		p["userInfo"] = user
	}
	err, roles := systemService.RoleAll()
	if err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p["allRole"] = roles
	response.BuildTpl(c, "backend/system/user/edit").ToTpl(p)
}

func UserSave(c *gin.Context) {
	var req *systemModel.SysUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	var err error
	if req != nil && req.Id != 0 {
		_, err = systemService.Update(req)
	} else {
		err = systemService.Add(req)
	}
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func UserSaveProfile(c *gin.Context) {
	var req *systemModel.SysUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	var err error
	cuser := systemService.GetProfile(c)
	if req == nil || req.Id == 0 || cuser.Id != req.Id {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	u, err := systemService.Update(req)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	// 更新session
	systemService.SaveUserToSession(u, c)
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func UserDelete(c *gin.Context) {
	var req *request.IdReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
		response.ErrorResp(c).SetMsg("参数有误").ToJson()
		return
	}
	item := systemModel.SysUser{}
	item.Id = uint(req.Id)
	err = systemService.DeleteUser(item)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func UserList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	paging := page.ToPaging(req)
	err, result, total := systemService.UserList(paging)
	if err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	data := model.TableDataInfo{
		Code: 0,
		Msg:  "操作成功",
		Rows: result,
		Page: *paging.CalcPage(total),
	}
	response.BuildTpl(c, "backend/system/user/list.html").TableTpl(data)
}

//验证登陆
func CheckLogin(c *gin.Context) {
	var req *request.UserReq
	//获取参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}

	if req == nil {
		response.ErrorResp(c).SetMsg("用户名或密码错误").ToJson()
		return
	}

	//比对验证码
	verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)

	if !verifyResult {
		response.ErrorResp(c).SetMsg("验证码不正确").ToJson()
		return
	}

	//验证账号密码
	if _, err := systemService.Login(req.UserName, req.Password, c); err != nil {
		response.ErrorResp(c).SetMsg("账号或密码不正确").ToJson()
	} else {
		response.SucessResp(c).SetMsg("登陆成功").ToJson()
	}
}
