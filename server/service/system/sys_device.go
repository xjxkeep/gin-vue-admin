
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type DeviceService struct {}
// CreateDevice 创建设备管理记录
// Author [yourname](https://github.com/yourname)
func (DMService *DeviceService) CreateDevice(ctx context.Context, DM *system.Device) (err error) {
	err = global.GVA_DB.Create(DM).Error
	return err
}

// DeleteDevice 删除设备管理记录
// Author [yourname](https://github.com/yourname)
func (DMService *DeviceService)DeleteDevice(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.Device{},"id = ?",ID).Error
	return err
}

// DeleteDeviceByIds 批量删除设备管理记录
// Author [yourname](https://github.com/yourname)
func (DMService *DeviceService)DeleteDeviceByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.Device{},"id in ?",IDs).Error
	return err
}

// UpdateDevice 更新设备管理记录
// Author [yourname](https://github.com/yourname)
func (DMService *DeviceService)UpdateDevice(ctx context.Context, DM system.Device) (err error) {
	err = global.GVA_DB.Model(&system.Device{}).Where("id = ?",DM.ID).Updates(&DM).Error
	return err
}

// GetDevice 根据ID获取设备管理记录
// Author [yourname](https://github.com/yourname)
func (DMService *DeviceService)GetDevice(ctx context.Context, ID string) (DM system.Device, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&DM).Error
	return
}
// GetDeviceInfoList 分页获取设备管理记录
// Author [yourname](https://github.com/yourname)
func (DMService *DeviceService)GetDeviceInfoList(ctx context.Context, info systemReq.DeviceSearch) (list []system.Device, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&system.Device{})
    var DMs []system.Device
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Device_id != nil && *info.Device_id != "" {
        db = db.Where("device_id = ?", *info.Device_id)
    }
    if info.Subscribe_id != nil && *info.Subscribe_id != "" {
        db = db.Where("subscribe_id = ?", *info.Subscribe_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&DMs).Error
	return  DMs, total, err
}
func (DMService *DeviceService)GetDevicePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
