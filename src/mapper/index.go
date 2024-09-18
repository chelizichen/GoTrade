package mapper

import (
	component_db "com_sgrid_gotrade/src/components/db"
	"com_sgrid_gotrade/src/object"
	"com_sgrid_gotrade/src/object/pojo"
)

func QueryStocks() []string {
	var DB = component_db.DB
	var resp []string
	DB.Model(&object.SaveConf{}).Select("q_code").Group("q_code").Find(&resp)
	return resp
}

func SaveMsg(obj pojo.TradeMsg) error {
	var DB = component_db.DB
	err := DB.Model(&pojo.TradeMsg{}).Save(obj).Error
	return err
}
