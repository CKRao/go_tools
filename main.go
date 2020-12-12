package main

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
