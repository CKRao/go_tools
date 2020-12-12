package mergo_lib

import (
	"github.com/imdario/mergo"
	"go.uber.org/zap"
)

// 合并结构体字段的库mergo
// go get github.com/imdario/mergo

type redisConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Dbs     int    `json:"dbs"`
}

var defaultConfig = redisConfig{
	Address: "127.0.0.1",
	Port:    6379,
	Dbs:     1,
}

func Run() {
	var config redisConfig
	config.Dbs = 2

	// 合并两个相同类型的结构或map
	if err := mergo.Merge(&config, defaultConfig); err != nil {
		zap.S().Errorf("mergo merge fail, err: %s", err.Error())
		return
	}

	zap.S().Info("redis address: ", config.Address)
	zap.S().Info("redis port: ", config.Port)
	zap.S().Info("redis db: ", config.Dbs)

	// 在结构和map之间赋值
	cfgMap := make(map[string]interface{})
	if err := mergo.Map(&cfgMap, defaultConfig); err != nil {
		zap.S().Errorf("mergo map fail, err: %s", err.Error())
		return
	}

	zap.S().Infof("merge map: %+v", cfgMap)
}
