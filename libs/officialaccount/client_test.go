package officialaccount

import (
	"github.com/silenceper/wechat/v2/officialaccount/draft"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"testing"
	"time"
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
	"weixin/common/util"
)

func init() {
	conf.InitConf("../../conf/app.dev.toml")

	log.Init()
}

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

func TestGetQuota(t *testing.T) {
	type args struct {
		cgiPath string
	}

	tests := []struct {
		name string
		args args
	}{
		{"获取AccessToken", args{"/cgi-bin/token"}},
		{"上传图文消息内的图片获取URL", args{"/cgi-bin/media/uploadimg"}},
		{"新增其他类型永久素材", args{"/cgi-bin/material/add_material"}},
		{"获取素材总数", args{"/cgi-bin/material/get_materialcount"}},
		{"获取素材列表", args{"/cgi-bin/material/batchget_material"}},
		{"新建草稿", args{"/cgi-bin/draft/add"}},
		{"获取草稿列表", args{"/cgi-bin/draft/batchget"}},
		{"发布文章", args{"/cgi-bin/freepublish/submit"}},
		{"发布状态轮询", args{"/cgi-bin/freepublish/get"}},
		{"获取成功发布列表", args{"/cgi-bin/freepublish/batchget"}},
		{"清空 api 的调用quota", args{"/cgi-bin/clear_quota"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quota, err := GetQuota(tt.args.cgiPath)
			if err != nil {
				t.Errorf("GetQuota() error = %+v", err)
				return
			}
			t.Logf("GetQuota() success = %+v", quota)
		})
	}
}

func TestAddDraft(t *testing.T) {

	content1 := `
<div class="content custom">
	<p>开发者可新增常用的素材到草稿箱中进行使用。上传到草稿箱中的素材被群发或发布后，该素材将从草稿箱中移除。新增草稿可在公众平台官网 - 草稿箱中查看和管理。</p> 
	<h2 id="接口请求说明"> 接口请求说明</h2> 
	<p>调用示例</p> 
	<div class="language-json extra-class">
		<pre class="language-json">
			<code>
				<span class="token punctuation">{</span>
				<span class="token property">"articles"</span>
				<span class="token comment">//若新增的是多图文素材，则此处应还有几段 articles 结构</span>
				<span class="token punctuation">]</span>
				<span class="token punctuation">}</span>
			</code>
		</pre>
	</div>
	<img title="上传图片" src="http://mmbiz.qpic.cn/mmbiz_png/gkxfsteFrSmo8gUibC7o9217LicrklVxa5XG3txHJcB3OKDc5k5SVVuHicZ8aOW7C3GgdytUicPEBMQINRxFCCPuyQ/0">
	<p>请求参数说明</p>
	<div class="table-wrp">
		<table>
			<thead>
			<tr><th>参数</th> <th>是否必须</th> <th>说明</th></tr>
			</thead> 
			<tbody>
			<tr><td>title</td> <td>是</td> <td>标题</td></tr> <tr><td>author</td> <td>否</td> <td>作者</td></tr> <tr><td>digest</td> <td>否</td> <td>图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。</td></tr> 
			<tr><td>content</td> <td>是</td> <td>图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。</td></tr> 
			<tr><td>content_source_url</td> <td>否</td> <td>图文消息的原文地址，即点击“阅读原文”后的URL</td></tr> 
			<tr><td>thumb_media_id</td> <td>是</td> <td>图文消息的封面图片素材id（必须是永久MediaID）</td></tr> 
			<tr><td>need_open_comment</td> <td>否</td> <td>Uint32 是否打开评论，0不打开(默认)，1打开</td></tr> 
			<tr><td>only_fans_can_comment</td> <td>否</td> <td>Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论</td></tr>
			</tbody>
		</table>
	</div>
	<h2 id="接口返回说明"> 接口返回说明</h2> 
	<p>返回参数说明</p> 
	<div class="table-wrp">
		<table>
		<thead><tr><th>参数</th> <th>描述</th></tr></thead> 
		<tbody><tr><td>media_id</td> <td>上传后的获取标志，长度不固定，但不会超过 128 字符</td></tr>
		</tbody>
		</table>
	</div>
</div>
`

	content2 := `
<div class="content custom">
<img title="永久素材" src="https://mmbiz.qpic.cn/mmbiz_png/gkxfsteFrSny3z3gbia08SMibBZeQRfEicVhmkdKbMzAylHiaEa1kxdiaQs8cSRPRkd0eRLCFtsCnGdwDC4j9NzmuCQ/0?wx_fmt=png">
</div>
`

	//content = strings.Replace(content, "\"", "\\\\\"", -1)

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
			Title:              "1测试title",
			Author:             "1测试作者",
			Digest:             "1图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空",
			Content:            content1,
			ContentSourceURL:   "https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html",
			ThumbMediaID:       "HEIhl0HtYZrgMLz4_jBrzk4TqXJpRKRvWy6yw8ggP7NDAwlhHXqTcPiPgheoaw_b",
			NeedOpenComment:    1,
			OnlyFansCanComment: 0,
		},
		{
			Title:              "2测试title",
			Author:             "2测试作者",
			Digest:             "2图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空",
			Content:            content2,
			ContentSourceURL:   "https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html",
			ThumbMediaID:       "HEIhl0HtYZrgMLz4_jBrzk4TqXJpRKRvWy6yw8ggP7NDAwlhHXqTcPiPgheoaw_b",
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
		//{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzsN0OyfofifBeaIk-gKHNt9-Qe4GLJoKKPe6-azsj4y4"}},
		//{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzvDwn26LA7sMY1J9tAyjJ5BXTaZk9yxLHF6A4Nk8ZT8O"}},
		//{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrznzMHcYcfrI0SP39dzwHR17GGgupzHp1Tu21PSE3kXDD"}},
		//{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzlq-GFnblA-54Po08DfgBM_3suHhSS7ZHENEK8gisIyF"}},
		//{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzgcvIAORLobR4TVCIOj4b0h4ApYhHQodxCJZfn1n_sff"}},
		//{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzjm_HzTjDYVw1TEGU0nUQgCnAM3mDSaiWWvBNccbkE5n"}},
		{"test1", args{draftId: "HEIhl0HtYZrgMLz4_jBrzt01Rbuz5Tl6d5wO5vYSTXPZ964LdqxaGbJwc491rn04"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publishID, err := Publish(
				tt.args.draftId,
			)
			if err != nil {
				t.Errorf("Publish() error = %+v", err)
				return
			}
			t.Logf("Publish() success = %+v", publishID)
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
		{"test1", args{2247483665}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publishStatus, err := PublishStatus(
				tt.args.publishID,
			)
			if err != nil {
				t.Errorf("PublishStatus() error = %+v", err)
				return
			}
			t.Logf("PublishStatus() success = %+v", publishStatus)
		})
	}
}

