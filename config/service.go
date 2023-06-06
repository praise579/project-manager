package config

type Service struct {
	Mode string `mapstructure:"mode" json:"mode"`
	Port uint   `mapstructure:"port" json:"port"`
}
