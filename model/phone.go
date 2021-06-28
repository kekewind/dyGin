/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

type Phone struct {
	ID       uint `gorm:"primaryKey;comment:'主键'"`
	Nickname string
	Status   int `gorm:"comment:1 未运行 2运行中`
	TaskType int `gorm:"comment:1采集 2直播加粉 3短视频粉 4 私聊`
	Created  int64  `gorm:"comment:'创建时间'"`
	Updated  int64  `gorm:"comment:'更新时间'"`
}
