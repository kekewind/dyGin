/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

type Collect struct {
	ID        uint   `gorm:"primaryKey;comment:'主键'"`
	Status    int    `gorm:"comment:'链接的状态 1 正常 2 正在使用 3 已使用'"`
	Age       string `gorm:"comment:'年龄'"`
	Address   string `gorm:"comment:'地址'"`
	Like      int    `gorm:"comment:'获赞'"`
	Following int    `gorm:"comment:'关注'"`
	Follows   int    `gorm:"comment:'粉丝'"`
	Sign      string `gorm:"comment:'签名'"`
	Sex       string `gorm:"comment:'性别'"`
	Type      int    `gorm:"comment:'调教抖音链接的类型 1直播间的链接 2是个人作品的链接';size:1"`
	DyNumber  string  `gorm:"comment:'抖音号'"`
	UseUser   string `gorm:"comment:'使用链接的手机'"`
	CollectUser   string `gorm:"comment:'采集链接的手机'"`
	CreatedAt string  `gorm:"comment:'创建时间'"`
	UpdatedAt string  `gorm:"comment:'更新时间'"`
}
