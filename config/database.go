package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

type Database struct {
	DbType   string `mapstructure:"db-type" json:"dbType"`
	DbName   string `mapstructure:"db-name" json:"dbName"`
	Host     string `mapstructure:"host" json:"host"`
	Port     uint   `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"userName"`
	Password string `mapstructure:"password" json:"password"`
}

func init() {
	cdb := Conf.Database

	switch cdb.DbType {
	case "mysql":
		DB = cdb.openMysql()
	default:
		panic("Unrecognized db type")
	}
}

func (d *Database) openMysql() *gorm.DB {
	fmt.Println("connecting...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username, d.Password, d.Host, d.Port, d.DbName)
	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		panic(fmt.Errorf(`failed to connect db:%w`, err))
	}
	return db
}
func (db *Database) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, db.DbName)
	return dsn
}
