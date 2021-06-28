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
	Status   string `gorm:"comment:1 未运行 2运行中`
	TaskType string `gorm:"comment:1采集 2直播加粉 3短视频粉 4 私聊`
	Create   int64  `gorm:"comment:'创建时间'"`
	Update   int64  `gorm:"comment:'更新时间'"`
}
