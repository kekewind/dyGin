package util

import (
	"fmt"
	"testing"
	"time"
)

//func TestName(t *testing.T) {
//	redis.Init()
//	fmt.Println(SetToken(36))
//
//}

func TestOne(t *testing.T) {

	var a = []string{"小米", "小黑", "小白"}
	fmt.Println(InArray(a, "小白"))
	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05","时间", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())
}



