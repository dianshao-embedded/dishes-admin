package schema

import "time"

// 远程连接
type Remote struct {
	ID        uint64    `json:"id,string"`
	DeviceID  uint64    `json:"device_id,string"`
	Address   string    `json:"address"` // IP 地址
	Key       string    `json:"key"`     // 公钥
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query parameters for db
type RemoteQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type RemoteQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type RemoteQueryResult struct {
	Data       Remotes
	PageResult *PaginationResult
}

// 远程连接 Object List
type Remotes []*Remote
