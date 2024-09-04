package component_interceptor

import (
	"com_sgrid_gotrade/src/utils"

	"github.com/gin-gonic/gin"
)

type componentInterceptor struct{}

var InterceptorComponent = new(componentInterceptor)

func (i *componentInterceptor) StockCodeCheck(c *gin.Context) {
	if len(c.Query("code")) != 6 {
		utils.AbortWithError(c, "query.code.length validate error")
	}
	c.Next()
}
