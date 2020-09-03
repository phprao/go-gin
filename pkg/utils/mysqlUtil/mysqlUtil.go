package mysqlUtil

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"votePlatfom/configs"
	"votePlatfom/pkg/utils/loggerUtil"
)

// ** query need rows.Close to release db ins
// ** exec will release automatic

// 连接管理
var connMap map[string]*gorm.DB

// 关闭所有连接
func CloseAll() {
	for _, db := range connMap {
		db.Close()
	}

	loggerUtil.Info("all mysql conn closed...")
}

// 获取连接：多数据库连接
func GetMysqlConn(connectionName ...string) *gorm.DB {
	// 获取连接名
	var name string
	if len(connectionName) == 0 {
		name = configs.DefaultConnection
	} else {
		name = connectionName[0]
	}

	// 查看是否已建立过连接池
	db, ok := connMap[name]
	if ok == true {
		return db
	}

	// 加载数据库配置
	config := configs.GetDbConnectionConfigByName(name)

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
		config.DB_CHARSET,
	)

	// connect and open db connection
	// db.DB() --> *sql.DB
	db, MysqlDbErr := gorm.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		panic("database data source name error: " + MysqlDbErr.Error())
	}

	loggerUtil.Info("connect to mysql success: " + name)

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(config.DB_MAX_OPEN_CONNS)
	db.DB().SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(config.DB_MAX_IDLE_CONNS)
	db.DB().SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(config.DB_MAX_LIFETIME_CONNS)
	db.DB().SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once avoid connect failed
	// else error will be reported until db first sql operate
	if MysqlDbErr = db.DB().Ping(); nil != MysqlDbErr {
		panic("database connect failed: " + MysqlDbErr.Error())
	}

	loggerUtil.Info("mysql ping is ok: " + name)

	// 初始化connMap，否则会报错
	if connMap == nil {
		connMap = make(map[string]*gorm.DB)
	}
	connMap[name] = db

	return db
}
