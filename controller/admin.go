package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/gin_project/dao/mysql"
	"github.com/wangyi/gin_project/model"
	"github.com/wangyi/gin_project/util"
	"github.com/wangyi/gin_project/verify"
	"net/http"
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
注册管理员
*/

func Register(context *gin.Context) {

	//校验的初始化   对token  包括接口用户的接口进行 权限控制
	b := verify.Verify(context)
	if !b {
		return //校验不通过  请求到这里就停止了
	}
	//获取 参数
	var newAdmin checkRegister
	//	newAdmin := checkRegister{AdminName: context.Query("admin_name"), AdminPassword: context.Query("admin_password"), Price: 0.00, PriceTwo: 0.01, Token: util.SetToken(36)}
	if err := context.ShouldBindQuery(&newAdmin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": -101,
			"err":  err.Error(),
		})
		return
	}
	//newAdmin = checkRegister{AdminName: context.Query("admin_name"), AdminPassword: context.Query("admin_password"), Price: 0.00, PriceTwo: 0.01, Token: util.SetToken(36)}

	if mysql.DB.HasTable(&model.Admin{}) {
		fmt.Println("数据库已经存在了!")
		mysql.DB.AutoMigrate(&model.Admin{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		mysql.DB.CreateTable(&model.Admin{})
	}

	//判断 这个用户是否注册过
	admin := model.Admin{}
	mysql.DB.Where("admin_name=?", context.Query("admin_name")).Find(&admin)
	if admin.ID != 0 {
		util.JsonWrite(context, -101, nil, "请不要重复注册")
		return
	}
	//	newAdmin = model.Admin{AdminName: context.Query("admin_name"), AdminPassword: context.Query("admin_password"), Price: 0.00, PriceTwo: 0.01, Token: util.SetToken(36),QqNum: "89237890",QqLink: "989",CreatedAt: time.Now().Unix(),UpdatedAt: time.Now().Unix()}
	new := model.Admin{}
	new.AdminName=context.Query("admin_name")
	new.AdminPassword=context.Query("admin_password")
	new.Token=util.SetToken(36)
	new.UpdatedAt=time.Now().Unix()
	new.CreatedAt=time.Now().Unix()

	result := mysql.DB.Save(&new).Error
	if result != nil {
		util.JsonWrite(context, -101, nil, "注册失败")
		return
	}
	util.JsonWrite(context, 200, nil, "注册成功")

	return

}
