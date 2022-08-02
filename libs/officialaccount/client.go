package officialaccount

import (
	"context"
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/draft"
	"github.com/silenceper/wechat/v2/officialaccount/freepublish"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"strings"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
)

//InitWechat 获取wechat实例
//在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat() *wechat.Wechat {
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        conf.Viper.GetString("redis.host"),
		Password:    conf.Viper.GetString("redis.password"),
		Database:    conf.Viper.GetInt("redis.database"),
		MaxActive:   conf.Viper.GetInt("redis.max_active"),
		MaxIdle:     conf.Viper.GetInt("redis.max_idle"),
		IdleTimeout: conf.Viper.GetInt("redis.idle_timeout"),
	}
	redisCache := cache.NewRedis(context.Background(), redisOpts)
	wc.SetCache(redisCache)
	return wc
}

func GetOfficialAccount() (officialAccount *officialaccount.OfficialAccount) {
	wc := InitWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache

	cfg := &config.Config{
		AppID:     conf.Viper.GetString("wxOfficialAccount.app_id"),
		AppSecret: conf.Viper.GetString("wxOfficialAccount.app_secret"),
		Token:     conf.Viper.GetString("wxOfficialAccount.token"),
	}

	officialAccount = wc.GetOfficialAccount(cfg)

	accessToken, _ := officialAccount.GetAccessToken()
	log.Trace.Info("GetAccessToken", accessToken)
	return
}

func GetAccessToken() (accessToken string, err error) {
	officialAccount := GetOfficialAccount()
	return officialAccount.GetAccessToken()
}

// MediaUpload 临时素材上传
func MediaUpload(mediaType material.MediaType, filename string) (media material.Media, err error) {

	officialAccount := GetOfficialAccount()
	newMaterial := officialAccount.GetMaterial()

	media, err = newMaterial.MediaUpload(mediaType, filename)
	if err != nil {
		log.Trace.Error("MediaUpload error", err.Error())
		fmt.Println("MediaUpload", err)
		return
	}

	log.Trace.Info("MediaUpload", media)
	return
}

// ImageUpload 永久图片上传
func MediaUploadImg(filename string) (string, error) {

	officialAccount := GetOfficialAccount()
	newMaterial := officialAccount.GetMaterial()

	url, err := newMaterial.ImageUpload(filename)
	if err != nil {
		log.Trace.Error("MediaUploadImg error", err.Error())
		return "", err
	}

	log.Trace.Info("MediaUploadImg", url)
	return url, err
}

// AddMaterial 上传永久性素材（处理视频需要单独上传）
func MediaAddMaterial(mediaType material.MediaType, filename string) (string, string, error) {

	officialAccount := GetOfficialAccount()
	newMaterial := officialAccount.GetMaterial()

	mediaID, url, err := newMaterial.AddMaterial(mediaType, filename)
	if err != nil {
		log.Trace.Error("MediaAddMaterial error", err.Error())
		return "", "", err
	}

	log.Trace.Info("MediaAddMaterial", mediaID, url)
	return mediaID, url, err
}

// BatchGetMaterial 批量获取永久素材
//reference:https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func BatchGetMaterial(permanentMaterialType material.PermanentMaterialType, offset, count int64) (articleList material.ArticleList, err error) {

	officialAccount := GetOfficialAccount()

	newMaterial := officialAccount.GetMaterial()

	articleList, err = newMaterial.BatchGetMaterial(permanentMaterialType, offset, count)

	if err != nil {
		log.Trace.Error("GetMaterialIndex error", err.Error())
		return
	}

	log.Trace.Info("GetMaterialIndex", articleList)
	return
}

// ClearQuota 清理接口调用次数
// 每个帐号每月共10次清零操作机会，清零生效一次即用掉一次机会；
// https://developers.weixin.qq.com/doc/offiaccount/openApi/clear_quota.html
func ClearQuota() error {
	officialAccount := GetOfficialAccount()
	b := basic.NewBasic(officialAccount.GetContext())
	return b.ClearQuota()
}

/**
Title              string `json:"title"`                 // 标题
Author             string `json:"author"`                // 作者
Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。
Content            string `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且去除JS
ContentSourceURL   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
ThumbMediaID       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久MediaID）
ShowCoverPic       uint   `json:"show_cover_pic"`        // 是否显示封面，0为false，即不显示，1为true，即显示(默认)
NeedOpenComment    uint   `json:"need_open_comment"`     // 是否打开评论，0不打开(默认)，1打开
OnlyFansCanComment uint   `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
*/
// AddDraft 新建草稿
func AddDraft(articles []*draft.Article) (string, error) {
	officialAccount := GetOfficialAccount()
	newDraft := officialAccount.GetDraft()

	mediaID, err := newDraft.AddDraft(articles)
	if err != nil {
		log.Trace.Error("AddDraft error", err.Error())
		return "", err
	}

	log.Trace.Info("AddDraft", mediaID)
	return mediaID, err
}

