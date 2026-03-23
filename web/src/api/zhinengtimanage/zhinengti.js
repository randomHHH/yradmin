import service from '@/utils/request'
// @Tags Zhinengti
// @Summary 创建智能体管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Zhinengti true "创建智能体管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /zhinengti/createZhinengti [post]
export const createZhinengti = (data) => {
  return service({
    url: '/zhinengti/createZhinengti',
    method: 'post',
    data
  })
}

// @Tags Zhinengti
// @Summary 删除智能体管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Zhinengti true "删除智能体管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhinengti/deleteZhinengti [delete]
export const deleteZhinengti = (params) => {
  return service({
    url: '/zhinengti/deleteZhinengti',
    method: 'delete',
    params
  })
}

// @Tags Zhinengti
// @Summary 批量删除智能体管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除智能体管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhinengti/deleteZhinengti [delete]
export const deleteZhinengtiByIds = (params) => {
  return service({
    url: '/zhinengti/deleteZhinengtiByIds',
    method: 'delete',
    params
  })
}

// @Tags Zhinengti
// @Summary 更新智能体管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Zhinengti true "更新智能体管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zhinengti/updateZhinengti [put]
export const updateZhinengti = (data) => {
  return service({
    url: '/zhinengti/updateZhinengti',
    method: 'put',
    data
  })
}

// @Tags Zhinengti
// @Summary 用id查询智能体管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Zhinengti true "用id查询智能体管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zhinengti/findZhinengti [get]
export const findZhinengti = (params) => {
  return service({
    url: '/zhinengti/findZhinengti',
    method: 'get',
    params
  })
}

// @Tags Zhinengti
// @Summary 分页获取智能体管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取智能体管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhinengti/getZhinengtiList [get]
export const getZhinengtiList = (params) => {
  return service({
    url: '/zhinengti/getZhinengtiList',
    method: 'get',
    params
  })
}

// @Tags Zhinengti
// @Summary 不需要鉴权的智能体管理接口
// @Accept application/json
// @Produce application/json
// @Param data query zhinengtimanageReq.ZhinengtiSearch true "分页获取智能体管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /zhinengti/getZhinengtiPublic [get]
export const getZhinengtiPublic = () => {
  return service({
    url: '/zhinengti/getZhinengtiPublic',
    method: 'get',
  })
}
