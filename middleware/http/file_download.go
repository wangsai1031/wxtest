package httpmiddleware

import (
	"io"
	"net/http"
	"os"

	"weixin/common/util"
)

const mimeJson = "application/json"

type serverWriter struct {
	basicWriter
	filePath string
	header   http.Header
}

func (s *serverWriter) Header() http.Header {
	if s.filePath == "" {
		s.filePath = util.GetDownloadFilePath(s.header)
		if s.filePath != "" {
			s.header.Del(util.XdFileDownloadHeader)
		}
	}

	return s.header
}

func (s *serverWriter) Write(b []byte) (int, error) {
	if s.filePath != "" {
		if _, ok := s.header[util.XdContentType]; ok {
			s.Header().Set("content-type", s.Header().Get(util.XdContentType))
			s.header.Del(util.XdContentType)
		}

		file, err := os.Open(s.filePath)

		if err != nil {
			return 0, err
		}

		defer file.Close()

		written, err := io.Copy(s.basicWriter, file)
		return int(written), err
	}
	return s.basicWriter.Write(b)
}

// FileDownload 处理文件下载相关逻辑
func FileDownload() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			sw := serverWriter{
				basicWriter: w,
				header:      w.Header(),
			}

			next.ServeHTTP(&sw, req)
		})
	}
}
