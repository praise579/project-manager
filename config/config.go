package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Service *Service `mapstructure:"service" json:"service"`
	DB      *DB      `mapstructure:"db" json:"db"`
	Minio   *Minio   `mapstructure:"minio" json:"minio"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取初始配置文件异常：%s", err))
	}

	viper.Unmarshal(Conf)
}

func PrintConfig() {
	fmt.Printf("%s", Conf)
}
