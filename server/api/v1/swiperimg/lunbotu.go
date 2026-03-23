package swiperimg

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg"
    swiperimgReq "github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type LunbotuApi struct {}



// CreateLunbotu 创建轮播图
// @Tags Lunbotu
// @Summary 创建轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body swiperimg.Lunbotu true "创建轮播图"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /lunbotu/createLunbotu [post]
func (lunbotuApi *LunbotuApi) CreateLunbotu(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var lunbotu swiperimg.Lunbotu
	err := c.ShouldBindJSON(&lunbotu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = lunbotuService.CreateLunbotu(ctx,&lunbotu)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteLunbotu 删除轮播图
// @Tags Lunbotu
// @Summary 删除轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body swiperimg.Lunbotu true "删除轮播图"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /lunbotu/deleteLunbotu [delete]
func (lunbotuApi *LunbotuApi) DeleteLunbotu(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := lunbotuService.DeleteLunbotu(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLunbotuByIds 批量删除轮播图
// @Tags Lunbotu
// @Summary 批量删除轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /lunbotu/deleteLunbotuByIds [delete]
func (lunbotuApi *LunbotuApi) DeleteLunbotuByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := lunbotuService.DeleteLunbotuByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateLunbotu 更新轮播图
// @Tags Lunbotu
// @Summary 更新轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body swiperimg.Lunbotu true "更新轮播图"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /lunbotu/updateLunbotu [put]
func (lunbotuApi *LunbotuApi) UpdateLunbotu(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var lunbotu swiperimg.Lunbotu
	err := c.ShouldBindJSON(&lunbotu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = lunbotuService.UpdateLunbotu(ctx,lunbotu)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindLunbotu 用id查询轮播图
// @Tags Lunbotu
// @Summary 用id查询轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询轮播图"
// @Success 200 {object} response.Response{data=swiperimg.Lunbotu,msg=string} "查询成功"
// @Router /lunbotu/findLunbotu [get]
func (lunbotuApi *LunbotuApi) FindLunbotu(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	relunbotu, err := lunbotuService.GetLunbotu(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(relunbotu, c)
}
// GetLunbotuList 分页获取轮播图列表
// @Tags Lunbotu
// @Summary 分页获取轮播图列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query swiperimgReq.LunbotuSearch true "分页获取轮播图列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /lunbotu/getLunbotuList [get]
func (lunbotuApi *LunbotuApi) GetLunbotuList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo swiperimgReq.LunbotuSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lunbotuService.GetLunbotuInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetLunbotuPublic 不需要鉴权的轮播图接口
// @Tags Lunbotu
// @Summary 不需要鉴权的轮播图接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lunbotu/getLunbotuPublic [get]
func (lunbotuApi *LunbotuApi) GetLunbotuPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    lunbotuService.GetLunbotuPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的轮播图接口信息",
    }, "获取成功", c)
}
