
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DeviceSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Device_id  *string `json:"device_id" form:"device_id"` 
      Subscribe_id  *string `json:"subscribe_id" form:"subscribe_id"` 
    request.PageInfo
}
