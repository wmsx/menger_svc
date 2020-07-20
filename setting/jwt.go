package setting

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/wmsx/xconf/pkg/client/source"
)

const (
	JWTNamespace = "jwt"
)

var JWTSetting = JWT{}

type JWT struct {
	Secret string
}

func setUpJwt(appName, env string) error {
	sourceUrl := XConfURL
	if env == "dev" {
		sourceUrl = DevXConfURL
	}

	jwtConfig, err := config.NewConfig(
		config.WithSource(source.NewSource(appName, env, JWTNamespace, source.WithURL(sourceUrl))),
	)
	if err != nil {
		log.Error("获取jwt配置失败")
		return err
	}
	err = jwtConfig.Scan(&JWTSetting)
	if err != nil {
		log.Error("获取jwt配置失败")
		return err
	}
	jwtWatcher, err := jwtConfig.Watch()
	if err != nil {
		log.Error("jwt配置watch失败")
		return err
	}

	go func() {
		for {
			// 会比较 value，内容不变不会返回
			v, err := jwtWatcher.Next()
			if err != nil {
				log.Error("jwt配置变更获取失败")
			} else {
				if err := v.Scan(&JWTSetting); err != nil {
					log.Error("watch获取jwt配置失败")
				} else {
					log.Info("watch获取jwt配置结果: ", JWTSetting)
				}
			}
		}
	}()
	return nil
}
