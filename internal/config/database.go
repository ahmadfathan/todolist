package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	MySQLHost   string
	MySQLPort   string
	MySQLUser   string
	MySQLPass   string
	MySQLDBName string
}

func GetDBConnectionString() string {
	cfg := getDBConfig()

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s",
		cfg.MySQLUser,
		cfg.MySQLPass,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDBName,
	)

	return connectionString
}

func getDBConfig() DBConfig {

	return DBConfig{
		MySQLHost:   os.Getenv("MYSQL_HOST"),
		MySQLPort:   os.Getenv("MYSQL_PORT"),
		MySQLUser:   os.Getenv("MYSQL_USER"),
		MySQLPass:   os.Getenv("MYSQL_PASSWORD"),
		MySQLDBName: os.Getenv("MYSQL_DBNAME"),
	}
}
