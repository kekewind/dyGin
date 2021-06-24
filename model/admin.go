/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

type Admin struct {
	ID            uint `gorm:"primaryKey;comment:'主键'"`
	AdminUser     string
	AdminPassword string
	UpdatedAt     int64 `gorm:"autoUpdateTime"`
	Update        int64
}
