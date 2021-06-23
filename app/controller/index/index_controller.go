package index

import (
	"festival/app/common/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.BuildTpl(c, "wecolme").ToTpl()
}
