package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"weixin/common/handlers/conf"
)

const (
	XdFileDownloadHeader = "xd-internal-file-download"
	XdContentType        = "Xd-Content-Type"
)

type AnnexInfo struct {
	AnnexName  string
	AnnexKey   string
	AnnexSize  int64
	AnnexModel string
}

type UploadFileReq struct {
	Namespace string
}

//支持预览的文件类型

var CanPreview = map[string]bool{
	".html": true,
	".pdf":  true,
	".jpg":  true,
	".png":  true,
	".jpeg": true,
}

const (
	CAN_PREVIEW_1 = 1 //可预览
	CAN_PREVIEW_2 = 2 //不可预览
)

// 获取下载路径
func GetDownloadFilePath(header http.Header) string {
	return header.Get(XdFileDownloadHeader)
}

type FileList struct {
	FilePath string
	FineName string
}

//GetPreviewPath 获取预览地址
func GetPreviewPath(key, namespace string) (previewUrl string) {
	if key == "" {
		return ""
	}
	u := url.Values{}
	u.Set("namespace", namespace)
	u.Set("key", key)

	previewUrl = fmt.Sprintf("%s/recruit/basic/preview?%s", conf.Viper.GetString("priview.domain"), u.Encode())
	return
}

//ClearDir 清空指定目录下的文件（dirPath  /Users/shulv/tmp/）
func ClearDir(dirPath string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, d := range dir {
		os.RemoveAll(dirPath + d.Name())
	}
	return nil
}
