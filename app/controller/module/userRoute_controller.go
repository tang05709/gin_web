package module

import (
	"festival/app/common/response"
	"festival/app/model"
	"festival/app/model/module"
	"festival/app/request"
	service "festival/app/service/module"
	"fmt"
	"net/http"
	"festival/app/common/utils/page"
	"github.com/gin-gonic/gin"
)

// 点亮线路表
// power by 7be.cn

func ModUserRouteList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	paging := page.ToPaging(req)
	err, result, total := service.ModUserRouteList(paging)
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
	response.BuildTpl(c, "backend/module/userRoute/list").TableTpl(data)
}

func ModUserRouteEdit(c *gin.Context) {
	var req *request.IdReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p := gin.H{}
	if req != nil {
		item, err := service.GetModUserRoute(req.Id)
		if err != nil {
			c.Redirect(http.StatusFound, "/500")
			return
		}
		p["item"] = item
	}
	response.BuildTpl(c, "backend/module/userRoute/edit").ToTpl(p)
}

func ModUserRouteDel(c *gin.Context) {
	var req *request.IdReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
		response.ErrorResp(c).SetMsg("参数有误").ToJson()
		return
	}
	item := module.ModUserRoute{}
	item.Id = uint(req.Id)
	err = service.DeleteModUserRoute(item)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func ModUserRouteSave(c *gin.Context) {
	var req *module.ModUserRoute
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	var err error
	if req != nil && req.Id != 0 {
		fmt.Println(req)
		err = service.UpdateModUserRoute(req)
	} else {
		err = service.AddModUserRoute(req)
	}
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}
