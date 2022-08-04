package mysql

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/utils"

	handlerLog "weixin/common/handlers/log"
)

const (
	callBackBeforeName = "core:before"
	callBackAfterName  = "core:after"
	startTime          = "_start_time"

	// CSTLayout China Standard Time Layout
	CSTLayout = "2006-01-02 15:04:05"
)

type TraceLogPlugin struct{}

type TraceSQL struct {
	Timestamp   string  `json:"timestamp"`     // 时间格式：2006-01-02 15:04:05
	Stack       string  `json:"stack"`         // 文件地址和行号
	SQL         string  `json:"sql"`           // SQL语句
	Rows        int64   `json:"rows_affected"` // 影响行数
	CostSeconds float64 `json:"cost_seconds"`  // 执行时长(单位秒)
}

func (op *TraceLogPlugin) Name() string {
	return "traceLogPlugin"
}

// Plugin GORM plugin interface
type Plugin interface {
	Name() string
	Initialize(*gorm.DB) error
}

func (op *TraceLogPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &TraceLogPlugin{}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	return
}

func after(db *gorm.DB) {
	_ctx := db.Statement.Context

	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	sqlInfo := &TraceSQL{
		Timestamp:   time.Now().Format(CSTLayout),
		Stack:       utils.FileWithLineNum(),
		SQL:         db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...),
		Rows:        db.Statement.RowsAffected,
		CostSeconds: time.Since(ts).Seconds(),
	}

	var sqlInfoStr []string
	t := reflect.TypeOf(*sqlInfo)
	v := reflect.ValueOf(*sqlInfo)
	for k := 0; k < t.NumField(); k++ {
		sqlInfoStr = append(sqlInfoStr, fmt.Sprintf("%s=%v", t.Field(k).Name, v.Field(k).Interface()))
	}
	// log 记录
	handlerLog.Trace.Info(_ctx, strings.Join(sqlInfoStr, "||"))
}
