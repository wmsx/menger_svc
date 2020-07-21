package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/wmsx/menger_svc/setting"
)

var db *gorm.DB

func SetUp() (err error) {
	db, err = gorm.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.ConnectStr)

	if err != nil {
		log.Error("数据库初始化失败 err: ", err)
		return err
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.AutoMigrate(&Menger{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return nil
}

func CloseDB() {
	db.Close()
}
