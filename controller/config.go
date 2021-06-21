/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/util"
	"net/http"
)

//获取版本信息
func GetVersion(context *gin.Context) {

	if mysql.DB.HasTable(&model.Config{}) {
		fmt.Println("数据库已经存在了!")
		mysql.DB.AutoMigrate(&model.Config{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		mysql.DB.CreateTable(&model.Config{})
		insert := model.Config{}
		insert.Version = 1
		mysql.DB.Save(&insert)

	}
	config := model.Config{}
	result := mysql.DB.Where("id=1").First(&config).Error
	if result != nil {
		util.JsonWrite(context, -101, nil, "查询失败"+result.Error())
		return
	}

	fmt.Println(config.Version)
	context.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
		"msg":    config.Version,
	})
	return

}
