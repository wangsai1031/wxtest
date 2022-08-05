package message

import (
	"encoding/xml"
	"github.com/silenceper/wechat/v2/officialaccount/message"

	"github.com/silenceper/wechat/v2/officialaccount/device"
	"github.com/silenceper/wechat/v2/officialaccount/freepublish"
)

// MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	CommonToken

	// 基本消息
	MsgID         int64   `xml:"MsgId"` // 其他消息推送过来是MsgId
	TemplateMsgID int64   `xml:"MsgID"` // 模板消息推送成功的消息是MsgID
	Content       string  `xml:"Content"`
	Recognition   string  `xml:"Recognition"`
	PicURL        string  `xml:"PicUrl"`
	MediaID       string  `xml:"MediaId"`
	Format        string  `xml:"Format"`
	ThumbMediaID  string  `xml:"ThumbMediaId"`
	LocationX     float64 `xml:"Location_X"`
	LocationY     float64 `xml:"Location_Y"`
	Scale         float64 `xml:"Scale"`
	Label         string  `xml:"Label"`
	Title         string  `xml:"Title"`
	Description   string  `xml:"Description"`
	URL           string  `xml:"Url"`

	// 事件相关
	Event       message.EventType `xml:"Event" json:"Event"`
	EventKey    string            `xml:"EventKey"`
	Ticket      string            `xml:"Ticket"`
	Latitude    string            `xml:"Latitude"`
	Longitude   string            `xml:"Longitude"`
	Precision   string            `xml:"Precision"`
	MenuID      string            `xml:"MenuId"`
	Status      string            `xml:"Status"`
	SessionFrom string            `xml:"SessionFrom"`
	TotalCount  int64             `xml:"TotalCount"`
	FilterCount int64             `xml:"FilterCount"`
	SentCount   int64             `xml:"SentCount"`
	ErrorCount  int64             `xml:"ErrorCount"`

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}

	subscribeMsgPopupEventList []SubscribeMsgPopupEvent `json:"-"`

	SubscribeMsgPopupEvent []struct {
		List SubscribeMsgPopupEvent `xml:"List"`
	} `xml:"SubscribeMsgPopupEvent"`

	// 事件相关：发布能力
	PublishEventInfo struct {
		PublishID     int64                     `xml:"publish_id"`     // 发布任务id
		PublishStatus freepublish.PublishStatus `xml:"publish_status"` // 发布状态
		ArticleID     string                    `xml:"article_id"`     // 当发布状态为0时（即成功）时，返回图文的 article_id，可用于“客服消息”场景
		ArticleDetail struct {
			Count uint `xml:"count"` // 文章数量
			Item  []struct {
				Index      uint   `xml:"idx"`         // 文章对应的编号
				ArticleURL string `xml:"article_url"` // 图文的永久链接
			} `xml:"item"`
		} `xml:"article_detail"` // 当发布状态为0时（即成功）时，返回内容
		FailIndex []uint `xml:"fail_idx"` // 当发布状态为2或4时，返回不通过的文章编号，第一篇为 1；其他发布状态则为空
	} `xml:"PublishEventInfo"`

	// 事件相关：图文消息推送
	ArticleUrlResult struct {
		Count      uint `xml:"Count"` // 文章数量
		ResultList struct {
			Item []struct {
				ArticleIdx uint   `xml:"ArticleIdx"` // 文章对应的编号
				ArticleUrl string `xml:"ArticleUrl"` // 图文的永久链接
			} `xml:"item"`
		} `xml:"ResultList"`
	} `xml:"ArticleUrlResult"`

	CopyrightCheckResult struct {
		Count      uint `xml:"Count"`      // 数量
		CheckState uint `xml:"CheckState"` // 整体校验结果：1-未被判为转载，可以群发，2-被判为转载，可以群发，3-被判为转载，不能群发
	} `xml:"CopyrightCheckResult"` // 原创校验

	// 第三方平台相关
	InfoType                     message.InfoType `xml:"InfoType"`
	AppID                        string           `xml:"AppId"`
	ComponentVerifyTicket        string           `xml:"ComponentVerifyTicket"`
	AuthorizerAppid              string           `xml:"AuthorizerAppid"`
	AuthorizationCode            string           `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64            `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string           `xml:"PreAuthCode"`
	AuthCode                     string           `xml:"auth_code"`
	Info                         struct {
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int    `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
	} `xml:"info"`

	// 卡券相关
	CardID              string `xml:"CardId"`
	RefuseReason        string `xml:"RefuseReason"`
	IsGiveByFriend      int32  `xml:"IsGiveByFriend"`
	FriendUserName      string `xml:"FriendUserName"`
	UserCardCode        string `xml:"UserCardCode"`
	OldUserCardCode     string `xml:"OldUserCardCode"`
	OuterStr            string `xml:"OuterStr"`
	IsRestoreMemberCard int32  `xml:"IsRestoreMemberCard"`
	UnionID             string `xml:"UnionId"`

	// 内容审核相关
	IsRisky       bool   `xml:"isrisky"`
	ExtraInfoJSON string `xml:"extra_info_json"`
	TraceID       string `xml:"trace_id"`
	StatusCode    int    `xml:"status_code"`

	// 设备相关
	device.MsgDevice
}

// SubscribeMsgPopupEvent 订阅通知事件推送的消息体
type SubscribeMsgPopupEvent struct {
	TemplateID            string `xml:"TemplateId" json:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString" json:"SubscribeStatusString"`
	PopupScene            int    `xml:"PopupScene" json:"PopupScene,string"`
}

// SetSubscribeMsgPopupEvents 设置订阅消息事件
func (s *MixMessage) SetSubscribeMsgPopupEvents(list []SubscribeMsgPopupEvent) {
	s.subscribeMsgPopupEventList = list
}

// GetSubscribeMsgPopupEvents 获取订阅消息事件数据
func (s *MixMessage) GetSubscribeMsgPopupEvents() []SubscribeMsgPopupEvent {
	if s.subscribeMsgPopupEventList != nil {
		return s.subscribeMsgPopupEventList
	}
	list := make([]SubscribeMsgPopupEvent, len(s.SubscribeMsgPopupEvent))
	for i, item := range s.SubscribeMsgPopupEvent {
		list[i] = item.List
	}
	return list
}

// EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

// EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}

// ResponseEncryptedXMLMsg 需要返回的消息体
type ResponseEncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	EncryptedMsg string   `xml:"Encrypt"      json:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64    `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string   `xml:"Nonce"        json:"Nonce"`
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name        `xml:"xml"`
	ToUserName   CDATA           `xml:"ToUserName" json:"ToUserName"`
	FromUserName CDATA           `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64           `xml:"CreateTime" json:"CreateTime"`
	MsgType      message.MsgType `xml:"MsgType" json:"MsgType"`
}

// SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName CDATA) {
	msg.ToUserName = toUserName
}

// SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName CDATA) {
	msg.FromUserName = fromUserName
}

// SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

// SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType message.MsgType) {
	msg.MsgType = msgType
}

// GetOpenID get the FromUserName value
func (msg *CommonToken) GetOpenID() string {
	return string(msg.FromUserName)
}
