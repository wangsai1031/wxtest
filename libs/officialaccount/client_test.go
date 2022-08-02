package officialaccount

import (
	"github.com/silenceper/wechat/v2/officialaccount/draft"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"testA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAccessToken, err := GetAccessToken()
			if err != nil {
				t.Errorf("GetAccessToken() error = %+v", err)
				return
			}
			t.Logf("GetAccessToken() success = %+v", gotAccessToken)
		})
	}
}

func TestMediaUploadImg(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{"./testmedia/process_daotu.png"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := MediaUploadImg(tt.args.filename)
			if err != nil {
				t.Errorf("MediaUploadImg() error = %+v", err)
				return
			}
			t.Logf("MediaUploadImg() success = %+v", url)
		})
	}
}

func TestMediaMaterial(t *testing.T) {
	type args struct {
		mediaType material.MediaType
		filename  string
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{material.MediaTypeImage, "./testmedia/2022-08-02_103202.png"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			media, err := MediaUpload(tt.args.mediaType, tt.args.filename)
			if err != nil {
				t.Errorf("MediaUpload() error = %+v", err)
				return
			}
			t.Logf("MediaUpload() success = %+v", media)
		})
	}
}

func TestMediaAddMaterial(t *testing.T) {
	type args struct {
		mediaType material.MediaType
		filename  string
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{material.MediaTypeImage, "./testmedia/process_daotu.png"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mediaID, url, err := MediaAddMaterial(tt.args.mediaType, tt.args.filename)
			if err != nil {
				t.Errorf("MediaUploadImg() error = %+v", err)
				return
			}
			t.Logf("MediaUploadImg() success = %+v, %+v", mediaID, url)
		})
	}
}

func TestBatchGetMaterial(t *testing.T) {
	type args struct {
		permanentMaterialType material.PermanentMaterialType
		offset                int64
		count                 int64
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{permanentMaterialType: material.PermanentMaterialTypeImage, count: 20}},
		//{"test2", args{permanentMaterialType: material.PermanentMaterialTypeNews, count: 20}},
		//{"test3", args{permanentMaterialType: material.PermanentMaterialTypeVideo, count: 20}},
		//{"test4", args{permanentMaterialType: material.PermanentMaterialTypeVoice, count: 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := BatchGetMaterial(
				tt.args.permanentMaterialType,
				tt.args.offset,
				tt.args.count,
			)
			if err != nil {
				t.Errorf("BatchGetMaterial() error = %+v", err)
				return
			}
			t.Logf("BatchGetMaterial() success = %+v", url)
		})
	}
}

func TestClearQuota(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ClearQuota()
			if err != nil {
				t.Errorf("ClearQuota() error = %+v", err)
				return
			}
			t.Logf("ClearQuota() success")
		})
	}
}

