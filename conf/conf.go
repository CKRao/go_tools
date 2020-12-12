package conf

import (
	"github.com/spf13/viper"
	"log"
)

var (
	ToolsCfg *ToolsConfig // 默认配置
)

func init() {
	cfg := viper.New()

	cfg.AddConfigPath("./")
	cfg.AddConfigPath("./conf")
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")

	// 读取配置
	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// 序列化到结构体
	if err := cfg.Unmarshal(&ToolsCfg); err != nil {
		log.Fatal(err)
	}
}

// ToolsConfig 工具配置
type ToolsConfig struct {
	TimeStamp   string `yaml:"TimeStamp"`
	Author      string `yaml:"Author"`
	PassWd      string `yaml:"PassWd"`
	Information struct {
		Name   string   `yaml:"Name"`
		Age    string   `yaml:"Age"`
		Alise  []string `yaml:"Alise"`
		Image  string   `yaml:"Image"`
		Public bool     `yaml:"Public"`
	} `yaml:"Information"`
	Favorite struct {
		Sport       []string `yaml:"Sport"`
		Music       []string `yaml:"Music"`
		LuckyNumber int      `yaml:"LuckyNumber"`
	} `yaml:"Favorite"`
	ConfigJsonPath string `yaml:"ConfigJsonPath"`
}
