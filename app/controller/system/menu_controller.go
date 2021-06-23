package system

import (
	"festival/app/common/response"
	"festival/app/model"
	systemModel "festival/app/model/system"
	"festival/app/request"
	systemService "festival/app/service/system"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MenuList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	err, result := systemService.MenuTreeList()
	if err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	data := model.TableDataInfo{
		Code: 0,
		Msg:  "操作成功",
		Rows: result,
	}
	response.BuildTpl(c, "backend/system/menu/list").TableTpl(data)
}

func MenuEdit(c *gin.Context) {
	var req *request.IdReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p := gin.H{}
	if req != nil {
		item, err := systemService.GetMenu(req.Id)
		if err != nil {
			c.Redirect(http.StatusFound, "/500")
			return
		}
		p["item"] = item
	}
	// 返回所有菜单
	err, menus := systemService.MenuAll()
	if err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p["allMenus"] = menus
	response.BuildTpl(c, "backend/system/menu/edit").ToTpl(p)
}

func MenuDel(c *gin.Context) {
	var req *request.IdReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
		response.ErrorResp(c).SetMsg("参数有误").ToJson()
		return
	}
	item := systemModel.SysMenu{}
	item.Id = uint(req.Id)
	err = systemService.DeleteMenu(item)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func MenuSave(c *gin.Context) {
	var req *systemModel.SysMenu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	var err error
	if req != nil && req.Id != 0 {
		fmt.Println(req)
		err = systemService.UpdateMenu(req)
		// 清理当前角色的菜单缓存
		systemService.ClearAllMenus()
	} else {
		err = systemService.AddMenu(req)
	}
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}
