package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"micro-base/internal/pkg/common"
	"micro-base/internal/pkg/core/ctx"
	"strconv"
)

func Access() gin.HandlerFunc {
	return func(context *gin.Context) {
		enterpriseID, _ := strconv.ParseInt(context.GetHeader(common.HeaderKeyEnum.EnterpriseID), 10, 64)
		enterpriseHash := context.GetHeader(common.HeaderKeyEnum.EnterpriseHash)
		context.Set(common.ContextKeyEnum.EnterpriseID, enterpriseID)
		context.Set(common.ContextKeyEnum.EnterpriseHash, enterpriseHash)
		requestID := context.GetHeader(common.HeaderKeyEnum.XRequestID)
		if requestID == "" {
			requestID = uuid.NewV4().String()
		}
		context.Set(ctx.RequestIDKey, requestID)

		context.Next()
	}
}