func TestAddDraft(t *testing.T) {

	content := `<div class="content custom"><p>开发者可新增常用的素材到草稿箱中进行使用。上传到草稿箱中的素材被群发或发布后，该素材将从草稿箱中移除。新增草稿可在公众平台官网 - 草稿箱中查看和管理。</p> <h2 id="接口请求说明"><a href="#接口请求说明" class="header-anchor">#</a> 接口请求说明</h2> <p>http 请求方式：POST（请使用 https 协议）https://api.weixin.qq.com/cgi-bin/draft/add?access_token=ACCESS_TOKEN</p> <p>调用示例</p> <div class="language-json extra-class"><pre class="language-json"><code><span class="token punctuation">{</span>
    <span class="token property">"articles"</span><span class="token operator">:</span> <span class="token punctuation">[</span>
        <span class="token punctuation">{</span>
            <span class="token property">"title"</span><span class="token operator">:</span>TITLE<span class="token punctuation">,</span>
            <span class="token property">"author"</span><span class="token operator">:</span>AUTHOR<span class="token punctuation">,</span>
            <span class="token property">"digest"</span><span class="token operator">:</span>DIGEST<span class="token punctuation">,</span>
            <span class="token property">"content"</span><span class="token operator">:</span>CONTENT<span class="token punctuation">,</span>
            <span class="token property">"content_source_url"</span><span class="token operator">:</span>CONTENT_SOURCE_URL<span class="token punctuation">,</span>
            <span class="token property">"thumb_media_id"</span><span class="token operator">:</span>THUMB_MEDIA_ID<span class="token punctuation">,</span>
            <span class="token property">"need_open_comment"</span><span class="token operator">:</span><span class="token number">0</span><span class="token punctuation">,</span>
            <span class="token property">"only_fans_can_comment"</span><span class="token operator">:</span><span class="token number">0</span>
        <span class="token punctuation">}</span>
        <span class="token comment">//若新增的是多图文素材，则此处应还有几段 articles 结构</span>
    <span class="token punctuation">]</span>
<span class="token punctuation">}</span>
</code></pre></div><p>请求参数说明</p> <div class="table-wrp"><table><thead><tr><th>参数</th> <th>是否必须</th> <th>说明</th></tr></thead> <tbody><tr><td>title</td> <td>是</td> <td>标题</td></tr> <tr><td>author</td> <td>否</td> <td>作者</td></tr> <tr><td>digest</td> <td>否</td> <td>图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。</td></tr> <tr><td>content</td> <td>是</td> <td>图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。</td></tr> <tr><td>content_source_url</td> <td>否</td> <td>图文消息的原文地址，即点击“阅读原文”后的URL</td></tr> <tr><td>thumb_media_id</td> <td>是</td> <td>图文消息的封面图片素材id（必须是永久MediaID）</td></tr> <tr><td>need_open_comment</td> <td>否</td> <td>Uint32 是否打开评论，0不打开(默认)，1打开</td></tr> <tr><td>only_fans_can_comment</td> <td>否</td> <td>Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论</td></tr></tbody></table></div><h2 id="接口返回说明"><a href="#接口返回说明" class="header-anchor">#</a> 接口返回说明</h2> <div class="language-json extra-class"><pre class="language-json"><code><span class="token punctuation">{</span>
   <span class="token property">"media_id"</span><span class="token operator">:</span>MEDIA_ID
<span class="token punctuation">}</span>
</code></pre></div><p>返回参数说明</p> <div class="table-wrp"><table><thead><tr><th>参数</th> <th>描述</th></tr></thead> <tbody><tr><td>media_id</td> <td>上传后的获取标志，长度不固定，但不会超过 128 字符</td></tr></tbody></table></div></div>`

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
	articles := []*draft.Article{
		{
			Title:              "测试title",
			Author:             "测试作者",
			Digest:             "图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空",
			Content:            content,
			ContentSourceURL:   "https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html",
			ThumbMediaID:       "HEIhl0HtYZrgMLz4_jBrzk4TqXJpRKRvWy6yw8ggP7NDAwlhHXqTcPiPgheoaw_b",
			ShowCoverPic:       1,
			NeedOpenComment:    1,
			OnlyFansCanComment: 0,
		},
	}

	type args struct {
		articles []*draft.Article
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{articles}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mediaId, err := AddDraft(tt.args.articles)
			if err != nil {
				t.Errorf("AddDraft() error = %+v", err)
				return
			}
			t.Logf("AddDraft() success = %+v", mediaId)
		})
	}
}

func TestPaginateDraft(t *testing.T) {
	type args struct {
		offset          int64
		count           int64
		noReturnContent bool
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{count: 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList, err := PaginateDraft(
				tt.args.offset,
				tt.args.count,
				tt.args.noReturnContent,
			)
			if err != nil {
				t.Errorf("PaginateDraft() error = %+v", err)
				return
			}
			t.Logf("PaginateDraft() success = %+v", articleList)
		})
	}
}

func TestPublish(t *testing.T) {
	type args struct {
		draftId string
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzsN0OyfofifBeaIk-gKHNt9-Qe4GLJoKKPe6-azsj4y4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList, err := Publish(
				tt.args.draftId,
			)
			if err != nil {
				t.Errorf("Publish() error = %+v", err)
				return
			}
			t.Logf("Publish() success = %+v", articleList)
		})
	}
}

func TestPublishStatus(t *testing.T) {
	type args struct {
		publishID int64
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{2247483653}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList, err := PublishStatus(
				tt.args.publishID,
			)
			if err != nil {
				t.Errorf("PublishStatus() error = %+v", err)
				return
			}
			t.Logf("PublishStatus() success = %+v", articleList)
		})
	}
}
