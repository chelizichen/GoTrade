package pojo

type SaveConf struct {
	ID                            int
	Q_TYPE                        int
	Q_CODE                        string
	Q_PERSONAL_ASSETS             string
	Q_STOCK_ASSETS                string
	Q_TYPE_1_PARAMS_RATE          string ` gorm:"column:q_type_1_params_rate"`          // 大单买入或卖出（一分钟内涨幅超多少进行买入）
	Q_TYPE_1_PARAMS_INIT_POSITION string ` gorm:"column:q_type_1_params_init_position"` // 初始仓位
	Q_STATUS                      int
	Q_USER_ID                     int
	CREATE_TIME                   string
}

type TradeMsg struct {
	Q_USER_ID   int     `json:"q_user_id,omitempty"`   // 用户
	Q_CODE      string  `json:"q_code,omitempty"`      // 个股
	T_TYPE      int     `json:"t_type,omitempty"`      // 交易类型
	T_PRICE     float64 `json:"t_price,omitempty"`     // 交易价格
	T_SEND_MSG  string  `json:"t_send_msg,omitempty"`  // 交易发的信息
	CREATE_TIME string  `json:"create_time,omitempty"` // 创建时间
}
