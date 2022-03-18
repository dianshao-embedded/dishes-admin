package schema

import "time"

// 固件
type Firmware struct {
	ID        uint64    `json:"id,string"`
	ProductID uint64    `json:"product_id,string"`
	Name      string    `json:"name" binding:"`
	Version   string    `json:"version" binding:"required"` // 版本
	Size      string    `json:"size"`                       // 大小
	Status    string    `json:"status"`                     // 更新状态
	Failure   int       `json:"failure"`                    // 失败次数
	Stage     string    `json:"stage"`                      // 固件阶段 dev or test or release
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type FirmwareQueryParam struct {
	PaginationParam
	ProductID uint64 `form:"productID"`
	Status    string `form:"status"`
}

// Query options for db (order or select fields)
type FirmwareQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type FirmwareQueryResult struct {
	Data       Firmwares
	PageResult *PaginationResult
}

// 固件 Object List
type Firmwares []*Firmware
