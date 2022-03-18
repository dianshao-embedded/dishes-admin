package demo

import (
	"context"

	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/util/structure"
)

// Get Demo db model
func GetDemoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Demo))
}

// Demo
type SchemaDemo schema.Demo

// Convert to Demo entity
func (a SchemaDemo) ToDemo() *Demo {
	item := new(Demo)
	structure.Copy(a, item)
	return item
}

// Demo entity
type Demo struct {
	util.Model
	Code           string `gorm:"size:50;index;"`   // 任务编号
	Name           string `gorm:"size:50;index;"`   // 任务名称
	Memo           string `gorm:"size:1024;"`       // 任务备注
	Number         int    `gorm:"index;default:0;"` // 整数
	UnsignedNumber int    `gorm:"index;default:0;"` // 无符号整数

}

// Convert to Demo schema
func (a Demo) ToSchemaDemo() *schema.Demo {
	item := new(schema.Demo)
	structure.Copy(a, item)
	return item
}

// Demo entity list
type Demos []*Demo

// Convert to Demo schema list
func (a Demos) ToSchemaDemos() []*schema.Demo {
	list := make([]*schema.Demo, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDemo()
	}
	return list
}
