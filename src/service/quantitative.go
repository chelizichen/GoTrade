package service

import (
	component_stock "com_sgrid_gotrade/src/components/stock"
	object_dto "com_sgrid_gotrade/src/object/dto"
	"com_sgrid_gotrade/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func V1_Quantitative_SaveConf(c *gin.Context) {
	var body *object_dto.SaveConf
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
	price := component_stock.StockComponent.GetDiff(code)
	utils.AbortWithSucc(c, price)
}
