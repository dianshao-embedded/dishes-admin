package schema

import "time"

// 产品
type Product struct {
	ID           uint64    `json:"id,string"`
	Name         string    `json:"name" binding:"required"` // 名称
	Description  string    `json:"description"`             // 描述
	Os           string    `json:"os"`                      // 操作系统
	UpdateMethod string    `json:"update_method"`           // 更新方式
	Format       string    `json:"format"`                  // 固件格式
	Network      string    `json:"network"`                 // 网络连接方式
	Type         string    `json:"type"`                    // 终端类型
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FirmwareList Firmwares `json:"firmware_list"`
}

// Query parameters for db
type ProductQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type ProductQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type ProductQueryResult struct {
	Data       Products
	PageResult *PaginationResult
}

// 产品 Object List
type Products []*Product
