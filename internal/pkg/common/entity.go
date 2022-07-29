package common

// ContextKeyEnum ctx.Context中存储的数据key枚举
var ContextKeyEnum = struct {
	TokenInfo      string
	EnterpriseHash string
	EnterpriseID   string
}{
	TokenInfo:      "token_info",
	EnterpriseHash: "enterprise_hash",
	EnterpriseID:   "enterprise_id",
}

// HeaderKeyEnum 请求Header中存储的key枚举
var HeaderKeyEnum = struct {
	AuthToken      string
	EnterpriseHash string
	EnterpriseID   string
	XRequestID     string
}{
	AuthToken:      "Auth-Token",
	EnterpriseHash: "Enterprise-Hash",
	EnterpriseID:   "Enterprise-ID",
	XRequestID:     "X-Request-ID",
}
