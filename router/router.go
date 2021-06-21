package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wangyi/gin_project/controller"
	"github.com/wangyi/gin_project/logger"
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
	r.GET("/centralism/register",controller.Register)








	/*	r.GET("/bookable", func(context *gin.Context) {

		var data Reg
		if err := context.ShouldBindQuery(&data); err != nil {
			fmt.Println("-----------------------------------------")
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := model.User{Password: data.Password, Username: data.Username, Token: util.SetToken(36), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
		fmt.Println(user)
		controller.Add(user)

	})*/

	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	return r
}

func getInformation(context *gin.Context) {

	 fmt.Println("====================")
	 fmt.Println(context.Params)



}

type Reg struct {
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
