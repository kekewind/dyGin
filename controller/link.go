package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/dyGin/dao/mysql"
	"github.com/wangyi/dyGin/model"
	"github.com/wangyi/dyGin/util"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
所有抖音的上传的 链接在这里处理
*/

//校验解析接口上传的参数

type CheckAnalysisLink struct {
	Type int    `form:"type" json:"type" xml:"type" binding:"required,min=1" `
	Link string `form:"link" json:"link" xml:"link" binding:"required,min=1"`
}

//解析上传的链接
func AnalysisLink(context *gin.Context) {
	kind := context.Query("type")
	link := context.Query("link")
	//	fmt.Println(kinds)
	var newAnalysisLink CheckAnalysisLink
	if err := context.ShouldBindQuery(&newAnalysisLink); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": -101,
			"err":  err.Error(),
		})
		return
	}

	arrayLinks := strings.Split(link, "\n")
	if mysql.DB.HasTable(&model.Link{}) {
		fmt.Println("数据库已经存在了!")
		mysql.DB.AutoMigrate(&model.Link{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		mysql.DB.CreateTable(&model.Link{})
	}
	for _, value := range arrayLinks {
		RoomId := GetDyId(value)
		//解析正则表达式，如果成功返回解释器
		reg1 := regexp.MustCompile(`\d{19}`)
		if reg1 == nil {
			continue
		}
		//根据规则提取关键信息
		result1 := reg1.FindAllStringSubmatch(RoomId, -1)

		if result1[0][0] != "" {
			//插入数据库
			//判断  这条链接是否存在
			LinkData := model.Link{}
			mysql.DB.Where("room_id=?", result1[0][0]).Find(&LinkData)
			if LinkData.ID != 0 {
				continue
			}
			new := model.Link{}
			new.RoomId = result1[0][0]
			new.Status = 1
			new.Link = link
			new.Type, _ = strconv.Atoi(kind)
			new.CreatedAt = time.Now().Unix()
			new.UpdatedAt = time.Now().Unix()
			mysql.DB.Save(&new)

		}
	}
	util.JsonWrite(context, 200, nil, "插入成功")

}

func GetDyId(link string) string {

	if link == "" {
		return ""
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(link)

	if err != nil {
		panic(err)
		return ""

	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

/**
  获取 抖音所有的 链接查询
*/

func GetLinksList(context *gin.Context) {

	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	Db := mysql.DB
	DB2 := mysql.DB
	var total int = 0

	links := make([]model.Link, 0)

	if status, isExist := context.GetQuery("status"); isExist == true {
		status, _ := strconv.Atoi(status)
		Db = Db.Where("status=?", status)
		DB2 = DB2.Model(links).Where("status=?", status)
	}

	Db.Limit(limit).Offset((page - 1) * limit).Order("updated_at desc")
	DB2.Count(&total)
	if err := Db.Find(&links).Error; err != nil {
		util.JsonWrite(context, -101, nil, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":   1,
		"count":  total,
		"result": links,
	})
	return

}
