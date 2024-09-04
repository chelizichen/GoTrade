package router

import (
	component_interceptor "com_sgrid_gotrade/src/components/interceptor"
	"com_sgrid_gotrade/src/service"

	"github.com/gin-gonic/gin"
)

func LoadRouter(engine *gin.Engine) {
	var validate = component_interceptor.InterceptorComponent
	// 行情
	engine.GET("/v1/stockBaseHQ/get_current_price", validate.StockCodeCheck, service.V1_StockBaseHQ_GET_CURRENT_PRICE)
	engine.GET("/v1/stockBaseHQ/get_kline_his", validate.StockCodeCheck, service.V1_StockBaseHQ_GET_KLINE_HIS)
	engine.GET("/v1/stockBaseHQ/get_kline_today", service.V1_StockBaseHQ_GET_KLINE_TODAY)

	// 交易
	engine.GET("/v1/stockTrade/trade_sj", service.V1_StockTrade_TRADE_SJ) // 市价
	engine.GET("/v1/stockTrade/trade_xj", service.V1_StockTrade_TRADE_XJ) // 限价

	// 调仓
	engine.GET("/v1/tradeHistory/get", service.V1_StockTrade_TRADE_XJ) // 限价

	// 自选股
	engine.POST("/v1/selectStocks/save")
	engine.GET("/v1/selectStocks/get")
	engine.GET("/v1/selectStocks/delete")
}
