package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"team.wphr.vip/technology-group/infrastructure/trace"
	"weixin/common/handlers/log"
	"weixin/common/util"
	message2 "weixin/libs/officialaccount/message"
	"weixin/model"
)

type MessageDao struct {
	BaseDao
}

var MessageDaoInstance = &MessageDao{
	BaseDao: BaseDao{
		NewEntityFunc: func() model.IModel {
			return &model.MessageEntity{}
		},
	},
}

func (d MessageDao) Save(ctx context.Context, msg *message2.MixMessage) (err error) {

	entity := model.MessageEntity{
		ToUserName:   string(msg.ToUserName),
		FromUserName: string(msg.FromUserName),
		MsgType:      string(msg.MsgType),
		Event:        string(msg.Event),
		CreateTime:   msg.CreateTime,
	}

	resourceByte, err := json.Marshal(msg)

	if err != nil {
		log.Trace.Errorf(ctx, trace.DLTagUndefined, "微信异步通知解析异常 err: %v, msg: %+v", err, msg)
		entity.Resource = "{}"
	} else {
		entity.Resource = string(resourceByte)
	}

	tx := d.getDb(ctx).Debug()
	dbResult := tx.Create(&entity)
	err = dbResult.Error

	if err != nil {
		log.Trace.Errorf(ctx, trace.DLTagUndefined, "微信异步通知保存失败 err: %v, msg: %+v", err, msg)
	}

	return
}

// GetListByConditions 通过条件搜索
func (d MessageDao) GetListByConditions(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB, order []Order, offset int64, limit int64) (modelSlice []*model.MessageEntity, total int64, err error) {
	tx := util.GetTransaction(ctx).Debug()

	if limit == 0 {
		limit = defaultLimit
	}

	tx = tx.Model(model.MessageEntity{})

	for _, scope := range scopes {
		tx.Scopes(scope)
	}

	for _, o := range order {
		if o.OrderType == OrderDesc {
			tx = tx.Order(fmt.Sprintf("%s desc", o.FieldName))
		} else {
			tx = tx.Order(o.FieldName)
		}
	}

	dbResult := tx.Count(&total)

	if dbResult.Error == nil {
		dbResult = tx.Offset(int(offset)).Limit(int(limit)).Find(&modelSlice)
	}

	err = dbResult.Error

	return
}

// GetAllByConditions 通过条件搜索
func (d MessageDao) GetAllByConditions(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB) (ret []*model.MessageEntity, err error) {
	tx := util.GetTransaction(ctx).Debug()

	tx = tx.Model(model.MessageEntity{})

	for _, scope := range scopes {
		tx.Scopes(scope)
	}

	dbResult := tx.Find(&ret)
	err = dbResult.Error

	return
}

// GetPluckByConditions 通过条件搜索
func (d MessageDao) GetPluckByConditions(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB, column string) (ret []interface{}, err error) {
	tx := util.GetTransaction(ctx).Debug()

	tx = tx.Model(model.MessageEntity{})

	for _, scope := range scopes {
		tx.Scopes(scope)
	}

	dbResult := tx.Distinct(column).Pluck(column, &ret)

	err = dbResult.Error
	return
}

// Exists 判断是否存在指定条件的数据
func (d MessageDao) Exists(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB) (bool, error) {
	tx := util.GetTransaction(ctx).Debug()
	tx = tx.Model(model.MessageEntity{})

	for _, scope := range scopes {
		tx.Scopes(scope)
	}

	var exists []int64
	dbResult := tx.Limit(1).Pluck("id", &exists)

	err := dbResult.Error
	if err != nil {
		return false, err
	}

	return len(exists) == 1, nil
}
