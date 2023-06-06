package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Service  *Service  `mapstructure:"service" json:"service"`
	Database *Database `mapstructure:"database" json:"database"`
	Minio    *Minio    `mapstructure:"minio" json:"minio"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取初始配置文件异常：%s", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("配置文件参数解析异常：%s", err))
	}
}
