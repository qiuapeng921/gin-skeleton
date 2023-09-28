package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
	"mime"
	"net/http"
	"time"
)

// ResBodyKey 返回数据存储 key
const ResBodyKey = "/res-body"

// Logger 访问日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		p := c.Request.URL.Path
		method := c.Request.Method
		e := log.Info(ctx.Wrap(c))

		e.Str("ip", c.ClientIP())
		e.Str("method", method)
		e.Str("url", c.Request.URL.String())
		e.Str("proto", c.Request.Proto)
		e.Str("request_id", c.GetHeader("X-Request-ID"))
		e.Str("user_agent", c.GetHeader("User-Agent"))

		// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
		if method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType == "application/json" {
				body, err := ioutil.ReadAll(c.Request.Body)
				c.Request.Body.Close()
				if err == nil {
					buf := bytes.NewBuffer(body)
					c.Request.Body = ioutil.NopCloser(buf)
					e.Int64("content_length", c.Request.ContentLength)
					e.Str("body", string(body))
				}
			}
		}
		c.Next()

		timeConsuming := time.Since(start).Nanoseconds() / 1e6
		e.Int("res_status", c.Writer.Status())
		e.Int("res_length", c.Writer.Size())

		if v, ok := c.Get(ResBodyKey); ok {
			if b, ok := v.([]byte); ok {
				e.Str("res_body", string(b))
			}
		}

		e.Msgf("[http] %s-%s-%s-%d(%dms)", p, c.Request.Method, c.ClientIP(), c.Writer.Status(), timeConsuming)
	}
}
