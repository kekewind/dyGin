package util

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/gin_project/dao/redis"
	"math/rand"
	"net/http"
	"time"
)

/**
工具
*/

//随机获取 36个位 字符串 生成token
func SetToken(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzQWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	//判断这个token redis 是否已经存在了?   5次循环
	for i := 0; i < 5; i++ {
		bExist, _ := redis.RDB.HExists("TOKEN", string(result)).Result()
		if !bExist {
			break
		}
	}
	return string(result)
}

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
