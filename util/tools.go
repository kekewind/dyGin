package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
