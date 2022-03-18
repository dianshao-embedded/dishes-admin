package firmware

import (
	"context"

	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/util/structure"
)

// Get Firmware db model
func GetFirmwareDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Firmware))
}

// Firmware
type SchemaFirmware schema.Firmware

// Convert to Firmware entity
func (a SchemaFirmware) ToFirmware() *Firmware {
	item := new(Firmware)
	structure.Copy(a, item)
	return item
}

// Firmware entity
type Firmware struct {
	util.Model
	ProductID uint64 `gorm:"index;not null;"` // 产品 ID
	Name      string `gorm:"size:200"`
	Version   string `gorm:"size:50;index;"` // 版本
	Size      string `gorm:"size:200;"`      // 大小
	Status    string `gorm:"size:50"`
	Failure   int    `gorm:"default:0;"` // 失败次数
	Stage     string `gorm:"size:50"`    // 固件阶段 dev or test or release
}

// Convert to Firmware schema
func (a Firmware) ToSchemaFirmware() *schema.Firmware {
	item := new(schema.Firmware)
	structure.Copy(a, item)
	return item
}

// Firmware entity list
type Firmwares []*Firmware

// Convert to Firmware schema list
func (a Firmwares) ToSchemaFirmwares() []*schema.Firmware {
	list := make([]*schema.Firmware, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaFirmware()
	}
	return list
}
