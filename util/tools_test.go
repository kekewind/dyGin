package util

import (
	"fmt"
	"github.com/wangyi/gin_project/dao/redis"
	"testing"
)

func TestName(t *testing.T) {
	redis.Init()
	fmt.Println(SetToken(36))

}

func TestOne(t *testing.T) {

	var a = []string{"小米", "小黑", "小白"}
	fmt.Println(InArray(a, "小白"))

}
