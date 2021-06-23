package main

import (
	"festival/app/common/cfg"
	_ "festival/app/common/cron"
	"festival/app/common/db"
	"festival/app/common/log"
	"festival/app/common/middleware"
	"festival/app/common/middleware/sessions"
	"festival/app/common/middleware/sessions/memstore"
	"festival/app/common/server"
	_ "festival/app/migrates"
	_ "festival/app/routers"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// @title 7be_cn 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {
	gin.SetMode("debug")

	config := cfg.Instance()
	if config == nil {
		fmt.Printf("参数错误")
		return
	}
	// 初始化
	db.Global()
	db.AutoMigrate()
	log.Global()
	//后台服务状态
	adminStatus := config.Status.Admin
	//api服务状态
	apiStatus := config.Status.Api

	if adminStatus {
		store := memstore.NewStore([]byte("secret"))
		admin := server.New("admin", config.Admin.Address, log.GinRecovery(true), log.GinLogger(), middleware.Pongo2(), sessions.Sessions("mysession", store))
		admin.Static(config.Admin.ServerRoot)
		admin.Start(g)
	}

	if apiStatus {
		api := server.New("api", config.Api.Address, log.GinRecovery(true), log.GinLogger())
		api.Start(g)
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

EXIT:
	for {
		sig := <-sc
		log.Global().Error("获取到信号[%s]", zap.Any("err", sig.String()))
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	tdb, _ := db.Global().DB()
	defer tdb.Close()
	fmt.Println("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
