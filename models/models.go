package models

import (
	"database/sql"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/wmsx/menger_svc/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func SetUp() (err error) {
	db, err = gorm.Open(mysql.Open(setting.DatabaseSetting.ConnectStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
	})

	if err != nil {
		log.Error("数据库初始化失败 err: ", err)
		return err
	}

	db.AutoMigrate(&Menger{}, &FavoritePost{}, &ThumbUpPost{})

	if sqlDB, err = db.DB(); err != nil {
		log.Error("获取数据库实例失败 err:", err)
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return nil
}

func CloseDB() {
	sqlDB.Close()
}
