package cmd

import (
	"fmt"
	"github.com/CKRao/go_tools/tool_lib/mergo_lib"
	"github.com/CKRao/go_tools/tool_lib/rx_go_lib"
	"github.com/CKRao/go_tools/tool_lib/wire_lib"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

// 注册 version cmd
func init() {
	rootCmd.AddCommand(versionCmd)
}

// 初始化version cmd
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "go_tools version v1.0.0",
	Long:  `This is go_tools version v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Info("go_tools version v1.0.0")
	},
}

var rootCmd = &cobra.Command{
	Use:   "go_tools",
	Short: "go_tools is clarkrao tool",
	Long:  `go_tools is clarkrao tool`,
	Run: func(cmd *cobra.Command, args []string) {
		mergo_lib.Run() // 测试mergo
		wire_lib.Run()  // 测试wire
		rx_go_lib.Run() // 测试rxgo
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
