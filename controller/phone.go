/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//获取手机列表 pc
func GetPhoneList(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	Db := mysql.DB
	var total int = 0
	phones := make([]model.Phone, 0)
	if status, isExist := context.GetQuery("status"); isExist == true {
		status, _ := strconv.Atoi(status)
		Db = Db.Where("status=?", status)
	}
	if taskType, isExist := context.GetQuery("task_type"); isExist == true {
		taskType, _ := strconv.Atoi(taskType)
		Db = Db.Where("type_task=?", taskType)

	}

	Db.Table("phones").Count(&total)
	Db = Db.Model(&phones).Offset((page - 1) * limit).Limit(limit).Order("updated desc")
	if err := Db.Find(&phones).Error; err != nil {
		util.JsonWrite(context, -101, nil, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":   1,
		"count":  total,
		"result": phones,
	})
	return
}

//针对脚本 获取自己的配置
func GetSelfConfig(context *gin.Context) {
	nickname := context.Query("nickname")
	phone := model.Phone{}
	mysql.DB.Where("nickname=?", nickname).Find(&phone)
	if phone.ID == 0 {
		//没有这部设备 添加设备
		addPhones := model.Phone{
			Nickname: nickname,
			Status:   1,
			TaskType: 0,
			Updated:  time.Now().Unix(),
			Created:  time.Now().Unix(),
		}
		result := mysql.DB.Save(&addPhones).Error
		if result != nil {
			util.JsonWrite(context, -101, nil, "新增设备失败")
			return
		}
		util.JsonWrite(context, 200, phone, "新增设备成功")
	}
	util.JsonWrite(context, 200, phone, "获取成功")
	return

}

//管理员修改脚本的运行
func SetAllOrOneTaskTypeValue(context *gin.Context) {
	kinds, _ := strconv.Atoi(context.Query("task_type"))
	ids := strings.Split(context.Query("ids"), "@")
	for _, id := range ids {
		updateDate := model.Phone{
			Updated:  time.Now().Unix(),
			TaskType: kinds,
		}
		mysql.DB.Where("id=?", id).Update(&updateDate)
	}

	util.JsonWrite(context, 200, nil, "修改成功")
	return

}
