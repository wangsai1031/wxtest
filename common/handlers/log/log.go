package log

import (
	"context"
	"log"

	"team.wphr.vip/technology-group/infrastructure/logger/xdlog"
	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/common/handlers/conf"
)

var (
	Trace *xdlog.XdLogHandle
)

type CommonLog interface {
	Debugf(ctx context.Context, prefix string, format string, args ...interface{})
	Errorf(ctx context.Context, prefix string, format string, args ...interface{})
	Warnf(ctx context.Context, prefix string, format string, args ...interface{})
	Infof(ctx context.Context, prefix string, format string, args ...interface{})
	Fatalf(ctx context.Context, prefix string, format string, args ...interface{})
	RegisterContextFormat(ctxFmt func(ctx context.Context) string)
}

func NewNormalLogger() (*xdlog.XdLogHandle, error) {
	config := xdlog.FileConfig{}
	config.AutoClear = conf.Viper.GetBool("log.auto_clear")
	config.ClearHours = conf.Viper.GetInt32("log.clear_hours")
	config.FilePrefix = conf.Viper.GetString("log.file_prefix")
	config.FileDir = conf.Viper.GetString("log.dir")
	config.Separate = conf.Viper.GetBool("log.separate")
	config.Level = conf.Viper.GetString("log.level")
	config.LogDebug = conf.Viper.GetBool("log.log_debug")

	logger, err := xdlog.NewLoggerWithCfg(&config)
	if err != nil {
		log.Fatalf("init NormalLogger err %v \n", err)
	}
	logger.RegisterContextFormat(trace.FormatCtx) // 注册 context 解析回调
	return logger, nil

}

func Init() {
	var err error

	Trace, err = NewNormalLogger()
	if err != nil {
		log.Fatalf("log init err %v \n", err)
	}
}
