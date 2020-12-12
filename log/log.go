package log

import (
	"fmt"
	"github.com/CKRao/go_tools/conf"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Logger         *zap.Logger
	ConfigJsonPath string
)

func init() {
	if ConfigJsonPath = conf.ToolsCfg.ConfigJsonPath; len(ConfigJsonPath) == 0 {
		ConfigJsonPath = "./log.yaml"
	}

	cfgBytes, err := ioutil.ReadFile(ConfigJsonPath)
	if err != nil {
		panic(err)
	}

	var cfg zap.Config
	if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		panic(err)
	}

	fmt.Printf("log cfg: %+v \n", cfg)

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Logger = logger
	zap.ReplaceGlobals(logger)
	zap.S().Info("logger construction succeeded")
}

func Sync() {
	Logger.Sync()
}
