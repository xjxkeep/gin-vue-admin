package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct {}

// InitDeviceRouter 初始化 设备管理 路由信息
func (s *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	DMRouter := Router.Group("DM").Use(middleware.OperationRecord())
	DMRouterWithoutRecord := Router.Group("DM")
	DMRouterWithoutAuth := PublicRouter.Group("DM")
	{
		DMRouter.POST("createDevice", DMApi.CreateDevice)   // 新建设备管理
		DMRouter.DELETE("deleteDevice", DMApi.DeleteDevice) // 删除设备管理
		DMRouter.DELETE("deleteDeviceByIds", DMApi.DeleteDeviceByIds) // 批量删除设备管理
		DMRouter.PUT("updateDevice", DMApi.UpdateDevice)    // 更新设备管理
	}
	{
		DMRouterWithoutRecord.GET("findDevice", DMApi.FindDevice)        // 根据ID获取设备管理
		DMRouterWithoutRecord.GET("getDeviceList", DMApi.GetDeviceList)  // 获取设备管理列表
	}
	{
	    DMRouterWithoutAuth.GET("getDevicePublic", DMApi.GetDevicePublic)  // 设备管理开放接口
	}
}
