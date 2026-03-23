package swiperimg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SwiperRouter struct{}

// InitSwiperRouter 初始化 轮播图 路由信息
func (s *SwiperRouter) InitSwiperRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	swiperRouter := Router.Group("swiper").Use(middleware.OperationRecord())
	swiperRouterWithoutRecord := Router.Group("swiper")
	swiperRouterWithoutAuth := PublicRouter.Group("swiper")
	{
		swiperRouter.POST("createSwiper", swiperApi.CreateSwiper)             // 新建轮播图
		swiperRouter.DELETE("deleteSwiper", swiperApi.DeleteSwiper)           // 删除轮播图
		swiperRouter.DELETE("deleteSwiperByIds", swiperApi.DeleteSwiperByIds) // 批量删除轮播图
		swiperRouter.PUT("updateSwiper", swiperApi.UpdateSwiper)              // 更新轮播图
	}
	{
		swiperRouterWithoutRecord.GET("findSwiper", swiperApi.FindSwiper) // 根据ID获取轮播图
	}
	{
		swiperRouterWithoutAuth.GET("getSwiperPublic", swiperApi.GetSwiperPublic) // 轮播图开放接口
		swiperRouterWithoutAuth.GET("getSwiperList", swiperApi.GetSwiperList)     // 获取轮播图列表

	}
}
