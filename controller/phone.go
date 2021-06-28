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
)

//获取手机列表
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
	if task_type, isExist := context.GetQuery("task_type"); isExist == true {
		task_type, _ := strconv.Atoi(task_type)
		Db = Db.Where("type=?", task_type)

	}

	Db.Table("phones").Count(&total)
	Db = Db.Model(&phones).Offset((page - 1) * limit).Limit(limit).Order("update desc")
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
