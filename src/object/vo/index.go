package vo

// 股票基础信息
type VoStockBaseInfo struct {
	TotalMarketValue      float64 `json:"totalMarketValue,omitempty"`      // 总市值
	TodayChangeValue      float64 `json:"todayChangeValue,omitempty"`      // 今日总成交量
	TodayChangeTotalValue float64 `json:"todayChangeTotalValue,omitempty"` // 今日总成交额
	EarnRatio             float64 `json:"earnRatio,omitempty"`             // 市盈率
	TurnoverRatio         float64 `json:"turnoverRatio,omitempty"`         // 换手率
	TodayPriceChangeRatio float64 `json:"todayPriceChangeRatio,omitempty"` // 今日涨跌幅
}

func NewVoStockNowHq(resp *StockNowHq) *VoStockNowHq {
	rsp := new(VoStockNowHq)
	// build buy
	rsp.Sell5 = float64(resp.Data.F31) / 100
	rsp.Sell4 = float64(resp.Data.F33) / 100
	rsp.Sell3 = float64(resp.Data.F35) / 100
	rsp.Sell2 = float64(resp.Data.F37) / 100
	rsp.Sell1 = float64(resp.Data.F39) / 100
	rsp.Sell5Vol = float64(resp.Data.F32)
	rsp.Sell4Vol = float64(resp.Data.F34)
	rsp.Sell3Vol = float64(resp.Data.F36)
	rsp.Sell2Vol = float64(resp.Data.F38)
	rsp.Sell1Vol = float64(resp.Data.F40)
	// build sell
	rsp.Buy1 = float64(resp.Data.F19) / 100
	rsp.Buy2 = float64(resp.Data.F17) / 100
	rsp.Buy3 = float64(resp.Data.F15) / 100
	rsp.Buy4 = float64(resp.Data.F13) / 100
	rsp.Buy5 = float64(resp.Data.F11) / 100
	rsp.Buy1Vol = float64(resp.Data.F20)
	rsp.Buy2Vol = float64(resp.Data.F18)
	rsp.Buy3Vol = float64(resp.Data.F16)
	rsp.Buy4Vol = float64(resp.Data.F14)
	rsp.Buy5Vol = float64(resp.Data.F12)
	// build other
	rsp.TurnoverRatio = float64(resp.Data.F168) / 100
	return rsp
}

