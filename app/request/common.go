package request

// 通用分页请求
type PageReq struct {
	Page     int `form:"page"`     //当前页码
	PageSize int `form:"pageSize"` //每页数
}

//通用修改请求
type IdReq struct {
	Id uint `json:"id" uri:"id"` //主键ID
}

//通用用户分页请求
type UserIdReq struct {
	UserId uint `json:"userId" form:"userId"` //用户ID
}


type UploadFile struct {
	Name string `json:"name" form:"name"`
	Url  string `json:"url" form:"url"`
}

type UserGoodsStatus struct {
	Id uint `json:"id" uri:"id"` //主键ID
	Type  string `json:"type" form:"type"`
}
