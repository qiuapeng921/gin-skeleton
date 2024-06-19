package handle

import (
	"github.com/gin-gonic/gin"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/ginplus"
)

type BaseHandle struct {
}

// Response 响应结构体
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

// List 结构体
type List struct {
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	List  interface{} `json:"list"`
}

func (b BaseHandle) ResSuccess(c *gin.Context, v interface{}) {
	ginplus.ResSuccess(c, v)
	return
}

func (b BaseHandle) ResError(c *gin.Context, err error) {
	ginplus.ResError(c, err)
	return
}

func (b BaseHandle) ResOk(c *gin.Context) {
	ginplus.ResOK(c)
	return
}

func (b BaseHandle) ResJson(c *gin.Context, status int, v interface{}) {
	ginplus.ResJSON(c, status, v)
	return
}

func (b BaseHandle) ResList(c *gin.Context, total int64, page, size int, list interface{}) {
	b.ResSuccess(c, List{
		Total: total,
		Page:  page,
		Size:  size,
		List:  list,
	})
}

func (b BaseHandle) WarpContext(c *gin.Context) ctx.Context {
	return ctx.Wrap(c)
}
