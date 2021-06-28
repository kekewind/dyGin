package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/wangyi/dyGin/controller"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/logger"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/util"
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
	//创建管理员
	r.GET("/admin/CreatAdmin", controller.CreatAdmin)
	//脚本获取 一个抖音号  GetDyOneInformation
	r.GET("/collect/GetDyOneInformation", controller.GetDyOneInformation)
	//GetPhoneList 获取手机列表
	r.GET("/phone/GetPhoneList", controller.GetPhoneList)
	//脚本 获取自己的 配置 GetSelfConfig
	r.GET("/phone/GetSelfConfig", controller.GetSelfConfig)
	//管理员修改 手机列表的 运行种类  SetAllOrOneTaskTypeValue
	r.GET("/phone/SetAllOrOneTaskTypeValue", controller.SetAllOrOneTaskTypeValue)
	//脚本上传
	r.POST("/uploadFile", func(context *gin.Context) {
		//单个文件
		file, err := context.FormFile("file")
		fmt.Println(file.Filename)

		if err != nil {
			util.JsonWrite(context, -101, nil, "上传文件错误!")
			return
		}
		path := file.Filename
		ok := context.SaveUploadedFile(file, `./static/script/`+path)
		if ok != nil {
			fmt.Println("保存出错完了" + ok.Error())
			util.JsonWrite(context, -1, nil, ok.Error())
			return
		}

		//更新数据 版本
		if err := mysql.DB.Model(model.Config{}).Where("id=1").Update("version", gorm.Expr("version+?", 1)).Error; err != nil {
			util.JsonWrite(context, 200, nil, "版本更新失!")

			return
		}

		util.JsonWrite(context, 200, nil, "上传成功!")

		return

	})
	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	return r
}
