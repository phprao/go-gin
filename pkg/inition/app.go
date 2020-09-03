package inition

import (
	"log"

	"github.com/joho/godotenv"
)

/**
初始化操作
*/
func init() {
	loadEnv()
}

/**
加载env环境配置
*/
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
