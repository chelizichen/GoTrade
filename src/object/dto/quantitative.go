package object_dto

type SaveConf struct {
	Q_TYPE                        int    `json:"q_type,omitempty"`                        // 量化类型 见 readme.md
	Q_CODE                        string `json:"q_code,omitempty"`                        // 量化个股
	Q_PERSONAL_ASSETS             string `json:"q_personal_assets,omitempty"`             // 个人资产
	Q_STOCK_ASSETS                string `json:"q_stock_assets,omitempty"`                // 股票市值
	Q_TYPE_1_PARAMS_RATE          string `json:"q_type_1_params_rate,omitempty"`          // 大单买入或卖出（一分钟内涨幅超多少进行买入）
	Q_TYPE_1_PARAMS_INIT_POSITION string `json:"q_type_1_params_init_position,omitempty"` // 初始仓位
	Q_STATUS                      int    `json:"q_status,omitempty"`                      // 状态 1 创建 2 开启 11 已停用
	Q_USER_ID                     int    `json:"q_user_id,omitempty"`                     // 用户ID
	CREATE_TIME                   string `json:"create_time,omitempty"`
}
