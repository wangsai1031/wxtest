package dao

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm"
	"weixin/common/consts"
	"weixin/common/util"
	"weixin/model"
)

var (
	errEntityIsNil        error = errors.New("entity is nil")
	errIllegalIdForDelete error = errors.New("illegal id for delete")
	errIllegalIdForUpdate error = errors.New("illegal id for update")
	defaultOffset         int64 = 0
	defaultLimit          int64 = 200
)

var ErrNoResultFound error = errors.New("no matched result")

type OrderType string

const (
	OrderDesc OrderType = "desc"
	OrderAsc  OrderType = "asc"
)

type Order struct {
	FieldName string
	OrderType OrderType
}

const (
	createTimeField       = "CreateTime"
	lastmodifiedTimeField = "ModifyTime"
	idField               = "Id"
	rowStatusField        = "RowStatus"
)

type BaseDao struct {
	NewEntityFunc func() model.IModel
}

/*
 * Create 新增数据
 * entity 参数必须是指向 models/model 中结构体的指针
 */
func (b BaseDao) Create(ctx context.Context, entity model.IModel) error {
	if util.IsNil(entity) {
		return errEntityIsNil
	}

	ts := time.Now().Unix()
	addIntField(createTimeField, ts, entity)
	addIntField(lastmodifiedTimeField, ts, entity)

	tx := b.getDb(ctx)

	dbResult := tx.Create(entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

/*
 * DeleteByID 根据id删除数据
 * 如果表中有row_status字段则实现为软删除，否则是真删除
 */
func (b BaseDao) DeleteByID(ctx context.Context, id int64) error {
	if id == 0 {
		return errIllegalIdForDelete
	}

	entity := b.NewEntityFunc()
	addIntField(idField, id, entity)

	tx := b.getDb(ctx)

	if fieldExists(entity, rowStatusField) {
		addIntField(lastmodifiedTimeField, time.Now().Unix(), entity)
		addIntField(rowStatusField, consts.RowStatusDelete, entity)

		dbResult := tx.Model(entity).Updates(entity)
		if dbResult.Error != nil {
			return dbResult.Error
		}
	} else {
		dbResult := tx.Delete(entity)
		if dbResult.Error != nil {
			return dbResult.Error
		}
	}

	return nil
}

/*
 * GetByID 根据id查询数据
 * entity 参数必须是指向 models/model 中结构体的指针
 */
func (b BaseDao) GetByID(ctx context.Context, id int64, entity model.IModel) error {
	tx := b.getDb(ctx)

	dbResult := tx.First(entity, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

/*
 * UpdateByID 修改数据
 * entity 参数必须是指向 models/model 中结构体的指针
 */
func (b BaseDao) UpdateByID(ctx context.Context, id int64, entity model.IModel) error {
	if id == 0 {
		return errIllegalIdForUpdate
	}

	if util.IsNil(entity) {
		return errEntityIsNil
	}

	addIntField(lastmodifiedTimeField, time.Now().Unix(), entity)

	addIntField(idField, id, entity)

	tx := b.getDb(ctx)

	dbResult := tx.Model(entity).Updates(entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

/*
 * GetByCondition 查询数据
 * modelSlice 参数必须是 models/model 中结构体的切片的 指针
 */
func (b BaseDao) GetByCondition(ctx context.Context, condition model.IModel, notCondition model.IModel,
	fuzzyMatching map[string]string, order []Order, offset *int64, limit *int64, modelSlice interface{}) (int64, error) {

	if util.IsNil(condition) {
		condition = b.NewEntityFunc()
	}

	if util.IsNil(notCondition) {
		notCondition = b.NewEntityFunc()
	}

	tx := b.getDb(ctx).Debug()

	if offset == nil {
		offset = &defaultOffset
	}
	if limit == nil {
		limit = &defaultLimit
	}

	var count int64 = 0

	tx = tx.Where(condition).Not(notCondition)
	for fieldName, fuzzyValue := range fuzzyMatching {
		tx = tx.Where(fmt.Sprintf("%s like ?", fieldName), fmt.Sprintf("%%%s%%", fuzzyValue))
	}

	for _, o := range order {
		if o.OrderType == OrderDesc {
			tx = tx.Order(fmt.Sprintf("%s desc", o.FieldName))
		} else {
			tx = tx.Order(o.FieldName)
		}
	}

	dbResult := tx.Model(condition).Count(&count)
	if dbResult.Error == nil {
		dbResult = tx.Offset(int(*offset)).Limit(int(*limit)).Find(modelSlice)
	}

	err := dbResult.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

// 获取一个事务 或 DB 对象
func (b BaseDao) getDb(ctx context.Context) *gorm.DB {
	return util.GetTransaction(ctx)
}

// 设置整型字段
func addIntField(field string, val int64, entity model.IModel) {

	v := reflect.ValueOf(entity).Elem()
	f := v.FieldByName(field)
	if !f.IsValid() {
		return
	}

	f.SetInt(val)
}

// 判断字段是否存在
func fieldExists(entity model.IModel, fieldName string) bool {
	v := reflect.ValueOf(entity).Elem()
	f := v.FieldByName(idField)
	if !f.IsValid() {
		return false
	}

	return true
}
