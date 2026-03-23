package swiperimg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LunbotuRouter struct{}

// InitLunbotuRouter 初始化 轮播图 路由信息
func (s *LunbotuRouter) InitLunbotuRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	lunbotuRouter := Router.Group("lunbotu").Use(middleware.OperationRecord())
	lunbotuRouterWithoutRecord := Router.Group("lunbotu")
	lunbotuRouterWithoutAuth := PublicRouter.Group("lunbotu")
	{
		lunbotuRouter.POST("createLunbotu", lunbotuApi.CreateLunbotu)             // 新建轮播图
		lunbotuRouter.DELETE("deleteLunbotu", lunbotuApi.DeleteLunbotu)           // 删除轮播图
		lunbotuRouter.DELETE("deleteLunbotuByIds", lunbotuApi.DeleteLunbotuByIds) // 批量删除轮播图
		lunbotuRouter.PUT("updateLunbotu", lunbotuApi.UpdateLunbotu)              // 更新轮播图
	}
	{
		lunbotuRouterWithoutRecord.GET("findLunbotu", lunbotuApi.FindLunbotu) // 根据ID获取轮播图
	}
	{
		lunbotuRouterWithoutAuth.GET("getLunbotuPublic", lunbotuApi.GetLunbotuPublic) // 轮播图开放接口
		lunbotuRouterWithoutAuth.GET("getLunbotuList", lunbotuApi.GetLunbotuList)     // 获取轮播图列表

	}
}
