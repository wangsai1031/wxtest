package util

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"weixin/common/handlers/mysql"
)

type ctxKey struct {
	name string
}

var transactionKey = ctxKey{"db_trans"}

var ErrNoTransaction error = errors.New("no transaction in context")

func BeginTransaction(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, transactionKey, mysql.Client.DB.Begin())
	return ctx
}

func Commit(ctx context.Context) error {
	db, ok := ctx.Value(transactionKey).(*gorm.DB)
	if !ok {
		return ErrNoTransaction
	}

	db = db.Commit()
	return db.Error
}

func Rollback(ctx context.Context) error {
	db, ok := ctx.Value(transactionKey).(*gorm.DB)
	if !ok {
		return ErrNoTransaction
	}

	db = db.Rollback()
	return db.Error
}

func GetTransaction(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(transactionKey).(*gorm.DB)
	if !ok {
		return mysql.Client.DB
	}

	return db
}