// 东方财富当前行情
type StockNowHq struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int    `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		F43  int     `json:"f43"`
		F44  int     `json:"f44"`
		F45  int     `json:"f45"`
		F46  int     `json:"f46"`
		F47  int     `json:"f47"`
		F48  float64 `json:"f48"`
		F49  int     `json:"f49"`
		F50  int     `json:"f50"`
		F51  int     `json:"f51"`
		F52  int     `json:"f52"`
		F57  string  `json:"f57"`
		F58  string  `json:"f58"`
		F59  int     `json:"f59"`
		F60  int     `json:"f60"`
		F71  int     `json:"f71"`
		F84  int64   `json:"f84"`
		F85  int64   `json:"f85"`
		F86  int     `json:"f86"`
		F92  float64 `json:"f92"`
		F107 int     `json:"f107"`
		F108 float64 `json:"f108"`
		F111 int     `json:"f111"`
		F116 float64 `json:"f116"`
		F117 float64 `json:"f117"`
		F152 int     `json:"f152"`
		F161 int     `json:"f161"`
		F162 int     `json:"f162"`
		F163 int     `json:"f163"`
		F164 int     `json:"f164"`
		F167 int     `json:"f167"`
		F168 int     `json:"f168"`
		F169 int     `json:"f169"`
		F170 int     `json:"f170"`
		F171 int     `json:"f171"`
		F177 int     `json:"f177"`
		F191 int     `json:"f191"`
		F192 int     `json:"f192"`
		F256 string  `json:"f256"`
		F257 int     `json:"f257"`
		F260 int     `json:"f260"`
		F261 int     `json:"f261"`
		F262 string  `json:"f262"`
		F269 string  `json:"f269"`
		F270 int     `json:"f270"`
		F277 int64   `json:"f277"`
		F278 int     `json:"f278"`
		F279 int     `json:"f279"`
		F285 string  `json:"f285"`
		F286 int     `json:"f286"`
		F288 int     `json:"f288"`
		F292 int     `json:"f292"`
		F294 int     `json:"f294"`
		F295 int     `json:"f295"`
		F301 int     `json:"f301"`
		F31  int     `json:"f31"`
		F32  int     `json:"f32"`
		F33  int     `json:"f33"`
		F34  int     `json:"f34"`
		F35  int     `json:"f35"`
		F36  int     `json:"f36"`
		F37  int     `json:"f37"`
		F38  int     `json:"f38"`
		F39  int     `json:"f39"`
		F40  int     `json:"f40"`
		F19  int     `json:"f19"`
		F20  int     `json:"f20"`
		F17  int     `json:"f17"`
		F18  int     `json:"f18"`
		F15  int     `json:"f15"`
		F16  int     `json:"f16"`
		F13  int     `json:"f13"`
		F14  int     `json:"f14"`
		F11  int     `json:"f11"`
		F12  int     `json:"f12"`
		F734 string  `json:"f734"`
		F747 string  `json:"f747"`
		F748 int     `json:"f748"`
	} `json:"data"`
}

type VoStockNowHq struct {
	Buy1          float64 `json:"buy1,omitempty"`          // 买1
	Buy2          float64 `json:"buy2,omitempty"`          // 买2
	Buy3          float64 `json:"buy3,omitempty"`          // 买3
	Buy4          float64 `json:"buy4,omitempty"`          // 买4
	Buy5          float64 `json:"buy5,omitempty"`          // 买5
	Buy1Vol       float64 `json:"buy1Vol,omitempty"`       // 买1单数
	Buy2Vol       float64 `json:"buy2Vol,omitempty"`       // 买2单数
	Buy3Vol       float64 `json:"buy3Vol,omitempty"`       // 买3单数
	Buy4Vol       float64 `json:"buy4Vol,omitempty"`       // 买4单数
	Buy5Vol       float64 `json:"buy5Vol,omitempty"`       // 买5单数
	Sell1         float64 `json:"sell1,omitempty"`         // 卖1
	Sell2         float64 `json:"sell2,omitempty"`         // 卖2
	Sell3         float64 `json:"sell3,omitempty"`         // 卖3
	Sell4         float64 `json:"sell4,omitempty"`         // 卖4
	Sell5         float64 `json:"sell5,omitempty"`         // 卖5
	Sell1Vol      float64 `json:"sell1Vol,omitempty"`      // 卖1单数
	Sell2Vol      float64 `json:"sell2Vol,omitempty"`      // 卖2单数
	Sell3Vol      float64 `json:"sell3Vol,omitempty"`      // 卖3单数
	Sell4Vol      float64 `json:"sell4Vol,omitempty"`      // 卖4单数
	Sell5Vol      float64 `json:"sell5Vol,omitempty"`      // 卖5单数
	TurnoverRatio float64 `json:"turnoverRatio,omitempty"` // 换手率
}

// 东方财富历史K
type VoKlineHis struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int    `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Code      string   `json:"code"`
		Market    int      `json:"market"`
		Name      string   `json:"name"`
		Decimal   int      `json:"decimal"`
		Dktotal   int      `json:"dktotal"`
		PreKPrice float64  `json:"preKPrice"`
		Klines    []string `json:"klines"`
	} `json:"data"`
}

// 东方财富日K
type VoKlineToday struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int    `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Code          string   `json:"code"`
		Market        int      `json:"market"`
		Type          int      `json:"type"`
		Status        int      `json:"status"`
		Name          string   `json:"name"`
		Decimal       int      `json:"decimal"`
		PreSettlement float64  `json:"preSettlement"`
		PreClose      float64  `json:"preClose"`
		Beticks       string   `json:"beticks"`
		TrendsTotal   int      `json:"trendsTotal"`
		Time          int      `json:"time"`
		Kind          int      `json:"kind"`
		PrePrice      float64  `json:"prePrice"`
		Trends        []string `json:"trends"`
	} `json:"data"`
}

// 东方财富板块
type StockBk struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int64  `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Total int `json:"total"`
		Diff  []struct {
			F3   int    `json:"f3"`
			F4   int    `json:"f4"`
			F12  string `json:"f12"`
			F13  int    `json:"f13"`
			F14  string `json:"f14"`
			F128 string `json:"f128"`
			F140 string `json:"f140"`
			F141 int    `json:"f141"`
			F152 int    `json:"f152"`
		} `json:"diff"`
	} `json:"data"`
}

type VoStockBk struct {
	BkName        string  `json:"bkName,omitempty"`
	BkChangeRatio float64 `json:"bkChangeRatio,omitempty"`
	BkTopStock    string  `json:"bkTopStock,omitempty"`
}

func NewVoStockBkArr(data *StockBk) []VoStockBk {
	rsp := make([]VoStockBk, 0)
	for _, v := range data.Data.Diff {
		rsp = append(rsp, VoStockBk{
			BkName:        v.F14,
			BkChangeRatio: float64(v.F3) / 100,
			BkTopStock:    v.F12,
		})
	}
	return rsp
}
