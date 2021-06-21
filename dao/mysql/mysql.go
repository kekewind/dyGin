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
	//数据迁移
	DB.AutoMigrate(&model.Link{})
	return err

}

func Close() {
	defer DB.Close()
}
