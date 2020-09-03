package redisUtil

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"votePlatfom/configs"
)

/**
DEMO

_, err := redisUtil.Conn().Do("set", "gin-name", "1111111")
if err != nil {
	fmt.Println(err)
}
res, err := redisUtil.Conn().Do("get", "gin-name")
if err != nil {
	fmt.Println(err)
}
fmt.Printf("%s\n", res)
*/

var poolMap map[string]*redis.Pool

// 使用特定连接
func Conn(names ...string) redis.Conn {
	conn := getRedisConnFromPool(names...)
	return conn
}

// 获取连接
func getRedisConnFromPool(names ...string) redis.Conn {
	var name string
	if len(names) == 0 {
		name = configs.DefaultRedisServer
	} else {
		name = names[0]
	}
	config := configs.GetRedisConfig(name)

	// 查看是否已建立过连接池
	pool, ok := poolMap[name]
	if ok == false {
		pool = redisPool(fmt.Sprintf("%s:%s", config.Host, config.Port), config.Password)
	}

	// 使用连接池
	conn := pool.Get()

	// 设置数据库
	_, err := conn.Do("select", config.Db)
	if err != nil {
		panic("redis connect failed")
	}

	// 初始化poolMap，否则会报错
	if poolMap == nil {
		poolMap = make(map[string]*redis.Pool)
	}
	poolMap[name] = pool

	return conn
}

func redisPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func redisConn(name ...string) redis.Conn {
	config := configs.GetRedisConfig(name...)
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port), redis.DialPassword(config.Password))

	if err != nil {
		panic("redis connect failed")
	}

	return conn
}
