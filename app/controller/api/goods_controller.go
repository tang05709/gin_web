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

// 获取所有礼品
func ApiGoodsList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ErrorResp(c).SetMsg("参数错误").ToJson()
		return
	}
	paging := page.ToPaging(req)
	err, result, total := service.ModGoodsList(paging)
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

//获取用户礼品
func UserGoodsList(c *gin.Context) {
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
	err, result, total := service.UserGoodsListByUser(userReq.UserId, paging)
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

// 获取礼品详情
func GoodsDetail(c *gin.Context) {
	var req *request.IdReq
	//获取参数
	if err := c.ShouldBindUri(&req); err != nil {
		response.ErrorResp(c).SetMsg("参数错误").ToJson()
		return
	}
	item, err := service.GetModUserGoods(req.Id)
	if err != nil {
		response.ErrorResp(c).SetMsg("无数据").ToJson()
		return
	}
	response.SucessResp(c).SetData(item).ToJson()
}

func UserGoodsSave(c *gin.Context) {
	var req *module.ModUserGoods
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg("参数错误" + err.Error()).ToJson()
		return
	}
	err := service.AddModUserGoods(req)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}
