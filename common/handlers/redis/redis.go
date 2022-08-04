package redis

import (
	"context"
	"os"

	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"

	"team.wphr.vip/technology-group/infrastructure/redis"
	"team.wphr.vip/technology-group/infrastructure/trace"
)

var (
	RedisClient *redis.RedisPool
)

func Init() {
	var err error

	rConf := redis.RedisConfig{
		Host:            conf.Viper.GetString("redis.host"),
		Port:            conf.Viper.GetString("redis.port"),
		Password:        conf.Viper.GetString("redis.password"),
		Database:        conf.Viper.GetInt("redis.database"),
		PrefixKey:       conf.Viper.GetString("redis.prefix_key"),
		MaxActive:       conf.Viper.GetInt("redis.max_active"),
		MaxIdle:         conf.Viper.GetInt("redis.max_idle"),
		IdleTimeout:     conf.Viper.GetInt("redis.idle_timeout"),
		MaxConnLifetime: conf.Viper.GetInt("redis.max_conn_lifetime"),
		ConnTimeoutMs:   conf.Viper.GetInt("redis.conn_timeout"),
		ReadTimeoutMs:   conf.Viper.GetInt("redis.read_timeout"),
		WriteTimeoutMs:  conf.Viper.GetInt("redis.write_timeout"),
	}

	pool, err := redis.NewRedisPool(&rConf)
	if err != nil {
		log.Trace.Errorf(context.Background(), trace.DLTagUndefined, "dial redis err %v", err)
		os.Exit(1)
	}

	RedisClient = pool
}
