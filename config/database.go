package config

import (
	"fmt"
	"github.com/praise579/project-manager/model/system"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
	DbType       string `mapstructure:"db-type" json:"dbType"`
	DbName       string `mapstructure:"db-name" json:"dbName"`
	Host         string `mapstructure:"host" json:"host"`
	Port         uint   `mapstructure:"port" json:"port"`
	Username     string `mapstructure:"username" json:"userName"`
	Password     string `mapstructure:"password" json:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	StringSize   uint   `mapstructure:"string-size" json:"stringSize"`
	LogLevel     string `mapstructure:"log-level" json:"logLevel"`
}

func init() {
	cdb := Conf.Database

	switch cdb.DbType {
	case "mysql":
		DB = cdb.openMysql()
		DB.Set("gorm:table_options", "ENGINE=InnoDB")
	default:
		panic("Unrecognized db type")
	}
}

func (d *Database) openMysql() *gorm.DB {

	db, err := gorm.Open(mysql.New(d.mysqlConfig()), &gorm.Config{
		Logger: d.logConfig(),
	})
	if err != nil {
		panic(fmt.Errorf(`failed to connect db:%w`, err))
	}

	// 连接池配置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(d.MaxIdleConns)
	sqlDB.SetMaxOpenConns(d.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func (d *Database) mysqlConfig() (mc mysql.Config) {
	mc = mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			d.Username, d.Password, d.Host, d.Port, d.DbName),
		DefaultStringSize: d.StringSize,
	}
	return
}

func (d *Database) logConfig() (l logger.Interface) {
	var level logger.LogLevel
	switch d.LogLevel {
	case "Silent":
		level = logger.Silent
	case "Error":
		level = logger.Error
	case "Warn":
		level = logger.Warn
	case "Info":
		level = logger.Info
	default:
		level = logger.Info
	}

	l = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  level,       // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	return
}

func doMigrate() {
	DB.AutoMigrate(
		&system.SysUser{},
	)
}
