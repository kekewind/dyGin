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
	"strconv"
	"time"
)

//上传资料信息

func UploadInformation(context *gin.Context) {
	age := context.Query("age")
	address := context.Query("address")
	like := context.Query("like")
	following := context.Query("following")
	follows := context.Query("follows")
	sign := context.Query("sign")
	sex := context.Query("sex")
	dyNumber := context.Query("dyNumber")
	usename := context.Query("username")
	kind := context.Query("type")
	if mysql.DB.HasTable(&model.Collect{}) {
		fmt.Println("数据库已经存在了!")
		mysql.DB.AutoMigrate(&model.Collect{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		mysql.DB.CreateTable(&model.Collect{})
	}

	insertData := model.Collect{}
	insertData.UseUser = usename
	insertData.Age = age
	insertData.Status = 1
	insertData.Type, _ = strconv.Atoi(kind)
	insertData.DyNumber = dyNumber
	insertData.Like, _ = strconv.Atoi(like)
	insertData.Following, _ = strconv.Atoi(following)
	insertData.Address = address
	insertData.Follows, _ = strconv.Atoi(follows)
	insertData.Sign = sign
	insertData.Sex = sex
	insertData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	insertData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	result := mysql.DB.Save(&insertData).Error
	if result != nil {
		util.JsonWrite(context, -101, result.Error(), "插入失败")
		return
	}
	util.JsonWrite(context, 200, nil, "插入成功")

}