package system

import (
	"festival/app/common/response"
	"festival/app/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("upload")
	if err != nil {
		response.ErrorResp(c).SetMsg("参数有误").ToJson()
		return
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	allowExts := []string {".jpg", ".png", ".gif", ".jpeg", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx", ".pdf"}
	for _, ext := range allowExts {
		if ext == fileExt {

		}
	}

	now := time.Now()
	//文件夹路径
	fileDir := fmt.Sprintf("resource/static/upload/%s", now.Format("200601"))
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		response.ErrorResp(c).SetMsg("无权限").ToJson()
		return
	}
	//文件路径
	timeStamp := now.Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	fmt.Println(filePathStr)
	//将浏览器客户端上传的文件拷贝到本地路径的文件里面，此处也可以使用io操作
	c.SaveUploadedFile(fileHeader,filePathStr)
	response.SucessResp(c).SetData(request.UploadFile{Name: fileName, Url: filepath.ToSlash(filePathStr[8 : ])}).ToJson()
}