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
