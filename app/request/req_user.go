package request

type UserReq struct {
	UserName     string `form:"username"  binding:"required,min=5,max=30"`
	Password     string `form:"password" binding:"required,min=5,max=30"`
	ValidateCode string `form:"validateCode" binding:"required,min=4,max=10"`
	IdKey        string `form:"idkey" binding:"required,min=5,max=30"`
}

type UserRoleReq struct {
	Id    uint     `json:"id"`
	Name  string   `json:"name"  binding:"required"`
	Desc  string   `json:"desc"  binding:"required"`
	Menus []string `json:"menus"`
}
