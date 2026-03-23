package swiperimg

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg"
    swiperimgReq "github.com/flipped-aurora/gin-vue-admin/server/model/swiperimg/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SwiperApi struct {}



// CreateSwiper 创建轮播图
// @Tags Swiper
// @Summary 创建轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body swiperimg.Swiper true "创建轮播图"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /swiper/createSwiper [post]
func (swiperApi *SwiperApi) CreateSwiper(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var swiper swiperimg.Swiper
	err := c.ShouldBindJSON(&swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = swiperService.CreateSwiper(ctx,&swiper)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSwiper 删除轮播图
// @Tags Swiper
// @Summary 删除轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body swiperimg.Swiper true "删除轮播图"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /swiper/deleteSwiper [delete]
func (swiperApi *SwiperApi) DeleteSwiper(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := swiperService.DeleteSwiper(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSwiperByIds 批量删除轮播图
// @Tags Swiper
// @Summary 批量删除轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /swiper/deleteSwiperByIds [delete]
func (swiperApi *SwiperApi) DeleteSwiperByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := swiperService.DeleteSwiperByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSwiper 更新轮播图
// @Tags Swiper
// @Summary 更新轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body swiperimg.Swiper true "更新轮播图"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /swiper/updateSwiper [put]
func (swiperApi *SwiperApi) UpdateSwiper(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var swiper swiperimg.Swiper
	err := c.ShouldBindJSON(&swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = swiperService.UpdateSwiper(ctx,swiper)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSwiper 用id查询轮播图
// @Tags Swiper
// @Summary 用id查询轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询轮播图"
// @Success 200 {object} response.Response{data=swiperimg.Swiper,msg=string} "查询成功"
// @Router /swiper/findSwiper [get]
func (swiperApi *SwiperApi) FindSwiper(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reswiper, err := swiperService.GetSwiper(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reswiper, c)
}
// GetSwiperList 分页获取轮播图列表
// @Tags Swiper
// @Summary 分页获取轮播图列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query swiperimgReq.SwiperSearch true "分页获取轮播图列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /swiper/getSwiperList [get]
func (swiperApi *SwiperApi) GetSwiperList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo swiperimgReq.SwiperSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := swiperService.GetSwiperInfoList(ctx,pageInfo)
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

// GetSwiperPublic 不需要鉴权的轮播图接口
// @Tags Swiper
// @Summary 不需要鉴权的轮播图接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /swiper/getSwiperPublic [get]
func (swiperApi *SwiperApi) GetSwiperPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    swiperService.GetSwiperPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的轮播图接口信息",
    }, "获取成功", c)
}
