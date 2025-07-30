package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DeviceApi struct {}



// CreateDevice 创建设备管理
// @Tags Device
// @Summary 创建设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Device true "创建设备管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /DM/createDevice [post]
func (DMApi *DeviceApi) CreateDevice(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var DM system.Device
	err := c.ShouldBindJSON(&DM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DMService.CreateDevice(ctx,&DM)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDevice 删除设备管理
// @Tags Device
// @Summary 删除设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Device true "删除设备管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /DM/deleteDevice [delete]
func (DMApi *DeviceApi) DeleteDevice(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := DMService.DeleteDevice(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDeviceByIds 批量删除设备管理
// @Tags Device
// @Summary 批量删除设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /DM/deleteDeviceByIds [delete]
func (DMApi *DeviceApi) DeleteDeviceByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := DMService.DeleteDeviceByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDevice 更新设备管理
// @Tags Device
// @Summary 更新设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.Device true "更新设备管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /DM/updateDevice [put]
func (DMApi *DeviceApi) UpdateDevice(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var DM system.Device
	err := c.ShouldBindJSON(&DM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DMService.UpdateDevice(ctx,DM)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDevice 用id查询设备管理
// @Tags Device
// @Summary 用id查询设备管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询设备管理"
// @Success 200 {object} response.Response{data=system.Device,msg=string} "查询成功"
// @Router /DM/findDevice [get]
func (DMApi *DeviceApi) FindDevice(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reDM, err := DMService.GetDevice(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reDM, c)
}
// GetDeviceList 分页获取设备管理列表
// @Tags Device
// @Summary 分页获取设备管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.DeviceSearch true "分页获取设备管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /DM/getDeviceList [get]
func (DMApi *DeviceApi) GetDeviceList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.DeviceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DMService.GetDeviceInfoList(ctx,pageInfo)
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

// GetDevicePublic 不需要鉴权的设备管理接口
// @Tags Device
// @Summary 不需要鉴权的设备管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DM/getDevicePublic [get]
func (DMApi *DeviceApi) GetDevicePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    DMService.GetDevicePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的设备管理接口信息",
    }, "获取成功", c)
}
