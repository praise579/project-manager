package config

type Minio struct {
	Endpoint  string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey string `mapstructure:"access-Key" json:"accessKey"`
	SecretKey string `mapstructure:"secret-key" json:"secretKey"`
}
