package response

import (
	"festival/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通用api响应
type TableResp struct {
	t *model.TableDataInfo
	c *gin.Context
}

//返回一个成功的消息体
func BuildTable(c *gin.Context, total int, rows interface{}) *TableResp {
	msg := model.TableDataInfo{
		Code: 0,
		Msg:  "操作成功",

		Rows: rows,
	}
	a := TableResp{
		t: &msg,
		c: c,
	}
	return &a
}

//输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOK, resp.t)
	resp.c.Abort()
}

//输出json到客户端
func (resp *TableResp) ToJson() {
	resp.c.JSON(http.StatusOK, resp.t)
	resp.c.Abort()
}
