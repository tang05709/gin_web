package api

import (
	"festival/app/common/response"
	"festival/app/common/utils/page"
	"festival/app/model"
	"festival/app/model/module"
	"festival/app/request"
	service "festival/app/service/module"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 获取所有线路
func ApiRouteList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ErrorResp(c).SetMsg("参数错误").ToJson()
		return
	}
	paging := page.ToPaging(req)
	err, result, total := service.ModRouteList(paging)
	if err != nil {
		response.SucessResp(c).SetData(nil).ToJson()
		return
	}
	data := model.TableDataInfo{
		Code: 0,
		Msg:  "操作成功",
		Rows: result,
		Page: *paging.CalcPage(total),
	}
	response.SucessResp(c).SetData(data).ToJson()
}

// 根据用户查找点亮的线路
func ApiUserRouteList(c *gin.Context) {
	var req *request.PageReq
	var userReq *request.UserIdReq
	//获取参数
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		response.ErrorResp(c).SetMsg("参数错误").ToJson()
		return
	}
	if err := c.ShouldBindBodyWith(&userReq, binding.JSON); err != nil {
		response.ErrorResp(c).SetMsg("参数错误").ToJson()
		return
	}

	paging := page.ToPaging(req)
	err, result, total := service.UserRouteListByUser(userReq.UserId, paging)
	if err != nil || total <= 0 {
		response.ErrorResp(c).SetMsg("无数据").ToJson()
		return
	}
	data := model.TableDataInfo{
		Code: 0,
		Msg:  "操作成功",
		Rows: result,
		Page: *paging.CalcPage(total),
	}
	response.SucessResp(c).SetData(data).ToJson()
}

// 用户点亮线路
func APiUserRouteSave(c *gin.Context) {
	var req *module.ModUserRoute
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	err := service.AddModUserRoute(req)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

// 删除点亮的线路
func APiUserRouteDelete(c *gin.Context) {
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