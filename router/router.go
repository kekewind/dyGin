package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wangyi/dyGin/controller"
	"github.com/wangyi/dyGin/logger"
	_ "sync/atomic"
)

/**
  注册路由
*/
func Setup() *gin.Engine {
	r := gin.New()
	//添加记录日志的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//设置 静态文件目录

	r.Static("/static", "./static")

	// 路由分组
	r.GET("/link/AnalysisLink", controller.AnalysisLink)
	// 获取链接的 列表  link/GetLinksList
	r.GET("/link/GetLinksList", controller.GetLinksList)
	//GetOneLink
	r.GET("/link/GetOneLink", controller.GetOneLink)

	//脚本获取版本号
	r.GET("/config/GetVersion", controller.GetVersion)

	//上传采集数据 UploadInformation
	r.GET("/collect/UploadInformation", controller.UploadInformation)
	//获取采集的资料了 GetCollectInformation
	r.GET("/collect/GetCollectInformation", controller.GetCollectInformation)
	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	return r
}