func TestPaginatePublish(t *testing.T) {
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
			articleList, err := PaginatePublish(
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

// 综合-发布文章
func TestPublishArticle(t *testing.T) {
	// 监控微信任务
	go util.SafeGo(TaskRun)

	type args struct {
		articles []*Article
	}

	content1 := `
<div class="content custom">
	<p>开发者可新增常用的素材到草稿箱中进行使用。上传到草稿箱中的素材被群发或发布后，该素材将从草稿箱中移除。新增草稿可在公众平台官网 - 草稿箱中查看和管理。</p> 
	<h2 id="接口请求说明"> 接口请求说明</h2> 
	<p>调用示例</p> 
	<div class="language-json extra-class">
		<pre class="language-json">
			<code>
				<span class="token punctuation">{</span>
				<span class="token property">"articles"</span>
				<span class="token comment">//若新增的是多图文素材，则此处应还有几段 articles 结构</span>
				<span class="token punctuation">]</span>
				<span class="token punctuation">}</span>
			</code>
		</pre>
	</div>
	<img title="上传图片1" src="{#1#}">
	<img title="上传图片2" src="{#2#}">
	<img title="上传图片3" src="{#3#}">
	<img title="上传图片2" src="{#2#}">
	<img title="上传图片1" src="{#1#}">
	<p>请求参数说明</p>
	<div class="table-wrp">
		<table>
			<thead>
			<tr><th>参数</th> <th>是否必须</th> <th>说明</th></tr>
			</thead> 
			<tbody>
			<tr><td>title</td> <td>是</td> <td>标题</td></tr> <tr><td>author</td> <td>否</td> <td>作者</td></tr> <tr><td>digest</td> <td>否</td> <td>图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。</td></tr> 
			<tr><td>content</td> <td>是</td> <td>图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。</td></tr> 
			<tr><td>content_source_url</td> <td>否</td> <td>图文消息的原文地址，即点击“阅读原文”后的URL</td></tr> 
			<tr><td>thumb_media_id</td> <td>是</td> <td>图文消息的封面图片素材id（必须是永久MediaID）</td></tr> 
			<tr><td>need_open_comment</td> <td>否</td> <td>Uint32 是否打开评论，0不打开(默认)，1打开</td></tr> 
			<tr><td>only_fans_can_comment</td> <td>否</td> <td>Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论</td></tr>
			</tbody>
		</table>
	</div>
	<h2 id="接口返回说明"> 接口返回说明</h2> 
	<p>返回参数说明</p> 
	<div class="table-wrp">
		<table>
		<thead><tr><th>参数</th> <th>描述</th></tr></thead> 
		<tbody><tr><td>media_id</td> <td>上传后的获取标志，长度不固定，但不会超过 128 字符</td></tr>
		</tbody>
		</table>
	</div>
</div>
`

	content2 := `
<div class="content custom">
	<p>开发者可新增常用的素材到草稿箱中进行使用。上传到草稿箱中的素材被群发或发布后，该素材将从草稿箱中移除。新增草稿可在公众平台官网 - 草稿箱中查看和管理。</p> 
	<h2 id="接口请求说明"> 接口请求说明</h2> 
	<p>调用示例</p> 
	<div class="language-json extra-class">
		<pre class="language-json">
			<code>
				<span class="token punctuation">{</span>
				<span class="token property">"articles"</span>
				<span class="token comment">//若新增的是多图文素材，则此处应还有几段 articles 结构</span>
				<span class="token punctuation">]</span>
				<span class="token punctuation">}</span>
			</code>
		</pre>
	</div>
	<img title="上传图片1" src="{#1#}">
	<img title="上传图片2" src="{#2#}">
	<img title="上传图片3" src="{#3#}">
	<img title="上传图片2" src="{#2#}">
	<img title="上传图片1" src="{#1#}">
    <p>请求参数说明</p>
	<div class="table-wrp">
		<table>
			<thead>
			<tr><th>参数</th> <th>是否必须</th> <th>说明</th></tr>
			</thead> 
			<tbody>
			<tr><td>title</td> <td>是</td> <td>标题</td></tr> <tr><td>author</td> <td>否</td> <td>作者</td></tr> <tr><td>digest</td> <td>否</td> <td>图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。</td></tr> 
			<tr><td>content</td> <td>是</td> <td>图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。</td></tr> 
			<tr><td>content_source_url</td> <td>否</td> <td>图文消息的原文地址，即点击“阅读原文”后的URL</td></tr> 
			<tr><td>thumb_media_id</td> <td>是</td> <td>图文消息的封面图片素材id（必须是永久MediaID）</td></tr> 
			<tr><td>need_open_comment</td> <td>否</td> <td>Uint32 是否打开评论，0不打开(默认)，1打开</td></tr> 
			<tr><td>only_fans_can_comment</td> <td>否</td> <td>Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论</td></tr>
			</tbody>
		</table>
	</div>
	<h2 id="接口返回说明"> 接口返回说明</h2> 
	<p>返回参数说明</p> 
	<div class="table-wrp">
		<table>
		<thead><tr><th>参数</th> <th>描述</th></tr></thead> 
		<tbody><tr><td>media_id</td> <td>上传后的获取标志，长度不固定，但不会超过 128 字符</td></tr>
		</tbody>
		</table>
	</div>
</div>
`

	articles := []*Article{
		{
			DraftArticle: draft.Article{
				Title:              "1综合测试title",
				Author:             "1综合测试作者",
				Digest:             "1图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空",
				Content:            content1,
				ContentSourceURL:   "https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html",
				ThumbMediaID:       "",
				NeedOpenComment:    1,
				OnlyFansCanComment: 0,
			},
			ContentImageFiles: []*ContentImageFile{
				{
					FilePath:    "./testmedia/2022-08-02_103202.png",
					Placeholder: "{#1#}",
				},
				{
					FilePath:    "./testmedia/process_daotu.png",
					Placeholder: "{#2#}",
				},
				{
					FilePath:    "./testmedia/String_struct.jpg",
					Placeholder: "{#3#}",
				},
			},
			CoverImageFile: "./testmedia/String_struct.jpg",
		},
		{
			DraftArticle: draft.Article{
				Title:        "2综合测试title",
				Content:      content2,
				ThumbMediaID: "",
			},
			ContentImageFiles: []*ContentImageFile{
				{
					FilePath:    "./testmedia/2022-08-02_103202.png",
					Placeholder: "{#3#}",
				},
				{
					FilePath:    "./testmedia/process_daotu.png",
					Placeholder: "{#2#}",
				},
				{
					FilePath:    "./testmedia/String_struct.jpg",
					Placeholder: "{#1#}",
				},
			},
			CoverImageFile: "./testmedia/process_daotu.png",
		},
	}

	tests := []struct {
		name string
		args args
	}{
		{"test1", args{articles}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList, err := PublishArticle(
				tt.args.articles,
			)
			if err != nil {
				t.Errorf("PaginateDraft() error = %+v", err)
				return
			}
			t.Logf("PaginateDraft() success = %+v", articleList)
		})
	}

	// 不立即结束，给异步监控发布状态一点时间
	time.Sleep(time.Minute)
}
