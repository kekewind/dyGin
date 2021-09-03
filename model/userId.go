/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model


type UserId struct {
	ID        uint   `gorm:"primaryKey;comment:'主键'"`
	UserId      string `gorm:"comment:'抖音用户的id'"`
	Status    int    `gorm:"comment:'链接的状态 1 正常 2 已使用'"`
	UseUser   string  `gorm:"comment:'使用链接的手机'"`
	Created  int64  `gorm:"comment:'创建时间'"`
	Updated  int64  `gorm:"comment:'更新时间'"`
}
