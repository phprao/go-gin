package server

import (
	"log"
	"os"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"votePlatfom/configs"
	"votePlatfom/pkg/common"
	_ "votePlatfom/pkg/inition"
	"votePlatfom/pkg/utils/mysqlUtil"
	"votePlatfom/routes"
)

var GRouter *gin.Engine

func initParam() string {
	// 设置运行模式
	gin.SetMode(common.GetAppEnv())

	// 初始化应用
	if os.Getenv("APP_DEBUG") == "true" {
		GRouter = gin.Default()
	} else {
		GRouter = gin.New()
		GRouter.Use(gin.Recovery())
	}

	// 注册路由
	routes.RegisterRoutes(GRouter)

	// 启动服务
	server := configs.GetServerConfig("http")
	serverAddr := server.Host + ":" + server.Port

	return serverAddr
}

// simple run
func Run() {
	// 服务关闭时关闭所有连接
	defer mysqlUtil.CloseAll()

	serverAddr := initParam()
	err := GRouter.Run(serverAddr)
	if nil != err {
		panic("server run error: " + err.Error())
	}
}

// with endless graceful restart
func RunWithGraceful() {
	// 服务关闭时关闭所有连接
	defer mysqlUtil.CloseAll()

	serverAddr := initParam()
	tmpServer := endless.NewServer(serverAddr, GRouter)
	tmpServer.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := tmpServer.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
