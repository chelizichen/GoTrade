package mapper

import (
	component_db "com_sgrid_gotrade/src/components/db"
	"com_sgrid_gotrade/src/object"
	"com_sgrid_gotrade/src/object/dto"
	"com_sgrid_gotrade/src/object/pojo"
	"com_sgrid_gotrade/src/utils"
	"fmt"
)

func QueryStocks() []string {
	var DB = component_db.DB
	var resp []string
	DB.Model(&object.SaveConf{}).Select("q_code").Group("q_code").Find(&resp)
	return resp
}

func QueryTrades() ([]*pojo.TradeMsg, int64, error) {
	var DB = component_db.DB
	var resp []*pojo.TradeMsg
	var total int64
	err := DB.Debug().Model(&object.TradeMsg{}).Count(&total).Find(&resp).Error
	return resp, total, err
}

func QueryConfs(obj *dto.PageBasicReq) ([]*object.SaveConf, int64, error) {
	var DB = component_db.DB
	var resp []*pojo.SaveConf
	fmt.Println("obj", obj)
	var total int64
	err := DB.Debug().Model(&pojo.SaveConf{}).
		Where(`q_code like ?`, "%"+obj.Keyword+"%").
		Count(&total).
		Offset((obj.PageNum - 1) * obj.PageSize).
		Limit(obj.PageSize).
		Find(&resp).
		Error

	var convertedResp []*object.SaveConf
	if err == nil {
		for _, item := range resp {
			// 创建一个新实例并复制字段
			convertedItem := &object.SaveConf{}
			// 自动转换 pojo.SaveConf 到 object.SaveConf
			utils.ConvertFiled(item, convertedItem)
			convertedResp = append(convertedResp, convertedItem)
		}
	}
	return convertedResp, total, err
}

func SaveMsg(obj pojo.TradeMsg) error {
	var DB = component_db.DB
	err := DB.Model(&pojo.TradeMsg{}).Save(obj).Error
	return err
}
