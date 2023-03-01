package main

import (
	"fmt"
	"github.com/tebie6/identifier_generate_service/conf"
	"github.com/tebie6/identifier_generate_service/models"
	"github.com/tebie6/identifier_generate_service/rpc"
	"github.com/tebie6/identifier_generate_service/tools/log"
	"github.com/tebie6/identifier_generate_service/tools/logger"
	"github.com/tebie6/identifier_generate_service/worker"
	"os"
)

func main() {

	// 初始化配置
	confpath := "./conf/conf.ini"
	err := conf.ParseConfigINI(confpath)
	if err != nil {
		fmt.Println("err : parse config failed", err.Error())
		os.Exit(1)
	}

	// 初始化log
	logPath := conf.GetConfigString("app", "log_path")

	logger.InitLogger(logPath, "20060102") // 日志
	el, err := conf.GetConfigInt("log", "level")
	if err != nil {
		el = 0
	}
	log.SetLogErrorLevel(int(el))

	// 初始化数据库
	models.InitModel()

	// 初始化ID生成节点
	worker.InitWorker()

	// 初始化RPC
	rpc.InitRpc()

}
