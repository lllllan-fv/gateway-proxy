package models

import (
	"fmt"

	"github.com/lllllan-fv/gateway-proxy/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	cfg := conf.GetConfig()

	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Addr,
		cfg.MySQL.DBName,
	)

	// Shanghai time zone
	dsn = dsn + "&loc=Asia%2FShanghai"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
}
