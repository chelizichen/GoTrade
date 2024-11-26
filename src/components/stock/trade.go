package component_stock

import (
	"com_sgrid_gotrade/src/components/constant"
	"com_sgrid_gotrade/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type TradeDirectionVo struct {
	RC     int64  `json:"rc"`
	Rt     int64  `json:"rt"`
	Svr    int64  `json:"svr"`
	Lt     int64  `json:"lt"`
	Full   int64  `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   data   `json:"data"`
}

type data struct {
	Total int64                `json:"total"`
	Diff  []map[string]float64 `json:"diff"`
}

func GetTrade() (ret *TradeDirectionVo) {
	URL, NAME := utils.ReplaceTarget(constant.TRADE_TOTAL, 0, "0")
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

func GetTradeVal() map[string]float64 {
	tdv := GetTrade()
	df := tdv.Data.Diff
	rsp := make(map[string]float64)
	for _, v := range df {
		for _k, _v := range v {
			if _k == "f64" {
				t, ok := rsp["超大单流入"]
				if !ok {
					rsp["超大单流入"] = 0
				}
				rsp["超大单流入"] = _v + t
			}
			if _k == "f65" {
				t, ok := rsp["超大单流出"]
				if !ok {
					rsp["超大单流出"] = 0
				}
				rsp["超大单流出"] = _v + t
			}
			if _k == "f70" {
				t, ok := rsp["大单流入"]
				if !ok {
					rsp["大单流入"] = 0
				}
				rsp["大单流入"] = _v + t
			}
			if _k == "f71" {
				t, ok := rsp["大单流出"]
				if !ok {
					rsp["大单流出"] = 0
				}
				rsp["大单流出"] = _v + t
			}
			if _k == "f76" {
				t, ok := rsp["中单流入"]
				if !ok {
					rsp["中单流入"] = 0
				}
				rsp["中单流入"] = _v + t
			}
			if _k == "f77" {
				t, ok := rsp["中单流出"]
				if !ok {
					rsp["中单流出"] = 0
				}
				rsp["中单流出"] = _v + t
			}
			if _k == "f82" {
				t, ok := rsp["小单流入"]
				if !ok {
					rsp["小单流入"] = 0
				}
				rsp["小单流入"] = _v + t
			}
			if _k == "f83" {
				t, ok := rsp["小单流出"]
				if !ok {
					rsp["小单流出"] = 0
				}
				rsp["小单流出"] = _v + t
			}
		}
	}
	now := time.Now().Format(time.DateTime)
	fmt.Printf(">> 资金流向 START >> %s \n", now)
	zl := (rsp["超大单流入"] + rsp["大单流入"]) - (rsp["超大单流出"] + rsp["大单流出"])
	zd := (rsp["中单流入"]) - (rsp["中单流出"])
	xd := (rsp["小单流入"]) - (rsp["小单流出"])
	cjl := (rsp["超大单流入"] + rsp["大单流入"] + rsp["中单流入"] + rsp["小单流入"])
	fmt.Printf(">> 成交量      : %f  亿元 \n", cjl/10000/10000)
	fmt.Printf(">> 主力净流入   : %f  亿元 \n", zl/10000/10000)
	fmt.Printf(">> 中单净流入   : %f  亿元 \n", zd/10000/10000)
	fmt.Printf(">> 小单净流入   : %f  亿元 \n", xd/10000/10000)
	fmt.Printf("<< ENDLESS 资金流向  %s \n", now)
	rsp["主力净流入"] = zl
	return rsp
}
