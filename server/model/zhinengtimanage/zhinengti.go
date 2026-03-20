// 自动生成模板Zhinengti
package zhinengtimanage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 智能体管理 结构体  Zhinengti
type Zhinengti struct {
	global.GVA_MODEL
	Title *string `json:"title" form:"title" gorm:"column:title;"` //标题
	Img   string  `json:"img" form:"img" gorm:"column:img;"`       //封面
	Desc  *string `json:"desc" form:"desc" gorm:"column:desc;"`    //描述
	Url   *string `json:"url" form:"url" gorm:"column:url;"`       //链接地址

}

// TableName 智能体管理 Zhinengti自定义表名 zhinengti
func (Zhinengti) TableName() string {
	return "yr_zhinengti"
}
