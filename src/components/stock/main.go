package component_stock

import (
	"com_sgrid_gotrade/src/components/constant"
	"com_sgrid_gotrade/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	fmt.Println("URL", URL)
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
