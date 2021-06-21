package main

import (
	"fmt"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/logger"
	"github.com/wangyi/dyGin/router"
	"github.com/wangyi/dyGin/setting"
	"go.uber.org/zap"
)

func main() {

	//加载配置
	if err := setting.Init(); err != nil {
		fmt.Println("配置文件初始化事变", err)
		return
	}

	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Println("日志初始化失败", err)
		return
	}

	defer zap.L().Sync() //缓存日志追加到日志文件中
	zap.L().Debug("LaLa")

	//链接数据库
	if err := mysql.Init(); err != nil {
	 	fmt.Println("mysql 链接失败,",err)
		return
	}

	defer mysql.Close()

	router.Setup()
}
