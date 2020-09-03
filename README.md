# go-gin
基于gin框架封装的Http Server  
[gin项目地址 ](https://github.com/gin-gonic/gin) 

本地环境
```bash
go env 
	set GO111MODULE=on
	set GOPATH=D:\dev\php\magook\trunk\server\golang\path
	set GOPROXY=https://goproxy.io
	set GOROOT=D:\Go

go version 
	go version go1.14.7 windows/amd64
``` 

### 主要功能
- .env 解析，采用 [godotenv](https://github.com/joho/godotenv)
- ini 配置文件 - todo
  - 采用 [goconfig](https://github.com/Unknwon/goconfig)
- 路由，路有组
- 路由中间件
- 路由和控制器绑定
- 日志按日期分割，采用 [logrus](https://github.com/sirupsen/logrus)
- mysql支持多库连接，采用 [mysql](https://github.com/go-sql-driver/mysql) 
- ORM 采用 [gorm](https://github.com/jinzhu/gorm)
- redis，支持多库连接 ，采用 [redigo](https://github.com/gomodule/redigo) 
- JWT支持，采用 [jwt-go](https://github.com/dgrijalva/jwt-go)
- 在开发模式下修改代码自动重启 [rizla](https://github.com/kataras/rizla)
  - rizla main.go
- 后台运行：nohup ./main > /dev/null 2>&1 &
- 线上部署，优雅的重启和零停机部署
  - 采用 [endless](https://github.com/fvbock/endless)
  - 启动|停止|状态|热重启 脚本：./bin/man help
