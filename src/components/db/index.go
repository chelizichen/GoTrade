package component_db

import (
	"com_sgrid_gotrade/src/framework/config"
	"com_sgrid_gotrade/src/object/pojo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func LoadDB(conf *config.SgridConf) {
	db_master := conf.GetString("db")
	db, err := gorm.Open(mysql.Open(db_master), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "trade_",
			SingularTable: true,
		},
	})
	db.AutoMigrate(&pojo.SaveConf{}, &pojo.TradeMsg{})
	if err != nil {
		panic(err)
	}
	DB = (db)
}
