package main

//clarkrao go 相关的库的练习项目

import (
	"github.com/CKRao/go_tools/cmd"
	"github.com/CKRao/go_tools/conf"
	"github.com/CKRao/go_tools/log"
	"go.uber.org/zap"
)

func main() {
	defer log.Sync()
	zap.S().Infof("load config: %+v", conf.ToolsCfg)
	cmd.Execute()
}
