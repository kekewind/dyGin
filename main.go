package main

import (
	"fmt"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/dao/redis"
	"github.com/wangyi/dyGin/logger"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/router"
	"github.com/wangyi/dyGin/setting"
	"go.uber.org/zap"
	"time"
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
		fmt.Println("mysql 链接失败,", err)
		return
	}

	defer mysql.Close()

	//redis 初始化
	//4.初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Println("redis文件初始化失败：", err)
		return
	}

	redis.Rdb.Set("soul", "1",0)
	//println(redis.Rdb.Get("soul").String())
	defer redis.Close()






	//进程 心跳 没30秒修改 数据设备的状态
	go func() {
		for true {
			err := mysql.DB.Model(&model.Phone{}).Update("status", 1).Error
			if err != nil {
				println("进程检测心跳失败")
			}
			time.Sleep(time.Second * 30)
		}
	}()

	router.Setup()

}
