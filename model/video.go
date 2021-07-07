/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

type Video struct {
	ID      uint `gorm:"primaryKey;comment:'主键'"`
	PathUrl int  `gorm:"comment:'文件路径'"`
	Status int  `gorm:"comment:'状态'"`
}

