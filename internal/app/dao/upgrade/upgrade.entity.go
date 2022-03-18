package upgrade

import (
	"context"

	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/util/structure"
)

// Get Upgrade db model
func GetUpgradeDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Upgrade))
}

// Upgrade
type SchemaUpgrade schema.Upgrade

// Convert to Upgrade entity
func (a SchemaUpgrade) ToUpgrade() *Upgrade {
	item := new(Upgrade)
	structure.Copy(a, item)
	return item
}

// Upgrade entity
type Upgrade struct {
	util.Model
	DeviceID   uint64 `gorm:"index;not null;"`  // 设备 ID
	FirmwareID uint64 `gorm:"index;not null;"`  // 固件 ID
	Name       string `gorm:"size:200"`         // 名称
	Version    string `gorm:"size:50;"`         // 版本号
	Size       string `gorm:"size:200;"`        // 大小
	Status     string `gorm:"size:50;"`         // 更新结果
	Failure    int    `gorm:"index;default:0;"` // 失败次数
	Stage      string `gorm:"size:50"`          // 固件阶段 dev or test or release
}

// Convert to Upgrade schema
func (a Upgrade) ToSchemaUpgrade() *schema.Upgrade {
	item := new(schema.Upgrade)
	structure.Copy(a, item)
	return item
}

// Upgrade entity list
type Upgrades []*Upgrade

// Convert to Upgrade schema list
func (a Upgrades) ToSchemaUpgrades() []*schema.Upgrade {
	list := make([]*schema.Upgrade, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUpgrade()
	}
	return list
}
