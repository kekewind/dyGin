/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

type Config struct {
	ID      uint `gorm:"primaryKey;comment:'主键'"`
	Version int  `gorm:"comment:'脚本的版本号'"`
}
