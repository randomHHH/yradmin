// 自动生成模板YrUser
package yunrong

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户 结构体  YrUser
type YrUser struct {
	global.GVA_MODEL
	Username string `json:"username" form:"username" gorm:"comment:用户名;column:username;" binding:"required"` //用户名
	Password string `json:"password" form:"password" gorm:"column:password;" binding:"required"`             //密码
}

// TableName 用户 YrUser自定义表名 yr_user
func (YrUser) TableName() string {
	return "yr_user"
}
