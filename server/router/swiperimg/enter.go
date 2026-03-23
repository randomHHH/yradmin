package swiperimg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	LunbotuRouter
	SwiperRouter
}

var (
	lunbotuApi = api.ApiGroupApp.SwiperimgApiGroup.LunbotuApi
	swiperApi  = api.ApiGroupApp.SwiperimgApiGroup.SwiperApi
)
