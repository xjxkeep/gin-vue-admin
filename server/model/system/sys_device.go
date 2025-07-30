
// 自动生成模板Device
package system
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 设备管理 结构体  Device
type Device struct {
    global.GVA_MODEL
  Device_id  *string `json:"device_id" form:"device_id" gorm:"uniqueIndex;comment:设备id;column:device_id;" binding:"required"`  //设备id
  Subscribe_id  *string `json:"subscribe_id" form:"subscribe_id" gorm:"uniqueIndex;comment:设备通信id;column:subscribe_id;" binding:"required"`  //设备通信id
  Token  *string `json:"token" form:"token" gorm:"comment:设备token;column:token;" binding:"required"`  //设备token
  Register_ip  *string `json:"register_ip" form:"register_ip" gorm:"column:register_ip;"`  //注册地ip
  Version  *string `json:"version" form:"version" gorm:"comment:软件版本号;column:version;"`  //软件版本号
  Mac  *string `json:"mac" form:"mac" gorm:"comment:硬件序列号;column:mac;"`  //硬件序列号
}


// TableName 设备管理 Device自定义表名 devices
func (Device) TableName() string {
    return "devices"
}





