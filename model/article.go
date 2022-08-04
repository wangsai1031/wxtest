package model

type ArticleEntity struct {
	Id                 int64  `gorm:"primary_key;column:id" json:"id"`
	Title              string `gorm:"column:title" json:"title"`                                 // 标题
	Author             string `gorm:"column:author" json:"author"`                               // 作者
	ThumbMediaId       string `gorm:"column:thumb_media_id" json:"thumb_media_id"`               // 封面缩略图的media_id
	ContentSourceUrl   string `gorm:"column:content_source_url" json:"content_source_url"`       // 阅读原文链接
	Content            string `gorm:"column:content" json:"content"`                             // 内容，支持 HTML 标签。具备微信支付权限的公众号，可以使用 a 标签，其他公众号不能使用
	Digest             string `gorm:"column:digest" json:"digest"`                               // 描述，如本字段为空，则默认抓取正文前64个字
	ShowCoverPic       int64  `gorm:"column:show_cover_pic" json:"show_cover_pic"`               // 是否显示封面，1为显示，0为不显示
	NeedOpenComment    int64  `gorm:"column:need_open_comment" json:"need_open_comment"`         // 是否打开评论，0不打开，1打开
	OnlyFansCanComment int64  `gorm:"column:only_fans_can_comment" json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
	MediaId            string `gorm:"column:media_id" json:"media_id"`                           // 图文上传ID，长度不固定，但不会超过 128 字符
	Index              int64  `gorm:"column:index" json:"index"`
	Url                string `gorm:"column:url" json:"url"`                 // 微信文章链接
	BizStatus          int64  `gorm:"column:biz_status" json:"biz_status"`   // 状态：0默认，1上传封面，2上传文章图片，3上传图文消息素材，4发布中，5发布完成
	UploadTime         int64  `gorm:"column:upload_time" json:"upload_time"` // 上传时间
	SendTime           int64  `gorm:"column:send_time" json:"send_time"`     // 群发消息时间
	CreateTime         int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	ModifyTime         int64  `gorm:"column:modify_time" json:"modify_time"` // 修改时间
}

func (ArticleEntity) TableName() string {
	return "article"
}
