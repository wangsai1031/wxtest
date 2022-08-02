package officialaccount

import (
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/draft"
	"github.com/silenceper/wechat/v2/officialaccount/freepublish"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"weixin/log"
)

const (
	AppId  = "wxedb8fc2aadd7bd51"
	Secret = "a3409c9c98a57b6a9bd01637f847b687"
	Token  = "wxtest"
)

func GetOfficialAccount() (officialAccount *officialaccount.OfficialAccount) {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     AppId,
		AppSecret: Secret,
		Token:     Token,
		Cache:     memory,
	}

	officialAccount = wc.GetOfficialAccount(cfg)

	accessToken, _ := officialAccount.GetAccessToken()
	fmt.Println("GetAccessToken", accessToken)
	log.Info.Println("GetAccessToken", accessToken)
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
		log.Error.Println("MediaUpload error", err.Error())
		fmt.Println("MediaUpload", err)
		return
	}

	log.Info.Println("MediaUpload", media)
	return
}

// ImageUpload 永久图片上传
func MediaUploadImg(filename string) (string, error) {

	officialAccount := GetOfficialAccount()
	newMaterial := officialAccount.GetMaterial()

	url, err := newMaterial.ImageUpload(filename)
	if err != nil {
		log.Error.Println("MediaUploadImg error", err.Error())
		fmt.Println("MediaUploadImg", err)
		return "", err
	}

	log.Info.Println("MediaUploadImg", url)
	return url, err
}

// AddMaterial 上传永久性素材（处理视频需要单独上传）
func MediaAddMaterial(mediaType material.MediaType, filename string) (string, string, error) {

	officialAccount := GetOfficialAccount()
	newMaterial := officialAccount.GetMaterial()

	mediaID, url, err := newMaterial.AddMaterial(mediaType, filename)
	if err != nil {
		log.Error.Println("MediaAddMaterial error", err.Error())
		fmt.Println("MediaAddMaterial", err)
		return "", "", err
	}

	log.Info.Println("MediaAddMaterial", mediaID, url)
	return mediaID, url, err
}

// BatchGetMaterial 批量获取永久素材
//reference:https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func BatchGetMaterial(permanentMaterialType material.PermanentMaterialType, offset, count int64) (articleList material.ArticleList, err error) {

	officialAccount := GetOfficialAccount()

	newMaterial := officialAccount.GetMaterial()

	articleList, err = newMaterial.BatchGetMaterial(permanentMaterialType, offset, count)

	if err != nil {
		log.Error.Println("GetMaterialIndex error", err.Error())
		fmt.Println("GetMaterialIndex", err)
		return
	}

	log.Info.Println("GetMaterialIndex", articleList)
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
func AddDraft(articles []*draft.Article) (string, error) {
	officialAccount := GetOfficialAccount()
	newDraft := officialAccount.GetDraft()

	mediaID, err := newDraft.AddDraft(articles)
	if err != nil {
		log.Error.Println("AddDraft error", err.Error())
		fmt.Println("AddDraft", err)
		return "", err
	}

	log.Info.Println("AddDraft", mediaID)
	return mediaID, err
}

func PaginateDraft(offset, count int64, noReturnContent bool) (articleList draft.ArticleList, err error) {
	officialAccount := GetOfficialAccount()
	newDraft := officialAccount.GetDraft()

	articleList, err = newDraft.PaginateDraft(offset, count, noReturnContent)
	if err != nil {
		log.Error.Println("PaginateDraft error", err.Error())
		fmt.Println("PaginateDraft", err)
		return
	}

	log.Info.Println("PaginateDraft", articleList)
	return
}

func Publish(draftId string) (publishID int64, err error) {
	officialAccount := GetOfficialAccount()
	newFreePublish := officialAccount.GetFreePublish()

	publishID, err = newFreePublish.Publish(draftId)
	if err != nil {
		log.Error.Println("Publish error", err.Error())
		fmt.Println("Publish", err)
		return
	}

	log.Info.Println("Publish", publishID)
	return
}

func PublishStatus(publishID int64) (publishStatusList freepublish.PublishStatusList, err error) {
	officialAccount := GetOfficialAccount()
	newFreePublish := officialAccount.GetFreePublish()

	publishStatusList, err = newFreePublish.SelectStatus(publishID)
	if err != nil {
		log.Error.Println("PublishStatus error", err.Error())
		fmt.Println("PublishStatus", err)
		return
	}

	log.Info.Println("PublishStatus", publishStatusList)
	return
}
