package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

//go:embed config.yaml
var Config embed.FS

var (
	// 私有变量，外部包无法直接访问
	config *viper.Viper
)

func init() {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("os.Executable() = " + err.Error())
	}
	config = viper.New()
	// 读工程目录配置
	if strings.Contains(executable, "___") {
		config.SetConfigFile("./infra/config/config.yaml")
		err := config.ReadInConfig()
		if err != nil {
			fmt.Println("读工程目录配置错误:", err)
		}
		return
	}
	// 读外部配置
	config.SetConfigFile("./config.yaml")
	err = config.ReadInConfig()
	if err == nil {
		return
	}
	fmt.Println("找不到外部配置目录，改为读嵌入配置文件", err)

	// 读嵌入配置
	data, e := Config.ReadFile("config.yaml")
	if e != nil {
		fmt.Println("读嵌入配置错误:", err)
	}

	config.SetConfigType("yaml")
	err = config.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("读嵌入配置错误:", err)
	}
}

func Get() *viper.Viper {
	return config
}
