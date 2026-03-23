import service from '@/utils/request'
// @Tags Lunbotu
// @Summary 创建轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Lunbotu true "创建轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /lunbotu/createLunbotu [post]
export const createLunbotu = (data) => {
  return service({
    url: '/lunbotu/createLunbotu',
    method: 'post',
    data
  })
}

// @Tags Lunbotu
// @Summary 删除轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Lunbotu true "删除轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lunbotu/deleteLunbotu [delete]
export const deleteLunbotu = (params) => {
  return service({
    url: '/lunbotu/deleteLunbotu',
    method: 'delete',
    params
  })
}

// @Tags Lunbotu
// @Summary 批量删除轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lunbotu/deleteLunbotu [delete]
export const deleteLunbotuByIds = (params) => {
  return service({
    url: '/lunbotu/deleteLunbotuByIds',
    method: 'delete',
    params
  })
}

// @Tags Lunbotu
// @Summary 更新轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Lunbotu true "更新轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lunbotu/updateLunbotu [put]
export const updateLunbotu = (data) => {
  return service({
    url: '/lunbotu/updateLunbotu',
    method: 'put',
    data
  })
}

// @Tags Lunbotu
// @Summary 用id查询轮播图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Lunbotu true "用id查询轮播图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lunbotu/findLunbotu [get]
export const findLunbotu = (params) => {
  return service({
    url: '/lunbotu/findLunbotu',
    method: 'get',
    params
  })
}

// @Tags Lunbotu
// @Summary 分页获取轮播图列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取轮播图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lunbotu/getLunbotuList [get]
export const getLunbotuList = (params) => {
  return service({
    url: '/lunbotu/getLunbotuList',
    method: 'get',
    params
  })
}

// @Tags Lunbotu
// @Summary 不需要鉴权的轮播图接口
// @Accept application/json
// @Produce application/json
// @Param data query swiperimgReq.LunbotuSearch true "分页获取轮播图列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lunbotu/getLunbotuPublic [get]
export const getLunbotuPublic = () => {
  return service({
    url: '/lunbotu/getLunbotuPublic',
    method: 'get',
  })
}
