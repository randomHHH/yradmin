import service from '@/utils/request'
// @Tags YrUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.YrUser true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /yrUser/createYrUser [post]
export const createYrUser = (data) => {
  return service({
    url: '/yrUser/createYrUser',
    method: 'post',
    data
  })
}

// @Tags YrUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.YrUser true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yrUser/deleteYrUser [delete]
export const deleteYrUser = (params) => {
  return service({
    url: '/yrUser/deleteYrUser',
    method: 'delete',
    params
  })
}

// @Tags YrUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yrUser/deleteYrUser [delete]
export const deleteYrUserByIds = (params) => {
  return service({
    url: '/yrUser/deleteYrUserByIds',
    method: 'delete',
    params
  })
}

// @Tags YrUser
// @Summary 更新用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.YrUser true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yrUser/updateYrUser [put]
export const updateYrUser = (data) => {
  return service({
    url: '/yrUser/updateYrUser',
    method: 'put',
    data
  })
}

// @Tags YrUser
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.YrUser true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yrUser/findYrUser [get]
export const findYrUser = (params) => {
  return service({
    url: '/yrUser/findYrUser',
    method: 'get',
    params
  })
}

// @Tags YrUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yrUser/getYrUserList [get]
export const getYrUserList = (params) => {
  return service({
    url: '/yrUser/getYrUserList',
    method: 'get',
    params
  })
}

// @Tags YrUser
// @Summary 不需要鉴权的用户接口
// @Accept application/json
// @Produce application/json
// @Param data query yunrongReq.YrUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /yrUser/getYrUserPublic [get]
export const getYrUserPublic = () => {
  return service({
    url: '/yrUser/getYrUserPublic',
    method: 'get',
  })
}
// AdminChangePassword 后台修改密码
// @Tags YrUser
// @Summary 后台修改密码
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /yrUser/adminChangePassword [PUT]
export const adminChangePassword = (data) => {
  return service({
    url: '/yrUser/adminChangePassword',
    method: 'PUT',
    data
  })
}
// Login 用户登陆
// @Tags YrUser
// @Summary 用户登陆
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /yrUser/login [POST]
export const login = () => {
  return service({
    url: '/yrUser/login',
    method: 'POST'
  })
}
// Register 用户注册
// @Tags YrUser
// @Summary 用户注册
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /yrUser/register [POST]
export const register = () => {
  return service({
    url: '/yrUser/register',
    method: 'POST'
  })
}
