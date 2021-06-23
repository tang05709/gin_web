package router

import (
	"github.com/gin-gonic/gin"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	TRACE   = "TRACE"
)

var GroupList = make([]*routerGroup, 0)

var PermissionMap = make(map[string]bool, 0)

//路由信息
type router struct {
	Method       string            //方法名称
	RelativePath string            //url路径
	NeedPerm     bool              //是否需要权限
	HandlerFunc  []gin.HandlerFunc //执行函数
}

//路由组信息
type routerGroup struct {
	ServerName   string            //服务名称
	RelativePath string            //url路径
	Handlers     []gin.HandlerFunc //中间件
	Router       []*router         //路由信息
}

//根据url获取权限字符串
func FindPermission(url string) bool {
	return PermissionMap[url]
}

//创建一个路由组
func New(serverName, relativePath string, middleware ...gin.HandlerFunc) *routerGroup {
	var rg routerGroup
	rg.ServerName = serverName
	rg.Router = make([]*router, 0)
	rg.RelativePath = relativePath
	rg.Handlers = middleware
	GroupList = append(GroupList, &rg)
	return &rg
}

//添加路由信息
func (group *routerGroup) Handle(method, relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	var r router
	r.Method = method
	r.NeedPerm = permiss
	r.RelativePath = relativePath
	r.HandlerFunc = handlers
	group.Router = append(group.Router, &r)
	PermissionMap[group.RelativePath+relativePath] = permiss
	return group
}

//添加路由信息-ANY
func (group *routerGroup) ANY(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle("ANY", relativePath, permiss, handlers...)
	return group
}

//添加路由信息-GET
func (group *routerGroup) GET(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(GET, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-POST
func (group *routerGroup) POST(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(POST, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-OPTIONS
func (group *routerGroup) OPTIONS(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(OPTIONS, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-PUT
func (group *routerGroup) PUT(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PUT, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-PATCH
func (group *routerGroup) PATCH(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PATCH, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-HEAD
func (group *routerGroup) HEAD(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(HEAD, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-DELETE
func (group *routerGroup) DELETE(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(DELETE, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-CONNECT
func (group *routerGroup) CONNECT(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(CONNECT, relativePath, permiss, handlers...)
	return group
}

//添加路由信息-TRACE
func (group *routerGroup) TRACE(relativePath string, permiss bool, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(TRACE, relativePath, permiss, handlers...)
	return group
}
