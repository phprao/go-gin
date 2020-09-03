package configs

import "os"

var DefaultRedisServer = "redis"

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       string
}

func GetRedisConfig(redisName ...string) RedisConfig {
	server := map[string]RedisConfig{
		"redis": RedisConfig{
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
			os.Getenv("REDIS_PASSWORD"),
			os.Getenv("REDIS_DB"),
		},
	}

	name := ""
	if len(redisName) == 0 {
		name = DefaultRedisServer
	} else {
		name = redisName[0]
	}

	if name == "" {
		panic("no server was selected")
	}

	config, ok := server[name]
	if ok == false {
		panic("the server name was not set")
	}

	return config
}
