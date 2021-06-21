package model

/**
link 链接
*/

type Link struct {
	ID        uint   `gorm:"primaryKey;comment:'主键'"`
	Status    int    `gorm:"comment:'链接的状态 1 正常 2 正在使用 3 已使用'"`
	RoomId    string `gorm:"comment:'抖音视频 抖音个人信息提取的id'"`
	Link      string `gorm:"comment:'提交的抖音链接'"`
	Type      int    `gorm:"comment:'提交抖音链接的类型 1直播间的链接 2是个人作品的链接';size:1"`
	UseUser   string  `gorm:"comment:'使用链接的手机'"`
	BelongUser   string  `gorm:"comment:'采集链接的手机'"`
	CreatedAt string  `gorm:"comment:'创建时间'"`
	UpdatedAt string  `gorm:"comment:'更新时间'"`
}
