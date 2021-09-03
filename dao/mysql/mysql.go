package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/wangyi/dyGin/model"
)

var (
	DB  *gorm.DB
	err error
)

func Init() error {
	// 参考 https://github.com/go-sql-driver/mysql#d	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情sn-data-source-name 获取详情
	//dsn := "root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err)
		panic(err)
		return err
	}
	//设置链接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	//日志
	//DB.LogMode(true)
	//数据迁移
	//DB.AutoMigrate(&model.Link{})
	//DB.AutoMigrate(&model.Collect{})
	//DB.AutoMigrate(&model.Admin{})
	//DB.AutoMigrate(&model.Phone{})

	CheckIsExistModelCollect()
	CheckIsExistModelLink()
	CheckIsExistModelPhone()
	CheckIsExistModelConfig()
	CheckIsExistModelUserId()

	//CheckIsExistModel(&model.Link{})
	//CheckIsExistModel(&model.Admin{})
	//CheckIsExistModel(&model.Collect{})
	return err

}

func Close() {
	defer DB.Close()
}

//检查数据是否存在不存在 自己创建一个
func CheckIsExistModelCollect() {
	if DB.HasTable(&model.Collect{}) {
		fmt.Println("数据库已经存在了!")
		DB.AutoMigrate(&model.Collect{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		DB.CreateTable(&model.Collect{})
	}
}

func CheckIsExistModelLink() {
	if DB.HasTable(&model.Link{}) {
		fmt.Println("数据库已经存在了!")
		DB.AutoMigrate(&model.Link{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		DB.CreateTable(&model.Link{})
	}
}
func CheckIsExistModelPhone() {
	if DB.HasTable(&model.Phone{}) {
		fmt.Println("数据库已经存在了!")
		DB.AutoMigrate(&model.Phone{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		DB.CreateTable(&model.Phone{})
	}
}
func CheckIsExistModelConfig() {
	if DB.HasTable(&model.Config{}) {
		fmt.Println("数据库已经存在了!")
		DB.AutoMigrate(&model.Config{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		DB.CreateTable(&model.Config{})
	}
}
func CheckIsExistModelUserId() {
	if DB.HasTable(&model.UserId{}) {
		fmt.Println("数据库已经存在了!")
		DB.AutoMigrate(&model.UserId{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		DB.CreateTable(&model.UserId{})
	}
}