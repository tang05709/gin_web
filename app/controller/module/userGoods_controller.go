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

// 兑换礼品表
// power by 7be.cn

func ModUserGoodsList(c *gin.Context) {
	var req *request.PageReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	paging := page.ToPaging(req)
	err, result, total := service.ModUserGoodsList(paging)
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
	response.BuildTpl(c, "backend/module/userGoods/list").TableTpl(data)
}

func ModUserGoodsEdit(c *gin.Context) {
	var req *request.IdReq
	//获取参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Redirect(http.StatusFound, "/500")
		return
	}
	p := gin.H{}
	if req != nil {
		item, err := service.GetModUserGoods(req.Id)
		if err != nil {
			c.Redirect(http.StatusFound, "/500")
			return
		}
		p["item"] = item
	}
	response.BuildTpl(c, "backend/module/userGoods/edit").ToTpl(p)
}

func ModUserGoodsDel(c *gin.Context) {
	var req *request.IdReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
		response.ErrorResp(c).SetMsg("参数有误").ToJson()
		return
	}
	item := module.ModUserGoods{}
	item.Id = uint(req.Id)
	err = service.DeleteModUserGoods(item)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func ModUserGoodsSave(c *gin.Context) {
	var req *module.ModUserGoods
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	var err error
	if req != nil && req.Id != 0 {
		fmt.Println(req)
		err = service.UpdateModUserGoods(req)
	} else {
		err = service.AddModUserGoods(req)
	}
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}

func ModUserGoodsStatus(c *gin.Context) {
	var req *request.UserGoodsStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}
	var status int = 0
	if req.Type == "reject" {
		status = -1
	}
	if req.Type == "send" {
		status = 1
	}
	err := service.UpdateModUserGoodsStatus(req.Id, status)
	if err != nil {
		response.ErrorResp(c).SetMsg("操作失败").ToJson()
		return
	}
	response.SucessResp(c).SetMsg("操作成功").ToJson()
}