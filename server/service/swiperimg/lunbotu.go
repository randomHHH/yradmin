
package swiperimg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg"
    swiperimgReq "github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg/request"
)

type LunbotuService struct {}
// CreateLunbotu 创建轮播图记录
// Author [yourname](https://github.com/yourname)
func (lunbotuService *LunbotuService) CreateLunbotu(ctx context.Context, lunbotu *swiperimg.Lunbotu) (err error) {
	err = global.GVA_DB.Create(lunbotu).Error
	return err
}

// DeleteLunbotu 删除轮播图记录
// Author [yourname](https://github.com/yourname)
func (lunbotuService *LunbotuService)DeleteLunbotu(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&swiperimg.Lunbotu{},"id = ?",ID).Error
	return err
}

// DeleteLunbotuByIds 批量删除轮播图记录
// Author [yourname](https://github.com/yourname)
func (lunbotuService *LunbotuService)DeleteLunbotuByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]swiperimg.Lunbotu{},"id in ?",IDs).Error
	return err
}

// UpdateLunbotu 更新轮播图记录
// Author [yourname](https://github.com/yourname)
func (lunbotuService *LunbotuService)UpdateLunbotu(ctx context.Context, lunbotu swiperimg.Lunbotu) (err error) {
	err = global.GVA_DB.Model(&swiperimg.Lunbotu{}).Where("id = ?",lunbotu.ID).Updates(&lunbotu).Error
	return err
}

// GetLunbotu 根据ID获取轮播图记录
// Author [yourname](https://github.com/yourname)
func (lunbotuService *LunbotuService)GetLunbotu(ctx context.Context, ID string) (lunbotu swiperimg.Lunbotu, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&lunbotu).Error
	return
}
// GetLunbotuInfoList 分页获取轮播图记录
// Author [yourname](https://github.com/yourname)
func (lunbotuService *LunbotuService)GetLunbotuInfoList(ctx context.Context, info swiperimgReq.LunbotuSearch) (list []swiperimg.Lunbotu, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&swiperimg.Lunbotu{})
    var lunbotus []swiperimg.Lunbotu
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

	err = db.Find(&lunbotus).Error
	return  lunbotus, total, err
}
func (lunbotuService *LunbotuService)GetLunbotuPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
