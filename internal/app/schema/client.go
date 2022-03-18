package schema

type Client struct {
	UpgradeID uint64 `json:"upgrade_id,string"`
	Filename  string `json:"file_name"`
	Verison   string `json:"version"`    // 版本
	Update    int    `json:"update"`     // 0-不更新；1-更新
	Event     string `json:"event"`      // 事件
	Stage     string `json:"stage"`      // 固件阶段 dev or test or release
	Size      string `json:"size"`       // 大小
	ServerKey string `json:"server_key"` // 服务端公钥
	ClientKey string `json:"client_key"` // 客户端公钥
}
