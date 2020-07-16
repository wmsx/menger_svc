package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/menger_svc/setting"
)

var db *gorm.DB

func SetUp() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.ConnectStr)

	if err != nil {
		log.Fatalf("数据库初始化失败 err: %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.AutoMigrate(&Menger{})

	db.Model(&Menger{}).AddUniqueIndex("udx_email_delete", "email", "delete_at")
	db.Model(&Menger{}).AddUniqueIndex("udx_name_delete", "name", "delete_at")

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	db.Close()
}
