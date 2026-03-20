package yunrong

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/yunrong"
	yunrongReq "github.com/flipped-aurora/gin-vue-admin/server/model/yunrong/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type YrUserApi struct{}

// CreateYrUser 创建用户
// @Tags YrUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body yunrong.YrUser true "创建用户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /yrUser/createYrUser [post]
func (yrUserApi *YrUserApi) CreateYrUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var yrUser yunrong.YrUser
	err := c.ShouldBindJSON(&yrUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = yrUserService.CreateYrUser(ctx, &yrUser)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteYrUser 删除用户
// @Tags YrUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body yunrong.YrUser true "删除用户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /yrUser/deleteYrUser [delete]
func (yrUserApi *YrUserApi) DeleteYrUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := yrUserService.DeleteYrUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteYrUserByIds 批量删除用户
// @Tags YrUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /yrUser/deleteYrUserByIds [delete]
func (yrUserApi *YrUserApi) DeleteYrUserByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := yrUserService.DeleteYrUserByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateYrUser 更新用户
// @Tags YrUser
// @Summary 更新用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body yunrong.YrUser true "更新用户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /yrUser/updateYrUser [put]
func (yrUserApi *YrUserApi) UpdateYrUser(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var yrUser yunrong.YrUser
	err := c.ShouldBindJSON(&yrUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = yrUserService.UpdateYrUser(ctx, yrUser)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindYrUser 用id查询用户
// @Tags YrUser
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户"
// @Success 200 {object} response.Response{data=yunrong.YrUser,msg=string} "查询成功"
// @Router /yrUser/findYrUser [get]
func (yrUserApi *YrUserApi) FindYrUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reyrUser, err := yrUserService.GetYrUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reyrUser, c)
}

// GetYrUserList 分页获取用户列表
// @Tags YrUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query yunrongReq.YrUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /yrUser/getYrUserList [get]
func (yrUserApi *YrUserApi) GetYrUserList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo yunrongReq.YrUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := yrUserService.GetYrUserInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetYrUserPublic 不需要鉴权的用户接口
// @Tags YrUser
// @Summary 不需要鉴权的用户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /yrUser/getYrUserPublic [get]
func (yrUserApi *YrUserApi) GetYrUserPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	yrUserService.GetYrUserPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户接口信息",
	}, "获取成功", c)
}

// AdminChangePassword 后台修改密码
// @Tags YrUser
// @Summary 后台修改密码
// @Accept application/json
// @Produce application/json
// @Param data query yunrongReq.YrUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /yrUser/adminChangePassword [PUT]
func (yrUserApi *YrUserApi) AdminChangePassword(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 请添加自己的业务逻辑
	// 2. 定义接收参数的结构体（需提前在 yunrongReq 包中定义）
	var req yunrongReq.AdminChangePassword
	// 3. 解析 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数解析失败!", zap.Error(err))
		response.FailWithMessage("参数格式错误", c)
		return
	}

	err := yrUserService.AdminChangePassword(ctx, req)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// Login 用户登陆
// @Tags YrUser
// @Summary 用户登陆
// @Accept application/json
// @Produce application/json
// @Param data query yunrongReq.YrUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /yrUser/login [POST]
func (yrUserApi *YrUserApi) Login(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req yunrongReq.YrUserLogin
	// 3. 解析 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数格式错误", c)
		return
	}

	// 请添加自己的业务逻辑
	user, err := yrUserService.Login(ctx, req.UserName, req.Password)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}

	token, _, err := utils.LoginWithToken(user)

	response.OkWithData(gin.H{
		"token": token,
		"user":  user,
	}, c)
}

// Register 用户注册
// @Tags YrUser
// @Summary 用户注册
// @Accept application/json
// @Produce application/json
// @Param data query yunrongReq.YrUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /yrUser/register [POST]
func (yrUserApi *YrUserApi) Register(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	var req yunrong.YrUser
	// 3. 解析 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数格式错误", c)
		return
	}

	// 请添加自己的业务逻辑
	err := yrUserService.CreateYrUser(ctx, &req)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}
