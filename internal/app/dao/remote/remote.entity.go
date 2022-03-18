package remote

import (
	"context"

	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/util/structure"
)

// Get Remote db model
func GetRemoteDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Remote))
}

// Remote
type SchemaRemote schema.Remote

// Convert to Remote entity
func (a SchemaRemote) ToRemote() *Remote {
	item := new(Remote)
	structure.Copy(a, item)
	return item
}

// Remote entity
type Remote struct {
	util.Model
	DeviceID uint64 `gorm:"index;not null;"` // 设备 ID
	Address  string `gorm:"size:50;"`        // IP 地址
	Key      string `gorm:"size:150;"`       // 公钥

}

// Convert to Remote schema
func (a Remote) ToSchemaRemote() *schema.Remote {
	item := new(schema.Remote)
	structure.Copy(a, item)
	return item
}

// Remote entity list
type Remotes []*Remote

// Convert to Remote schema list
func (a Remotes) ToSchemaRemotes() []*schema.Remote {
	list := make([]*schema.Remote, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRemote()
	}
	return list
}
