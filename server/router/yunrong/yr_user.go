package yunrong

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YrUserRouter struct{}

func (s *YrUserRouter) InitYrUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	yrUserRouter := Router.Group("yrUser").Use(middleware.OperationRecord())
	yrUserRouterWithoutRecord := Router.Group("yrUser")
	yrUserRouterWithoutAuth := PublicRouter.Group("yrUser")
	{
		yrUserRouter.POST("createYrUser", yrUserApi.CreateYrUser)
		yrUserRouter.DELETE("deleteYrUser", yrUserApi.DeleteYrUser)
		yrUserRouter.DELETE("deleteYrUserByIds", yrUserApi.DeleteYrUserByIds)
		yrUserRouter.PUT("updateYrUser", yrUserApi.UpdateYrUser)
	}
	{
		yrUserRouterWithoutRecord.GET("findYrUser", yrUserApi.FindYrUser)
		yrUserRouterWithoutRecord.GET("getYrUserList", yrUserApi.GetYrUserList)
	}
	{
		yrUserRouterWithoutAuth.GET("getYrUserPublic", yrUserApi.GetYrUserPublic)
		yrUserRouterWithoutAuth.PUT("adminChangePassword", yrUserApi.AdminChangePassword)
		yrUserRouterWithoutAuth.POST("login", yrUserApi.Login)
		yrUserRouterWithoutAuth.POST("register", yrUserApi.Register)
	}
}
