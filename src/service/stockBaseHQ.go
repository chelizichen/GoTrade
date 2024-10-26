package service

import (
	"com_sgrid_gotrade/src/components/constant"
	component_stock "com_sgrid_gotrade/src/components/stock"
	"com_sgrid_gotrade/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func V1_StockBaseHQ_GET_CURRENT_PRICE(c *gin.Context) {
	code := c.Query("code")
	marketCode := component_stock.StockComponent.GetMarket(code)
	market := ""
	if marketCode == 0 {
		market = "sz"
	} else {
		market = "sh"
	}
	target := fmt.Sprintf(constant.TARGET_GET_CURRENT_PRICE, market, code)
	resp, err := component_stock.StockComponent.GetStockHQ(target)
	if err != nil {
		utils.AbortWithError(c, err.Error())
		return
	}
	utils.AbortWithSucc(c, resp)
}

func V1_StockBaseHQ_GET_KLINE_HIS(c *gin.Context) {
	code := c.Query("code")
	resp := component_stock.StockComponent.GetKlineHis(code)
	utils.AbortWithSucc(c, resp.Data)
}

func V1_StockBaseHQ_GET_KLINE_TODAY(c *gin.Context) {
	code := c.Query("code")
	resp := component_stock.StockComponent.GetKlineToday(code)
	utils.AbortWithSucc(c, resp.Data)
}

func V1_StockBaseHQ_GET_TRADE_VAR(c *gin.Context) {
	rsp := component_stock.GetTradeVal()
	utils.AbortWithSucc(c, rsp)
}

type stockBaseInfo struct {
	TotalMarketValue      float64 `json:"totalMarketValue,omitempty"`      // 总市值
	TodayChangeValue      float64 `json:"todayChangeValue,omitempty"`      // 今日总成交量
	TodayChangeTotalValue float64 `json:"todayChangeTotalValue,omitempty"` // 今日总成交额
	EarnRatio             float64 `json:"earnRatio,omitempty"`             // 市盈率
	TurnoverRatio         float64 `json:"turnoverRatio,omitempty"`         // 换手率
	TodayPriceChangeRatio float64 `json:"todayPriceChangeRatio,omitempty"` // 今日涨跌幅
}

func V1_StockBaseHQ_GET_STOCK_BASE_INFO(c *gin.Context) {
	code := c.Query("code")
	marketCode := component_stock.StockComponent.GetMarket(code)
	resp, err := component_stock.StockComponent.GetStockBaseInfo(code, marketCode)
	if err != nil {
		utils.AbortWithError(c, err.Error())
		return
	}
	rsp := new(stockBaseInfo)
	rsp.TotalMarketValue = resp.Data.F116 / 10000 / 10000     // 总市值
	rsp.TodayChangeValue = float64(resp.Data.F47) / 100       // 今日总成交量
	rsp.TodayChangeTotalValue = resp.Data.F48 / 10000 / 10000 // 今日总成交额
	rsp.TodayPriceChangeRatio = float64(resp.Data.F170) / 100 // 今日涨跌幅
	rsp.EarnRatio = float64(resp.Data.F162) / 100             // 市盈率
	rsp.TurnoverRatio = float64(resp.Data.F168) / 100         // 换手率
	fmt.Println("rsp:", rsp)
	utils.AbortWithSucc(c, rsp)
}
