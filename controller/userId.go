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
	"github.com/wangyi/dyGin/dao/redis"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/util"
	"strconv"
	"strings"
	"time"
)

/**
  上传 抖音用户的ids
*/
func UploadUserIds(context *gin.Context) {
	message := context.PostForm("ids")
	path := context.PostForm("path")
	if message == "" || path == "" {
		util.JsonWrite(context, -101, nil, "上传的用户id不可以为空")
		return
	}
	ids := strings.Split(message, "\n")
	for _, ids := range ids {
		//判断是否存在重复
		phone := model.UserId{}
		mysql.DB.Where("user_id=?", ids).Find(&phone)
		if phone.ID == 0 {
			addDate := model.UserId{
				UserId:  ids,
				Status:  1,
				UseUser: "",
				Updated: time.Now().Unix(),
				Created: time.Now().Unix(),
			}
			addDate.Path, _ = strconv.Atoi(path)
			mysql.DB.Model(&model.UserId{}).Save(&addDate)
		}
	}
	util.JsonWrite(context, 200, nil, "添加成功")
	return
}

//脚本获取用户抖音Id
func GetDyUserIdOne(context *gin.Context) {
	soul, _ := redis.Rdb.Get("soul").Result()
	if soul == "1" {
		redis.Rdb.Set("soul", "2", 0)
		nickname := context.Query("nickname")
		path := context.Query("path")
		phone := model.UserId{}
		mysql.DB.Where("status=1 ").Where("path=?", path).Find(&phone)
		if phone.ID == 0 {
			util.JsonWrite(context, -101, "", "没有库存了")
			return
		}
		//更新数据
		updateData := model.UserId{
			UseUser: nickname,
			Status:  2,
			Updated: time.Now().Unix(),
		}
		mysql.DB.Model(&model.UserId{}).Where("id= ?", phone.ID).Update(&updateData)
		redis.Rdb.Set("soul", "1", 0)
		util.JsonWrite(context, 200, "", phone.UserId)
		return
	}else {
		util.JsonWrite(context, -101, "", "访问频繁,请稍等")
		return
	}
}
