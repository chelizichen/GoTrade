package service

import (
	component_stock "com_sgrid_gotrade/src/components/stock"
	"com_sgrid_gotrade/src/mapper"
	"com_sgrid_gotrade/src/object"
	"com_sgrid_gotrade/src/object/dto"
	"com_sgrid_gotrade/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func V1_Quantitative_SaveConf(c *gin.Context) {
	var body *object.SaveConf
	err := c.BindJSON(&body)
	if err != nil {
		utils.AbortWithError(c, err.Error())
		return
	}
	fmt.Println("body", body)
	utils.AbortWithSucc(c, nil)

}

func V1_Quantitative_StopConf(c *gin.Context) {

}

func V1_Quantitative_StartConf(c *gin.Context) {

}

func V1_Quantitative_GetCodeDiffPrice(c *gin.Context) {
	code := c.Query("code")
	s := component_stock.StockComponent.GetDiff(code)
	utils.AbortWithSucc(c, s)
}

func V1_Quantitative_GetStockList(c *gin.Context) {
	stocks := mapper.QueryStocks()
	utils.AbortWithSucc(c, stocks)
}

func V1_Quantitative_QueryTrades(c *gin.Context) {
	list, total, err := mapper.QueryTrades()
	if err != nil {
		utils.AbortWithError(c, err.Error())
		return
	}
	utils.AbortWithSuccList(c, list, total)
}

func V1_Quantitative_QueryConfs(c *gin.Context) {
	var obj *dto.PageBasicReq
	if err := c.ShouldBindJSON(&obj); err != nil {
		utils.AbortWithError(c, err.Error())
		return
	}
	list, total, err := mapper.QueryConfs(obj)
	if err != nil {
		utils.AbortWithError(c, err.Error())
		return
	}
	utils.AbortWithSuccList(c, list, total)
}
