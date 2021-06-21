package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
工具
*/

//随机获取 36个位 字符串 生成token

//封装方法 判断数组中是否存在某一个元素
func InArray(str []string, arg1 string) bool {
	for _, element := range str {
		if element == arg1 {
			return true
		}
	}
	return false
}

func JsonWrite(context *gin.Context, status int, result interface{}, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code":   status,
		"result": result,
		"msg":    msg,
	})
}


//时间戳转时间
func UnixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

//时间转时间戳
func StrToUnix(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}

