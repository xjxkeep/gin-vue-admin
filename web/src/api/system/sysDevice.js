import service from '@/utils/request'
// @Tags Device
// @Summary 创建设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Device true "创建设备管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /DM/createDevice [post]
export const createDevice = (data) => {
  return service({
    url: '/DM/createDevice',
    method: 'post',
    data
  })
}

// @Tags Device
// @Summary 删除设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Device true "删除设备管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DM/deleteDevice [delete]
export const deleteDevice = (params) => {
  return service({
    url: '/DM/deleteDevice',
    method: 'delete',
    params
  })
}

// @Tags Device
// @Summary 批量删除设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除设备管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DM/deleteDevice [delete]
export const deleteDeviceByIds = (params) => {
  return service({
    url: '/DM/deleteDeviceByIds',
    method: 'delete',
    params
  })
}

// @Tags Device
// @Summary 更新设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Device true "更新设备管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DM/updateDevice [put]
export const updateDevice = (data) => {
  return service({
    url: '/DM/updateDevice',
    method: 'put',
    data
  })
}

// @Tags Device
// @Summary 用id查询设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Device true "用id查询设备管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DM/findDevice [get]
export const findDevice = (params) => {
  return service({
    url: '/DM/findDevice',
    method: 'get',
    params
  })
}

// @Tags Device
// @Summary 分页获取设备管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取设备管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DM/getDeviceList [get]
export const getDeviceList = (params) => {
  return service({
    url: '/DM/getDeviceList',
    method: 'get',
    params
  })
}

// @Tags Device
// @Summary 不需要鉴权的设备管理接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.DeviceSearch true "分页获取设备管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DM/getDevicePublic [get]
export const getDevicePublic = () => {
  return service({
    url: '/DM/getDevicePublic',
    method: 'get',
  })
}
