package work

import (
	"fmt"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/clients"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
	"micro-scrm/internal/pkg/http"
	"micro-scrm/internal/pkg/sdk"
	"time"
)

func Request(ctx ctx.Context, url string, req interface{}, response interface{}, options ...clients.Option) error {
	if req != nil {
		options = append(options, clients.JSONBody(req))
	}
	options = append(options, clients.Header(sdk.HeaderKeyEnum.XRequestID, ctx.ID()))

	client := http.NewHTTPClient(clients.Timeout(60 * time.Second))
	res, err := client.Post(ctx, url, options...)
	fmt.Println(res, err)
	return nil
}
