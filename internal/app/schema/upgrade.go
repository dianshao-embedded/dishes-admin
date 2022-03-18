package schema

import "time"

// 更新记录
type Upgrade struct {
	ID         uint64    `json:"id,string"`
	DeviceID   uint64    `json:"device_id,string"`
	Name       string    `json:"name" binding:"`
	FirmwareID uint64    `json:"firmware_id,string"`
	Size       string    `json:"size"`    // 大小
	Version    string    `json:"version"` // 版本号
	Status     string    `json:"status"`  // 更新结果
	Failure    int       `json:"failure"` // 失败次数
	Stage      string    `json:"stage"`   // 固件阶段 dev or test or release
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Query parameters for db
type UpgradeQueryParam struct {
	PaginationParam
	DeviceID uint64 `form:"device_id"`
	Status   string `form:"status"`
}

// Query options for db (order or select fields)
type UpgradeQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type UpgradeQueryResult struct {
	Data       Upgrades
	PageResult *PaginationResult
}

// 更新记录 Object List
type Upgrades []*Upgrade
