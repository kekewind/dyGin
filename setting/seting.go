package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	// 指定配置文件 如果是json 就写json
	viper.SetConfigFile("config.yaml")
	//指定查找配置文件的路径
	viper.AddConfigPath(".")
	//读取配置文件
	err = viper.ReadInConfig()
	// 读取配置信息失败
	if err != nil {
		fmt.Println("读取配置信息失败:", err)
		return err
	}

	//监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了!")
	})


     print("值;"+viper.GetString("mysql.user"))


	fmt.Println("获取配置成功 \n")
	return err

}
