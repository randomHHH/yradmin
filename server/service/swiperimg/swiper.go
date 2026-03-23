
package swiperimg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg"
    swiperimgReq "github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg/request"
)

type SwiperService struct {}
// CreateSwiper 创建轮播图记录
// Author [yourname](https://github.com/yourname)
func (swiperService *SwiperService) CreateSwiper(ctx context.Context, swiper *swiperimg.Swiper) (err error) {
	err = global.GVA_DB.Create(swiper).Error
	return err
}

// DeleteSwiper 删除轮播图记录
// Author [yourname](https://github.com/yourname)
func (swiperService *SwiperService)DeleteSwiper(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&swiperimg.Swiper{},"id = ?",ID).Error
	return err
}

// DeleteSwiperByIds 批量删除轮播图记录
// Author [yourname](https://github.com/yourname)
func (swiperService *SwiperService)DeleteSwiperByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]swiperimg.Swiper{},"id in ?",IDs).Error
	return err
}

// UpdateSwiper 更新轮播图记录
// Author [yourname](https://github.com/yourname)
func (swiperService *SwiperService)UpdateSwiper(ctx context.Context, swiper swiperimg.Swiper) (err error) {
	err = global.GVA_DB.Model(&swiperimg.Swiper{}).Where("id = ?",swiper.ID).Updates(&swiper).Error
	return err
}

// GetSwiper 根据ID获取轮播图记录
// Author [yourname](https://github.com/yourname)
func (swiperService *SwiperService)GetSwiper(ctx context.Context, ID string) (swiper swiperimg.Swiper, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&swiper).Error
	return
}
// GetSwiperInfoList 分页获取轮播图记录
// Author [yourname](https://github.com/yourname)
func (swiperService *SwiperService)GetSwiperInfoList(ctx context.Context, info swiperimgReq.SwiperSearch) (list []swiperimg.Swiper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&swiperimg.Swiper{})
    var swipers []swiperimg.Swiper
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&swipers).Error
	return  swipers, total, err
}
func (swiperService *SwiperService)GetSwiperPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
