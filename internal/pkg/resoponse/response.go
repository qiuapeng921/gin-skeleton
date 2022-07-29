package response

import (
	"github.com/gin-gonic/gin"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/errorx"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ginplus"
	"net/http"
)

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

// SuccessList 成功列表返回格式
func SuccessList(c *gin.Context, total int64, page, size int, list interface{}) {
	Success(c, List{
		Total: total,
		Page:  page,
		Size:  size,
		List:  list,
	})
}

// Success 返回格式
func Success(c *gin.Context, obj interface{}) {
	data := Response{
		Code:      http.StatusOK,
		Message:   "success",
		RequestID: ctx.Wrap(c).ID(),
		Data:      obj,
	}

	ginplus.ResJSON(c, http.StatusOK, data)
}

// Notice 返回格式
func Notice(c *gin.Context, e *errorx.Error, v ...interface{}) {
	var obj interface{}
	if len(v) > 0 {
		obj = v[0]
	}
	data := Response{
		Code:      e.Code,
		Message:   e.Err.Error(),
		RequestID: ctx.Wrap(c).ID(),
		Data:      obj,
	}

	ginplus.ResJSON(c, http.StatusOK, data)
}
