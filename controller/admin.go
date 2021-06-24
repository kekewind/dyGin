package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/util"
	"time"
)

// 注册参数 检验的定义
type checkRegister struct {
	AdminName     string `form:"admin_name" json:"admin_name" xml:"admin_name" binding:"required"`
	AdminPassword string `form:"admin_password" json:"admin_password" xml:"admin_password" binding:"required"`
	Token         string
	Price         float32
	PriceTwo      float32
	QqNum         string
	QqLink        string
	CreatedAt     int
	UpdatedAt     int
}

/**
管理员
*/

func CreatAdmin(content *gin.Context) {

	username := content.Query("username")
	password := content.Query("password")

	one := model.Admin{}
	one.AdminUser = username
	one.AdminPassword = password
	//err := mysql.DB.Save(&one).Error
	updateData := model.Admin{
		AdminPassword: "11323243",
		UpdatedAt:     time.Now().Unix(),
		Update:        time.Now().Unix(),
	}
	err := mysql.DB.Model(&model.Admin{}).Update(&updateData).Error
	if err != nil {
		util.JsonWrite(content, -101, nil, "添加失败")
		return
	}

	util.JsonWrite(content, 200, nil, "添加成功")
	return

}
