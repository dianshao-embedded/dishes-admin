package schema

import "time"

// 设备
type Device struct {
	ID          uint64    `json:"id,string"`
	ProductID   uint64    `json:"product_id,string"`
	Name        string    `json:"name" binding:"required"` // 设备名称
	Version     string    `json:"version"`                 // 当前版本
	Stage       string    `json:"stage"`                   // 设备所属阶段
	SN          string    `json:"sn"`                      // 设备SN
	Status      string    `json:"status"`                  // 设备更新状态
	Failure     int       `json:"failure"`                 // 失败次数
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpgradeList Upgrades  `json:"upgrade_list"`
}

// Query parameters for db
type DeviceQueryParam struct {
	PaginationParam
	ProductID uint64 `form:"productID"`
	Name      string `form:"name"`
	SN        string `form:"sn"`
	Stage     string `form:"stage"`
}

// Query options for db (order or select fields)
type DeviceQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type DeviceQueryResult struct {
	Data       Devices
	PageResult *PaginationResult
}

// 设备 Object List
type Devices []*Device
