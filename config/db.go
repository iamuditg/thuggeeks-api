package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sql.DB

type DataBaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func CreateDbConnection() *sql.DB {
	viper.SetConfigName("")
}