// PaginateDraft 获取草稿列表
func PaginateDraft(offset, count int64, noReturnContent bool) (articleList draft.ArticleList, err error) {
	officialAccount := GetOfficialAccount()
	newDraft := officialAccount.GetDraft()

	articleList, err = newDraft.PaginateDraft(offset, count, noReturnContent)
	if err != nil {
		log.Trace.Error("PaginateDraft error", err.Error())
		fmt.Println("PaginateDraft", err)
		return
	}

	log.Trace.Info("PaginateDraft", articleList)
	return
}

// Publish 发布接口。需要先将图文素材以草稿的形式保存（见“草稿箱/新建草稿”，
// 如需从已保存的草稿中选择，见“草稿箱/获取草稿列表”），选择要发布的草稿 media_id 进行发布
func Publish(draftId string) (publishID int64, err error) {
	officialAccount := GetOfficialAccount()
	newFreePublish := officialAccount.GetFreePublish()

	publishID, err = newFreePublish.Publish(draftId)
	if err != nil {
		log.Trace.Error("Publish error", err.Error())
		fmt.Println("Publish", err)
		return
	}

	log.Trace.Info("Publish", publishID)
	return
}

// PublishStatus 获取文章发布状态
func PublishStatus(publishID int64) (publishStatus freepublish.PublishStatusList, err error) {
	officialAccount := GetOfficialAccount()
	newFreePublish := officialAccount.GetFreePublish()

	publishStatus, err = newFreePublish.SelectStatus(publishID)
	if err != nil {
		log.Trace.Error("PublishStatus error", err.Error())
		fmt.Println("PublishStatus", err)
		return
	}

	log.Trace.Info("PublishStatus", publishStatus)
	return
}

// Paginate 获取成功发布列表
func PaginatePublish(offset, count int64, noReturnContent bool) (publishList freepublish.ArticleList, err error) {
	officialAccount := GetOfficialAccount()
	newFreePublish := officialAccount.GetFreePublish()

	publishList, err = newFreePublish.Paginate(offset, count, noReturnContent)
	if err != nil {
		log.Trace.Error("PaginatePublish error", err.Error())
		fmt.Println("PaginatePublish", err)
		return
	}

	log.Trace.Info("PaginatePublish", publishList)
	return
}

type ContentImageFile struct {
	FilePath    string // 文件路径
	Placeholder string // 图片在文章内容中的占位符（该占位符将会替换为微信图片链接）
}

type Article struct {
	draftArticle      draft.Article
	ContentImageFiles []*ContentImageFile
	CoverImageFile    string
}

// 综合-发布文章接口，流程示例，非最终方案
func PublishArticle(articles []*Article) {

	// 可以同时发布多篇文章
	var draftArticles []*draft.Article

	for _, article := range articles {
		draftArticle := article.draftArticle

		contentText := draftArticle.Content

		// 文章中的图片文件
		contentImageFiles := article.ContentImageFiles

		for _, imgFile := range contentImageFiles {
			// 1. 上传文章中的图片，获取图片url
			imgUrl, err := MediaUploadImg(imgFile.FilePath)
			if err != nil {
				log.Trace.Error("MediaUploadImg", err)
				return
			}
			// 2. 将文章内容中的图片替换为微信的图片链接
			contentText = strings.Replace(contentText, imgFile.Placeholder, imgUrl, -1)
		}

		draftArticle.Content = contentText

		if article.CoverImageFile != "" && draftArticle.ShowCoverPic == 1 {
			// 3. 上传文章封面图片，获取media_id
			coverMediaID, _, err := MediaAddMaterial(material.MediaTypeImage, article.CoverImageFile)
			if err != nil {
				log.Trace.Error("MediaAddMaterial() error = ", err)
				return
			}

			draftArticle.ThumbMediaID = coverMediaID
		}

		draftArticles = append(draftArticles, &draftArticle)
	}

	draftId, err := AddDraft(draftArticles)
	if err != nil {
		log.Trace.Error("AddDraft() error = ", err)
		return
	}

	// 5. 发布文章
	publishID, err := Publish(draftId)
	if err != nil {
		log.Trace.Error("Publish() error = ", err)
		return
	}

	// 6. 轮询监控发布状态（异步执行）
	TriggerPublishStatusCheckEvent(publishID)

	return
}
