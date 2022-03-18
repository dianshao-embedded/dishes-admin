package device

import (
	"context"

	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/util/structure"
)

// Get Device db model
func GetDeviceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Device))
}

// Device
type SchemaDevice schema.Device

// Convert to Device entity
func (a SchemaDevice) ToDevice() *Device {
	item := new(Device)
	structure.Copy(a, item)
	return item
}

// Device entity
type Device struct {
	util.Model
	ProductID uint64 `gorm:"index;not null;"`  // 产品 ID
	Name      string `gorm:"size:100;index;"`  // 设备名称
	Version   string `gorm:"size:50;"`         // 当前版本
	Stage     string `gorm:"size:50;"`         // 设备所属阶段
	SN        string `gorm:"size:100;"`        // 设备SN
	Status    string `gorm:"size:50;"`         // 设备更新状态
	Failure   int    `gorm:"index;default:0;"` // 失败次数

}

// Convert to Device schema
func (a Device) ToSchemaDevice() *schema.Device {
	item := new(schema.Device)
	structure.Copy(a, item)
	return item
}

// Device entity list
type Devices []*Device

// Convert to Device schema list
func (a Devices) ToSchemaDevices() []*schema.Device {
	list := make([]*schema.Device, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDevice()
	}
	return list
}
