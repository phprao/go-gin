package configs

import "os"

type DbConnectionConfig struct {
	DB_HOST               string
	DB_PORT               string
	DB_DATABASE           string
	DB_USERNAME           string
	DB_PASSWORD           string
	DB_CHARSET            string
	DB_MAX_OPEN_CONNS     string // 连接池最大连接数
	DB_MAX_IDLE_CONNS     string // 连接池最大空闲数
	DB_MAX_LIFETIME_CONNS string // 连接池连接最长生命周期
}

var DefaultConnection = "opdata"

func GetDbConnectionConfigByName(connectionName ...string) DbConnectionConfig {
	connections := map[string]DbConnectionConfig{
		"opdata": DbConnectionConfig{
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			"utf8mb4",
			"20",
			"2",
			"7200",
		},
		"uc": DbConnectionConfig{
			os.Getenv("DB_HOST_UCENTER"),
			os.Getenv("DB_PORT_UCENTER"),
			os.Getenv("DB_DATABASE_UCENTER"),
			os.Getenv("DB_USERNAME_UCENTER"),
			os.Getenv("DB_PASSWORD_UCENTER"),
			"utf8mb4",
			"20",
			"2",
			"7200",
		},
	}
	connection := ""
	if len(connectionName) == 0 {
		connection = DefaultConnection
	}else{
		connection = connectionName[0]
	}

	if connection == "" {
		panic("no connection was selected")
	}

	config, ok := connections[connection]
	if ok == false {
		panic("the connection name was not set")
	}

	return config
}
