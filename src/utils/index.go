package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var AbortWithError = func(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusOK, &gin.H{
		"code":    -1,
		"message": err,
		"data":    nil,
	})
}

// Done
var AbortWithSucc = func(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, &gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
	})
}

// List
var AbortWithSuccList = func(c *gin.Context, data interface{}, total int64) {
	c.AbortWithStatusJSON(http.StatusOK, &gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
		"total":   total,
	})
}

func GenerateRandomCallbackName() string {
	rand.Seed(time.Now().UnixNano())
	randomPart := rand.Int63n(10000000000000000)
	timestamp := time.Now().UnixMilli()
	return fmt.Sprintf("jQuery%d_%d", randomPart, timestamp)
}

// 替换 URL 中的参数
func ReplaceTarget(url string, market int, code string) (string, string) {
	params := "[JSONPCALLBACK]"
	params2 := "[TIME]"
	params3 := "[MARKET]"
	params4 := "[CODE]"

	name := GenerateRandomCallbackName()
	name2 := time.Now().UnixMilli()

	newURL := strings.ReplaceAll(url, params, name)
	newURL = strings.ReplaceAll(newURL, params2, strconv.FormatInt(name2, 10))
	newURL = strings.ReplaceAll(newURL, params3, strconv.Itoa(market))
	newURL = strings.ReplaceAll(newURL, params4, code)

	return newURL, name
}
