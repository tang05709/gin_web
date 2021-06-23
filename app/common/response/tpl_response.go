package response

import (
	"festival/app/common/middleware/sessions"
	"festival/app/common/utils/tool"
	"festival/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

//返回一个tpl响应
func BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

//返回一个错误的tpl响应
func ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/error.html",
	}
	return &t
}

//返回一个无操作权限tpl响应
func ForbiddenTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/unauth.html",
	}
	return &t
}

//输出页面模板
func (resp *TplResp) WriteTpl(params ...gin.H) {
	session := sessions.Default(resp.c)
	uid := session.Get(model.USER_ID)
	if params == nil || len(params) == 0 {
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{"uid": uid})
	} else {
		params[0]["uid"] = uid
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
	//resp.c.Abort()
}

//返回一个错误的tpl响应
func (resp *TplResp) TableTpl(t model.TableDataInfo) {
	resp.c.Set("template", resp.tpl)
	resp.c.Set("data", tool.Struct2Map(t))
}

// pongo2 模板引擎
func (resp *TplResp) ToTpl(params ...gin.H) {
	resp.c.Set("template", resp.tpl)
	resp.c.Set("data", tool.H2Map(params))
	// session := sessions.Default(resp.c)
	// uid := session.Get(model.USER_ID)
	// if params == nil || len(params) == 0 {
	// 	resp.c.Set("template", resp.tpl)
	// 	resp.c.Set("data", tool.H2Map(gin.H{"uid": uid}))
	// } else {
	// 	params[0]["uid"] = uid
	// 	resp.c.Set("template", resp.tpl)
	// 	resp.c.Set("data", tool.H2Map(params[0]))
	// }
	//resp.c.Abort()
}
