package schema

import "time"

// 实例
type Demo struct {
	ID             uint64    `json:"id,string"`
	Code           string    `json:"code" binding:"required"` // 任务编号
	Name           string    `json:"name" binding:"required"` // 任务名称
	Memo           string    `json:"memo"`                    // 任务备注
	Number         int       `json:"number"`                  // 整数
	UnsignedNumber int       `json:"unsigned_number"`         // 无符号整数
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Query parameters for db
type DemoQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type DemoQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type DemoQueryResult struct {
	Data       Demos
	PageResult *PaginationResult
}

// 实例 Object List
type Demos []*Demo
