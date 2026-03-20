package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/yunrong"
	"github.com/flipped-aurora/gin-vue-admin/server/model/zhinengtimanage"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(yunrong.YrUser{}, zhinengtimanage.Zhinengti{})
	if err != nil {
		return err
	}
	return nil
}
