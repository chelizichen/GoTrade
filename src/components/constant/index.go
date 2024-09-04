package constant

const (
	TARGET_GET_CURRENT_PRICE = "http://qt.gtimg.cn/q=s_%s%s"
	TARGET_KLine_HIS         = "https://push2his.eastmoney.com/api/qt/stock/kline/get?cb=[JSONPCALLBACK]&secid=[MARKET].[CODE]&ut=fa5fd1943c7b386f172d6893dbfba10b&fields1=f1%2Cf2%2Cf3%2Cf4%2Cf5%2Cf6&fields2=f51%2Cf52%2Cf53%2Cf54%2Cf55%2Cf56%2Cf57%2Cf58%2Cf59%2Cf60%2Cf61&klt=101&fqt=1&end=20500101&lmt=120&_=[TIME]"
	TARGET_KLine_TDY         = "https://push2his.eastmoney.com/api/qt/stock/trends2/get?fields1=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13&fields2=f51,f52,f53,f54,f55,f56,f57,f58&ut=fa5fd1943c7b386f172d6893dbfba10b&secid=[MARKET].[CODE]&ndays=1&iscr=0&iscca=0&cb=[JSONPCALLBACK]"
)
