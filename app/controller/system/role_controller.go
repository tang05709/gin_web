package system

import (
	"festival/app/common/response"
	"festival/app/common/utils/page"
	"festival/app/model"
	"festival/app/model/system"
	systemModel "festival/app/model/system"
	"festival/app/request"
	systemService "festival/app/service/system"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	paging := page.ToPaging(req)
	err, result, total := systemService.RoleList(paging)
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
	response.BuildTpl(c, "backend/system/role/list").TableTpl(data)
}

func RoleEdit(c *gin.Context) {
	var req *request.IdReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p := gin.H{}
	if req != nil {
		item, err := systemService.GetRole(req.Id)
		if err != nil {
			c.Redirect(http.StatusFound, "/500")
			return
		}
		p["item"] = item
	}
	// 返回所有菜单树
	err, menus := systemService.MenuAll()
	err, menusTree := systemService.MenuTreeList()
	if err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p["allMenus"] = menus
	p["menusTree"] = menusTree
	response.BuildTpl(c, "backend/system/role/edit").ToTpl(p)
}

func RoleSave(c *gin.Context) {
	var req *request.UserRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	fmt.Println(req)
	// 找出所有菜单
	err, menus := systemService.MenuInIds(req.Menus)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	var role system.SysRole
	role.Id = req.Id
	role.Name = req.Name
	role.Desc = req.Desc
	role.Menus = menus
	if req != nil && req.Id != 0 {
		err = systemService.UpdateRole(role)
		// 清理当前角色的菜单缓存
		systemService.ClearRoleMenus(req.Id)
	} else {
		err = systemService.AddRole(role)
	}
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}

	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func RoleDelete(c *gin.Context) {
	var req *request.IdReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
		response.ErrorResp(c).SetMsg("参数有误").ToJson()
		return
	}
	item := systemModel.SysRole{}
	item.Id = uint(req.Id)
	err = systemService.DeleteRole(item)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}
