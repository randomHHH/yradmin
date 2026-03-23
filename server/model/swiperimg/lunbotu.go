
// 自动生成模板Lunbotu
package swiperimg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 轮播图 结构体  Lunbotu
type Lunbotu struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"column:title;"`  //轮播图
  Urlpath  *string `json:"urlpath" form:"urlpath" gorm:"column:urlpath;"`  //图片地址
  Desc  *string `json:"desc" form:"desc" gorm:"column:desc;"`  //图片描述
}


// TableName 轮播图 Lunbotu自定义表名 lunbotu
func (Lunbotu) TableName() string {
    return "lunbotu"
}





