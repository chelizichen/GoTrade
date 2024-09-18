package object

type TradeMsg struct {
	Q_USER_ID   int    `json:"q_user_id,omitempty"`   // 用户
	Q_CODE      string `json:"q_code,omitempty"`      // 个股
	T_TYPE      int    `json:"t_type,omitempty"`      // 交易类型
	T_PRICE     int    `json:"t_price,omitempty"`     // 交易价格
	T_SEND_MSG  string `json:"t_send_msg,omitempty"`  // 交易发的信息
	CREATE_TIME string `json:"create_time,omitempty"` // 创建时间
}

type SaveConf struct {
	ID                            int    `json:"id,omitempty"`
	Q_TYPE                        int    `json:"q_type,omitempty"`                         // 类型
	Q_CODE                        string `json:"q_code,omitempty"`                         // 量化个股
	Q_PERSONAL_ASSETS             string `json:"q_personal_assets,omitempty"`              // 个人资产
	Q_STOCK_ASSETS                string `json:"q_stock_assets,omitempty"`                 // 股票市值
	Q_TYPE_1_PARAMS_RATE          string `json:"q_type_1_params_rate,omitempty" `          // 大单买入或卖出（一分钟内涨幅超多少进行买入）
	Q_TYPE_1_PARAMS_INIT_POSITION string `json:"q_type_1_params_init_position,omitempty" ` // 初始仓位
	Q_STATUS                      int    `json:"q_status,omitempty"`                       // 状态 1 创建 2 开启 11 已停用
	Q_USER_ID                     int    `json:"q_user_id,omitempty"`                      // 用户ID
	CREATE_TIME                   string `json:"create_time,omitempty"`
}
