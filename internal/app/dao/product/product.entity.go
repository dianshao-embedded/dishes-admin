package product

import (
	"context"

	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/util/structure"
)

// Get Product db model
func GetProductDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Product))
}

// Product
type SchemaProduct schema.Product

// Convert to Product entity
func (a SchemaProduct) ToProduct() *Product {
	item := new(Product)
	structure.Copy(a, item)
	return item
}

// Product entity
type Product struct {
	util.Model
	Name         string `gorm:"size:50;index;"` // 名称
	Description  string `gorm:"size:200;"`      // 描述
	Os           string `gorm:"size:50;"`       // 操作系统
	UpdateMethod string `gorm:"size:50;"`       // 更新方式
	Format       string `gorm:"size:50;"`       // 固件格式
	Network      string `gorm:"size:50;"`       // 网络连接方式
	Type         string `gorm:"size:50;"`       // 终端类型

}

// Convert to Product schema
func (a Product) ToSchemaProduct() *schema.Product {
	item := new(schema.Product)
	structure.Copy(a, item)
	return item
}

// Product entity list
type Products []*Product

// Convert to Product schema list
func (a Products) ToSchemaProducts() []*schema.Product {
	list := make([]*schema.Product, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaProduct()
	}
	return list
}
