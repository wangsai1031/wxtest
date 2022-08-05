package httpmiddleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"weixin/common/consts"
)

// InitPage
func InitPage() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if strings.Contains(req.Header.Get("Content-Type"), "application/json") {
				limit := consts.DefaultLimit
				offset := consts.DefaultOffset

				body := make([]byte, req.ContentLength)
				req.Body.Read(body)
				jsonStr := string(body)

				var mapResult map[string]interface{}
				err := json.Unmarshal([]byte(jsonStr), &mapResult)
				if err != nil {
					// 记录日志
					return
				}

				if pn, okPageNum := mapResult["page_num"]; okPageNum {
					if pageNum, typeOk := pn.(float64); typeOk {
						limit = int64(pageNum)
					}
					if page, okPage := mapResult["page"]; okPage {
						if p, pageTypeOk := page.(float64); pageTypeOk {
							pageNumber := int64(p)
							if pageNumber > 0 {
								offset = (pageNumber - 1) * limit
							} else {
								offset = pageNumber
							}
						}
					}
					if limit > consts.DefaultMaxLimit {
						limit = consts.DefaultMaxLimit
					}
					mapResult["limit"] = limit
					mapResult["offset"] = offset
				}

				bufBody, errBuf := json.Marshal(&mapResult)
				if errBuf != nil {
					// 记录日志

					return
				}
				req.Body = ioutil.NopCloser(bytes.NewBuffer(bufBody))
			}

			next.ServeHTTP(w, req)
		})
	}
}
