package tool

import (
	"festival/app/common/cfg"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(cfg.Instance().Logger.Path, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName("latest_log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if gin.Mode() == gin.DebugMode {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func H2Map(params []gin.H) map[string]interface{} {
	tmp := make(map[string]interface{})
	if params == nil || len(params) == 0 {
		return tmp
	}
	for k, v := range params[0] {
		tmp[k] = v
	}
	return tmp
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
