package demo

import (
	"festival/app/common/response"

	"github.com/gin-gonic/gin"
)

func Chart(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/chart").ToTpl()
}

func Buttons(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/buttons").ToTpl()
}

func Editors(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/editors").ToTpl()
}
func Genernal(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/general").ToTpl()
}

func Form(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/generalForm").ToTpl()
}

func Modals(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/modals").ToTpl()
}

func Modify(c *gin.Context) {
	response.BuildTpl(c, "backend/system/demo/modify").ToTpl()
}
