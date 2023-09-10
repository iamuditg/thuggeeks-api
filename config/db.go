package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

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

func CreateDbConnection() (*sql.DB, error) {
	viper.SetConfigName("db")
	viper.AddConfigPath("./resources/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	var dbConfig DataBaseConfig
	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		log.Fatalf("Error Unmarshaling config: %v", err)
		return nil, err
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}
	return db, nil
}
