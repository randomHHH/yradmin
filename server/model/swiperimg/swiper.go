
// 自动生成模板Swiper
package swiperimg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 轮播图 结构体  Swiper
type Swiper struct {
    global.GVA_MODEL
  Title  *string `json:"title" form:"title" gorm:"column:title;"`  //名称
  Urlpath  string `json:"urlpath" form:"urlpath" gorm:"column:urlpath;"`  //图片地址
  Desc  *string `json:"desc" form:"desc" gorm:"column:desc;"`  //图片描述
}


// TableName 轮播图 Swiper自定义表名 yr_swiper
func (Swiper) TableName() string {
    return "yr_swiper"
}





