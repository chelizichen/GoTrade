package schedule

import (
	component_cache "com_sgrid_gotrade/src/components/cache"
	component_stock "com_sgrid_gotrade/src/components/stock"
	"com_sgrid_gotrade/src/mapper"
	"com_sgrid_gotrade/src/object/pojo"
	"encoding/json"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

const QuantitativeStocks = "Quantitative|Stocks"

func InitSchedule() {
	loadStocksTask()
	var cronInstance = cron.New()
	quantitativeJob(cronInstance)
	loadStocks(cronInstance)
	cronInstance.Start()
}

func quantitativeJob(cronInstance *cron.Cron) {
	var scheduledTask = func() {
		fmt.Println("quantitativeJob 开始执行任务")
		// Redis 拿到注册的股票
		var Cache = component_cache.Cache
		var Rsp []string
		bytes, err := Cache.Get(QuantitativeStocks).Bytes()
		if err != nil {
			fmt.Println("Cache.Get(QuantitativeStocks).Bytes().error 任务失败:", err)
		}
		err = json.Unmarshal(bytes, &Rsp)
		if err != nil {
			fmt.Println("json.Unmarshal(bytes, &Rsp).error 任务失败:", err)
		}
		fmt.Println("Rsp", Rsp)
		for _, code := range Rsp {
			s := component_stock.StockComponent.GetDiff(code)
			fmt.Println(s.Info())
			// 下跌朝上 买入趋势
			if s.GetRate() > -0.8 && s.GetRate() < 0 && s.GetDiffRate() > 0.3 {
				fmt.Println("下跌朝上 买入趋势")
				err = pushMessage(s, 1)
				if err != nil {
					fmt.Println("pushMessage(s, 1).error 任务失败:", err.Error())
				}
			}
		}
		fmt.Println("quantitativeJob 执行任务结束")
	}
	// 定义每天9:30到15:00之间每两分钟执行一次
	err := cronInstance.AddFunc("0 * 9-15 * * 1-5", scheduledTask) // 从9:00到14:59的每分钟
	if err != nil {
		fmt.Println("任务注册错误:", err)
		return
	}
	err = cronInstance.AddFunc("30-59 9 * * 1-5", scheduledTask) // 从9:30到9:59的每分钟
	if err != nil {
		fmt.Println("任务注册错误:", err)
		return
	}
	err = cronInstance.AddFunc("0-59 15 * * 1-5", scheduledTask) // 从15:00到15:59的每分钟
	if err != nil {
		fmt.Println("任务注册错误:", err)
		return
	}
}

// Redis 拿到注册的股票
var loadStocksTask = func() {
	var Cache = component_cache.Cache
	stocks := mapper.QueryStocks()
	bytes, err := json.Marshal(stocks)
	if err != nil {
		fmt.Println("json.Marshal.error 任务失败:", err)
	}
	err = Cache.Set(QuantitativeStocks, bytes, time.Hour*24).Err()
	if err != nil {
		fmt.Println("loadStocksTask.error 任务失败:", err)
	}
	fmt.Println("stocks ", stocks)
}

func loadStocks(cronInstance *cron.Cron) {
	err := cronInstance.AddFunc("0 15 9 * * MON-FRI", loadStocksTask) // 9:30到9:59
	if err != nil {
		fmt.Println("任务注册错误:", err)
		return
	}
}

func pushMessage(s *component_stock.StockPrice, tradeType int) error {
	var msg string = ""
	if tradeType == 1 {
		msg = "下跌朝上 买入趋势 | "
	}
	return mapper.SaveMsg(pojo.TradeMsg{
		Q_USER_ID:   0,
		Q_CODE:      s.Code,
		T_TYPE:      tradeType,
		T_PRICE:     s.CurrentPrice,
		T_SEND_MSG:  msg + s.Info(),
		CREATE_TIME: time.Now().Format("2006-01-02 15:04:05"),
	})
}
