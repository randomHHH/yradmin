package yunrong

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/yunrong"
	yunrongReq "github.com/flipped-aurora/gin-vue-admin/server/model/yunrong/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type YrUserService struct{}

// CreateYrUser 创建用户记录
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) CreateYrUser(ctx context.Context, yrUser *yunrong.YrUser) (err error) {
	var count int64
	err = global.GVA_DB.Model(&yunrong.YrUser{}).Where("username = ?", yrUser.Username).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已注册")
	}
	hashPwd := utils.BcryptHash(yrUser.Password)
	yrUser.Password = hashPwd

	err = global.GVA_DB.Create(yrUser).Error
	return err
}

// DeleteYrUser 删除用户记录
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) DeleteYrUser(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&yunrong.YrUser{}, "id = ?", ID).Error
	return err
}

// DeleteYrUserByIds 批量删除用户记录
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) DeleteYrUserByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]yunrong.YrUser{}, "id in ?", IDs).Error
	return err
}

// UpdateYrUser 更新用户记录
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) UpdateYrUser(ctx context.Context, yrUser yunrong.YrUser) (err error) {
	err = global.GVA_DB.Model(&yunrong.YrUser{}).Where("id = ?", yrUser.ID).Updates(&yrUser).Error
	return err
}

// GetYrUser 根据ID获取用户记录
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) GetYrUser(ctx context.Context, ID string) (yrUser yunrong.YrUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&yrUser).Error
	return
}

// GetYrUserInfoList 分页获取用户记录
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) GetYrUserInfoList(ctx context.Context, info yunrongReq.YrUserSearch) (list []yunrong.YrUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&yunrong.YrUser{})
	var yrUsers []yunrong.YrUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Username != nil && *info.Username != "" {
		db = db.Where("username LIKE ?", "%"+*info.Username+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&yrUsers).Error
	return yrUsers, total, err
}
func (yrUserService *YrUserService) GetYrUserPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// AdminChangePassword 后台修改密码
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) AdminChangePassword(ctx context.Context, req yunrongReq.AdminChangePassword) (err error) {
	newPwd := utils.BcryptHash(req.Password)
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&yunrong.YrUser{}).Where("id=?", req.UserID).Update("password", newPwd)
	return db.Error
}

// Login 用户登陆
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) Login(ctx context.Context, username string, password string) (user yunrong.YrUser, err error) {
	// 请在这里实现自己的业务逻辑
	err = global.GVA_DB.First(&user, "username = ?", username).Error
	if err != nil {
		return user, errors.New("用户不存在")
	}
	ok := utils.BcryptCheck(password, user.Password)
	if !ok {
		return user, errors.New("密码错误")
	}
	return
}

// Register 用户注册
// Author [yourname](https://github.com/yourname)
func (yrUserService *YrUserService) Register(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&yunrong.YrUser{})
	return db.Error
}
