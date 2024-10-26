package component_stock

import (
	"com_sgrid_gotrade/src/components/constant"
	"com_sgrid_gotrade/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type stockComponent struct{}

var StockComponent = new(stockComponent)

func (s *stockComponent) GetMarket(stockCode string) int {
	if strings.HasPrefix(stockCode, "11") {
		return 1
	} else if strings.HasPrefix(stockCode, "12") {
		return 0
	} else {
		firstChar := string(stockCode[0])
		switch firstChar {
		case "6", "9", "5", "7":
			return 1
		default:
			return 0
		}
	}
}

// GetStockHQ retrieves stock information from a given target URL.
func (s *stockComponent) GetStockHQ(target string) (map[string]string, error) {
	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(body), "~")

	if len(data) < 6 {
		return nil, fmt.Errorf("invalid data format")
	}

	_, name, code, price, change, changePercent := data[0], data[1], data[2], data[3], data[4], data[5]

	result := map[string]string{
		"name":          name,
		"code":          code,
		"price":         price,
		"change":        change,
		"changePercent": changePercent,
	}

	return result, nil
}

type KlineHisVo struct {
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

func (s *stockComponent) GetKlineHis(stockCode string) (ret KlineHisVo) {
	market := StockComponent.GetMarket(stockCode)
	URL, NAME := utils.ReplaceTarget(constant.TARGET_KLine_HIS, market, stockCode)
	fmt.Println("URL", URL)
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error:", err)
		return ret
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ret
	}

	// 将字节切片转换为字符串
	bodyString := string(bodyBytes)
	bodyString = strings.ReplaceAll(bodyString, NAME, "")
	// 移除 JSONP 回调函数部分
	callbackIndex := strings.Index(bodyString, "(")
	if callbackIndex >= 0 {
		bodyString = bodyString[callbackIndex+1 : len(bodyString)-2]
		bodyString = strings.TrimSpace(bodyString)
	}
	json.Unmarshal([]byte(bodyString), &ret)
	return ret
}

type KlineTodayVo struct {
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

func (s *stockComponent) GetKlineToday(stockCode string) (ret KlineTodayVo) {
	market := StockComponent.GetMarket(stockCode)
	URL, NAME := utils.ReplaceTarget(constant.TARGET_KLine_TDY, market, stockCode)
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 将字节切片转换为字符串
	bodyString := string(bodyBytes)
	bodyString = strings.ReplaceAll(bodyString, NAME, "")
	// 移除 JSONP 回调函数部分
	callbackIndex := strings.Index(bodyString, "(")
	if callbackIndex >= 0 {
		bodyString = bodyString[callbackIndex+1 : len(bodyString)-2]
		bodyString = strings.TrimSpace(bodyString)
	}
	json.Unmarshal([]byte(bodyString), &ret)
	return ret
}

func (s *stockComponent) GetDiff(stockCode string) *StockPrice {
	GKT := s.GetKlineToday(stockCode)
	var length = len(GKT.Data.Trends)
	var curr = strings.Split(GKT.Data.Trends[length-1], ",")[1]
	var last = strings.Split(GKT.Data.Trends[length-3], ",")[1]
	var open = strings.Split(GKT.Data.Trends[0], ",")[1]
	currPrice, _ := strconv.ParseFloat(curr, 64)
	lastPrice, _ := strconv.ParseFloat(last, 64)
	openPrice, _ := strconv.ParseFloat(open, 64)
	ret := &StockPrice{
		CurrentPrice: currPrice,
		LastPrice:    lastPrice,
		OpenPrice:    openPrice,
		Name:         GKT.Data.Name,
		Code:         GKT.Data.Code,
	}
	return ret
}

type StockBaseInfo struct {
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
		F260 string  `json:"f260"`
		F261 string  `json:"f261"`
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
		F295 string  `json:"f295"`
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

func (s *stockComponent) GetStockBaseInfo(code string, market int) (ret *StockBaseInfo, err error) {
	URL, NAME := utils.ReplaceTarget(constant.CODE_BASE_INFO, market, code)
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 将字节切片转换为字符串
	bodyString := string(bodyBytes)
	bodyString = strings.ReplaceAll(bodyString, NAME, "")
	// 移除 JSONP 回调函数部分
	callbackIndex := strings.Index(bodyString, "(")
	if callbackIndex >= 0 {
		bodyString = bodyString[callbackIndex+1 : len(bodyString)-2]
		bodyString = strings.TrimSpace(bodyString)
	}
	json.Unmarshal([]byte(bodyString), &ret)
	return ret, nil
}

type StockPrice struct {
	CurrentPrice float64
	LastPrice    float64
	OpenPrice    float64
	Code         string
	Name         string
}

func (s *StockPrice) GetDiff() float64 {
	return s.CurrentPrice - s.LastPrice
}

func (s *StockPrice) GetDiffRate() float64 {
	return s.GetDiff() / s.CurrentPrice * 100
}

func (s *StockPrice) GetRate() float64 {
	return (s.CurrentPrice - s.OpenPrice) / s.CurrentPrice * 100
}

// 打印结构体的所有信息
func (s *StockPrice) Info() string {
	return fmt.Sprintf("股票代码: %s, 股票名称: %s, 当前价格: %.2f, 前两分钟价格: %.2f, 今日开盘价: %.2f, 价格差异: %.2f, 价格差异率: %.2f%%, 开盘至当前涨跌幅: %.2f%%\n",
		s.Code, s.Name, s.CurrentPrice, s.LastPrice, s.OpenPrice, s.GetDiff(), s.GetDiffRate(), s.GetRate())
}
