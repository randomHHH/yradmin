package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/swiperimg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/yunrong"
	"github.com/flipped-aurora/gin-vue-admin/server/router/zhinengti"
	"github.com/flipped-aurora/gin-vue-admin/server/router/zhinengtimanage"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	Yunrong         yunrong.RouterGroup
	Zhinengti       zhinengti.RouterGroup
	Zhinengtimanage zhinengtimanage.RouterGroup
	Swiperimg       swiperimg.RouterGroup
}
