package swiperimg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	LunbotuApi
	SwiperApi
}

var (
	lunbotuService = service.ServiceGroupApp.SwiperimgServiceGroup.LunbotuService
	swiperService  = service.ServiceGroupApp.SwiperimgServiceGroup.SwiperService
)
