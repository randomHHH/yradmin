package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/swiperimg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/yunrong"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zhinengti"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zhinengtimanage"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup          system.ServiceGroup
	ExampleServiceGroup         example.ServiceGroup
	YunrongServiceGroup         yunrong.ServiceGroup
	ZhinengtiServiceGroup       zhinengti.ServiceGroup
	ZhinengtimanageServiceGroup zhinengtimanage.ServiceGroup
	SwiperimgServiceGroup       swiperimg.ServiceGroup
}
