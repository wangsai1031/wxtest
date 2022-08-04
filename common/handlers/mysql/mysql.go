package mysql

import (
	"context"
	"os"
	"time"

	"gorm.io/gorm"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"

	mysql_v2 "team.wphr.vip/technology-group/infrastructure/mysql-v2"
	"team.wphr.vip/technology-group/infrastructure/trace"
)

var (
	//Client对象是线程安全的
	Client *mysql_v2.MysqlConnection
)

func Init() {
	dsn := conf.Viper.GetString("mysql.dsn")

	db, err := mysql_v2.NewConnection(dsn, &gorm.Config{})
	if err != nil {
		log.Trace.Errorf(context.Background(), trace.DLTagUndefined, "dail mysql v2 err: %v ", err)
		os.Exit(1)
	}
	err = db.SetConnMaxLifetime(time.Duration(conf.Viper.GetInt("mysql.conn_max_lifetime")) * time.Second)
	if err != nil {
		log.Trace.Errorf(context.Background(), trace.DLTagUndefined, "set mysql v2 conn_max_lifetime err: %v ", err)
		os.Exit(1)
	}
	err = db.SetMaxIdleConns(conf.Viper.GetInt("mysql.max_idle_conns"))
	if err != nil {
		log.Trace.Errorf(context.Background(), trace.DLTagUndefined, "set mysql v2 max_idle_conns err: %v ", err)
		os.Exit(1)
	}
	err = db.SetMaxOpenConns(conf.Viper.GetInt("mysql.max_open_conns"))
	if err != nil {
		log.Trace.Errorf(context.Background(), trace.DLTagUndefined, "set mysql v2 max_open_conns err: %v ", err)
		os.Exit(1)
	}
	db.Use(&TraceLogPlugin{})
	Client = db
}
