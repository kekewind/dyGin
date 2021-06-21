package controller



// 注册参数 检验的定义
type checkRegister struct {
	AdminName     string `form:"admin_name" json:"admin_name" xml:"admin_name" binding:"required"`
	AdminPassword string `form:"admin_password" json:"admin_password" xml:"admin_password" binding:"required"`
	Token         string
	Price         float32
	PriceTwo      float32
	QqNum         string
	QqLink        string
	CreatedAt     int
	UpdatedAt     int
}

/**
管理员
*/


