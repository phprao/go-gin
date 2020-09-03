package configs

import "os"

var DefaultServer = "http"

type ServerConfig struct {
	Host string
	Port string
}

func GetServerConfig(serverName ...string) ServerConfig {
	server := map[string]ServerConfig{
		"http": ServerConfig{
			os.Getenv("HTTP_HOST"),
			os.Getenv("HTTP_PORT"),
		},
	}

	name := ""
	if len(serverName) == 0 {
		name = DefaultServer
	} else {
		name = serverName[0]
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
