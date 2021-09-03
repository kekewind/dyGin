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
	"strings"
	"time"
)

/**
  上传 抖音用户的ids
*/
func UploadUserIds(context *gin.Context) {
	message := context.PostForm("ids")
	if message == "" {
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
			mysql.DB.Model(&model.UserId{}).Save(&addDate)
		}
	}
	util.JsonWrite(context, 200, nil, "添加成功")
	return
}

//脚本获取用户抖音Id
func GetDyUserIdOne(context *gin.Context) {
	nickname := context.Query("nickname")

	phone := model.UserId{}
	mysql.DB.Where("status=1").Find(&phone)
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
	util.JsonWrite(context, 200, "", phone.UserId)
	println()

}
