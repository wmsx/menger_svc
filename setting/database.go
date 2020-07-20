package setting

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2/config"
	"github.com/wmsx/xconf/pkg/client/source"
)

const DBNamespace = "database"

var DatabaseSetting = DataBase{}

type DataBase struct {
	TablePrefix string `json:"table_prefix"`
	Type        string `json:"type"`
	ConnectStr  string `json:"connect_str"`
}

func setUpDB(appName, env string) error {
	sourceUrl := XConfURL
	if env == "dev" {
		sourceUrl = DevXConfURL
	}

	dbConfig, err := config.NewConfig(
		config.WithSource(source.NewSource(appName, env, DBNamespace, source.WithURL(sourceUrl))),
	)
	if err != nil {
		log.Error("获取db配置失败")
		return err
	}
	err = dbConfig.Scan(&DatabaseSetting)
	if err != nil {
		log.Error("获取db配置失败")
		return err
	}
	dbWatcher, err := dbConfig.Watch()
	if err != nil {
		log.Error("db配置watch失败")
		return err
	}

	go func() {
		for {
			// 会比较 value，内容不变不会返回
			v, err := dbWatcher.Next()
			if err != nil {
				log.Error("db配置变更获取失败")
			} else {
				if err := v.Scan(&DatabaseSetting); err != nil {
					log.Error("watch获取db配置失败")
				} else {
					log.Info("watch获取db配置结果: ", DatabaseSetting)
				}
			}
		}
	}()
	return nil
}
