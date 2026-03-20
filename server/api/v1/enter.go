package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/yunrong"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/zhinengti"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/zhinengtimanage"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup          system.ApiGroup
	ExampleApiGroup         example.ApiGroup
	YunrongApiGroup         yunrong.ApiGroup
	ZhinengtiApiGroup       zhinengti.ApiGroup
	ZhinengtimanageApiGroup zhinengtimanage.ApiGroup
}
