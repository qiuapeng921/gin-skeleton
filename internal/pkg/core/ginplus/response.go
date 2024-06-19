package ginplus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/errorx"
	"micro-base/internal/pkg/core/log"
	"net/http"
)

const (
	// UnknownErrCode 未知错误业务代码
	UnknownErrCode = 999
	// ParamsErrCode 参数错误业务代码
	ParamsErrCode = 400
)

// ---------------------- 返回结构定义 ------------------------

// APIStatusText 定义HTTP状态文本
type APIStatusText string

func (t APIStatusText) String() string {
	return string(t)
}

// OKStatusText 定义HTTP状态文本常量
const (
	OKStatusText APIStatusText = "OK"
)

// APIResponse HTTP响应错误项
type APIResponse struct {
	Code      int         `json:"code" desc:"业务状态码"`     // 状态码
	Message   string      `json:"message" desc:"业务状态信息"` // 状态信息
	RequestID string      `json:"request_id" desc:"请求ID"`
	Data      interface{} `json:"data,omitempty"desc:"业务数据"` // 状态信息
}

// APIStatus HTTP响应状态
type APIStatus struct {
	Status string `json:"status"` // 状态(OK)
}

// APIList HTTP响应列表数据
type APIList struct {
	List       interface{}    `json:"list"`
	Pagination *APIPagination `json:"pagination,omitempty"`
}

// APIPagination HTTP分页数据
type APIPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageIndex int // 页索引
	PageSize  int // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total int // 总数据条数
}

// ---------------------- 响应返回辅助方法 ------------------------

// ResPage 响应分页数据
func ResPage(c *gin.Context, v interface{}, pr *APIPagination) {
	list := APIList{
		List:       v,
		Pagination: pr,
	}
	ResSuccess(c, list)
}

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	ResSuccess(c, APIList{List: v})
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	ResSuccess(c, APIStatus{Status: OKStatusText.String()})
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	response := APIResponse{
		Code:      http.StatusOK,
		Message:   "success",
		RequestID: ctx.Wrap(c).ID(),
		Data:      v,
	}

	ResJSON(c, http.StatusOK, response)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Data(status, "application/json; charset=utf-8", buf)
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	cc := ctx.Wrap(c)
	response := APIResponse{
		Code:      UnknownErrCode,
		Message:   "未知错误信息",
		RequestID: ctx.Wrap(c).ID(),
	}

	if err != nil {
		response.Message = err.Error()
		var e *errorx.Error
		if errors.As(err, &e) {
			response.Code = e.Code
		}
	}

	if err != nil {
		response.Message = err.Error()
	}

	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	if statusCode >= 400 && statusCode < 500 {
		log.Warn(cc).Msg(response.Message)
	} else if statusCode >= 500 {
		span := log.Error(cc)
		span = span.Str("stack", fmt.Sprintf("%+v", err))
		span.Msg(response.Message)
	}

	ResJSON(c, statusCode, response)
}
