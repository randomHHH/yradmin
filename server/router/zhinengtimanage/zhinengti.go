package zhinengtimanage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZhinengtiRouter struct{}

// InitZhinengtiRouter 初始化 智能体管理 路由信息
func (s *ZhinengtiRouter) InitZhinengtiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	zhinengtiRouter := Router.Group("zhinengti").Use(middleware.OperationRecord())
	zhinengtiRouterWithoutRecord := Router.Group("zhinengti")
	zhinengtiRouterWithoutAuth := PublicRouter.Group("zhinengti")
	{
		zhinengtiRouter.POST("createZhinengti", zhinengtiApi.CreateZhinengti)             // 新建智能体管理
		zhinengtiRouter.DELETE("deleteZhinengti", zhinengtiApi.DeleteZhinengti)           // 删除智能体管理
		zhinengtiRouter.DELETE("deleteZhinengtiByIds", zhinengtiApi.DeleteZhinengtiByIds) // 批量删除智能体管理
		zhinengtiRouter.PUT("updateZhinengti", zhinengtiApi.UpdateZhinengti)              // 更新智能体管理
	}
	{
		zhinengtiRouterWithoutRecord.GET("findZhinengti", zhinengtiApi.FindZhinengti) // 根据ID获取智能体管理
		//zhinengtiRouterWithoutRecord.GET("getZhinengtiList", zhinengtiApi.GetZhinengtiList)  // 获取智能体管理列表
	}
	{
		zhinengtiRouterWithoutAuth.GET("getZhinengtiPublic", zhinengtiApi.GetZhinengtiPublic) // 智能体管理开放接口
		zhinengtiRouterWithoutAuth.GET("getZhinengtiList", zhinengtiApi.GetZhinengtiList)     // 获取智能体管理列表

	}
}
