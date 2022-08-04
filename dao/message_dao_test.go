package dao

import (
	"context"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
	"weixin/common/handlers/mysql"
	"weixin/model"
)

func init() {
	conf.InitConf("../conf/app.dev.toml")

	log.Init()
	mysql.Init()
}

func TestMessageDao_Exists(t *testing.T) {
	type fields struct {
		BaseDao BaseDao
	}
	type args struct {
		ctx    context.Context
		scopes []func(db *gorm.DB) *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MessageDao{
				BaseDao: tt.fields.BaseDao,
			}
			got, err := d.Exists(tt.args.ctx, tt.args.scopes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageDao_GetAllByConditions(t *testing.T) {
	type fields struct {
		BaseDao BaseDao
	}
	type args struct {
		ctx    context.Context
		scopes []func(db *gorm.DB) *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet []*model.MessageEntity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MessageDao{
				BaseDao: tt.fields.BaseDao,
			}
			gotRet, err := d.GetAllByConditions(tt.args.ctx, tt.args.scopes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByConditions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("GetAllByConditions() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestMessageDao_GetListByConditions(t *testing.T) {
	type fields struct {
		BaseDao BaseDao
	}
	type args struct {
		ctx    context.Context
		scopes []func(db *gorm.DB) *gorm.DB
		order  []Order
		offset int64
		limit  int64
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantModelSlice []*model.MessageEntity
		wantTotal      int64
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MessageDao{
				BaseDao: tt.fields.BaseDao,
			}
			gotModelSlice, gotTotal, err := d.GetListByConditions(tt.args.ctx, tt.args.scopes, tt.args.order, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListByConditions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotModelSlice, tt.wantModelSlice) {
				t.Errorf("GetListByConditions() gotModelSlice = %v, want %v", gotModelSlice, tt.wantModelSlice)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("GetListByConditions() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func TestMessageDao_GetPluckByConditions(t *testing.T) {
	type fields struct {
		BaseDao BaseDao
	}
	type args struct {
		ctx    context.Context
		scopes []func(db *gorm.DB) *gorm.DB
		column string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MessageDao{
				BaseDao: tt.fields.BaseDao,
			}
			gotRet, err := d.GetPluckByConditions(tt.args.ctx, tt.args.scopes, tt.args.column)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPluckByConditions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("GetPluckByConditions() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestMessageDao_Save(t *testing.T) {
	type args struct {
		msg *message.MixMessage
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{msg: &message.MixMessage{
			CommonToken: message.CommonToken{
				ToUserName:   "gh_fdefd501c79b",
				FromUserName: "o7fbK55MPSGavlS-7E2fRC7olytc",
				CreateTime:   1659512227,
				MsgType:      "event",
			},

			Event:       "MASSSENDJOBFINISH",
			MsgID:       1000000001,
			Status:      "send success",
			TotalCount:  2,
			FilterCount: 2,
			SentCount:   2,
			ErrorCount:  0,
			// 缺少该参数
			//CopyrightCheckResult: CopyrightCheckResult{
			//	Count: "0",
			//	CheckState: "1"
			//},

		}}},
		{"test2", args{msg: &message.MixMessage{
			CommonToken: message.CommonToken{
				ToUserName:   "gh_fdefd501c79b",
				FromUserName: "o7fbK58Nq4POkwq_kMGzAheED9P4",
				CreateTime:   1659421865,
				MsgType:      "text",
			},
			Content: "你好呀",
			MsgID:   23757208460886931,
		}}},
		{"test3", args{msg: &message.MixMessage{
			CommonToken: message.CommonToken{
				ToUserName:   "gh_fdefd501c79b",
				FromUserName: "o7fbK58Nq4POkwq_kMGzAheED9P4",
				CreateTime:   1659421883,
				MsgType:      "image",
			},
			PicURL:  "http://mmbiz.qpic.cn/mmbiz_jpg/P5eXRa2DySVUbnWO6DbJRoMibRrFgIvffNvk8h6CCRVTY1ia3EiczjIlhoKnZbjdf7a2UK7MicDUSjgDGSVvmf8ib4A/0",
			MsgID:   23757208157957746,
			MediaID: "Wl_TTIXqibH_46V_m8sJbkvwL9aUx0Uxo0NllxbuQVN06ShH_QZ5_VBO83g3YqSR",
		}}},
		{"test4", args{msg: &message.MixMessage{
			CommonToken: message.CommonToken{
				ToUserName:   "gh_fdefd501c79b",
				FromUserName: "o7fbK59QFqPxKppu19ftpd0mAIXw",
				CreateTime:   1659511270,
				MsgType:      "event",
			},
			Event: "subscribe",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MessageDao{
				BaseDao: BaseDao{
					NewEntityFunc: func() model.IModel {
						return &model.MessageEntity{}
					},
				},
			}
			if err := d.Save(context.Background(), tt.args.msg); err != nil {
				t.Errorf("Save() error = %v", err)
			}
		})
	}
}
