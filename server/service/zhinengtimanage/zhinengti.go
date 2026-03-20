
package zhinengtimanage

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/zhinengtimanage"
    zhinengtimanageReq "github.com/flipped-aurora/gin-vue-admin/server/model/zhinengtimanage/request"
)

type ZhinengtiService struct {}
// CreateZhinengti 创建智能体管理记录
// Author [yourname](https://github.com/yourname)
func (zhinengtiService *ZhinengtiService) CreateZhinengti(ctx context.Context, zhinengti *zhinengtimanage.Zhinengti) (err error) {
	err = global.GVA_DB.Create(zhinengti).Error
	return err
}

// DeleteZhinengti 删除智能体管理记录
// Author [yourname](https://github.com/yourname)
func (zhinengtiService *ZhinengtiService)DeleteZhinengti(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&zhinengtimanage.Zhinengti{},"id = ?",ID).Error
	return err
}

// DeleteZhinengtiByIds 批量删除智能体管理记录
// Author [yourname](https://github.com/yourname)
func (zhinengtiService *ZhinengtiService)DeleteZhinengtiByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]zhinengtimanage.Zhinengti{},"id in ?",IDs).Error
	return err
}

// UpdateZhinengti 更新智能体管理记录
// Author [yourname](https://github.com/yourname)
func (zhinengtiService *ZhinengtiService)UpdateZhinengti(ctx context.Context, zhinengti zhinengtimanage.Zhinengti) (err error) {
	err = global.GVA_DB.Model(&zhinengtimanage.Zhinengti{}).Where("id = ?",zhinengti.ID).Updates(&zhinengti).Error
	return err
}

// GetZhinengti 根据ID获取智能体管理记录
// Author [yourname](https://github.com/yourname)
func (zhinengtiService *ZhinengtiService)GetZhinengti(ctx context.Context, ID string) (zhinengti zhinengtimanage.Zhinengti, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&zhinengti).Error
	return
}
// GetZhinengtiInfoList 分页获取智能体管理记录
// Author [yourname](https://github.com/yourname)
func (zhinengtiService *ZhinengtiService)GetZhinengtiInfoList(ctx context.Context, info zhinengtimanageReq.ZhinengtiSearch) (list []zhinengtimanage.Zhinengti, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&zhinengtimanage.Zhinengti{})
    var zhinengtis []zhinengtimanage.Zhinengti
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title LIKE ?", "%"+ *info.Title+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&zhinengtis).Error
	return  zhinengtis, total, err
}
func (zhinengtiService *ZhinengtiService)GetZhinengtiPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
