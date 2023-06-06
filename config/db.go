package config

import "fmt"

type DB struct {
	DbType   string `mapstructure:"db-type" json:"dbType"`
	Host     string `mapstructure:"host" json:"host"`
	Port     uint   `mapstructure:"port" json:"port"`
	Username string `mapstructure:"user-name" json:"userName"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

func (db *DB) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, db.Database)
	return dsn
}
