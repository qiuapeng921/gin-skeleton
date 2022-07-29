package api

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/docs/swagger"
	"micro-base/internal/modules/activity/delivery/http"
)

func endpoints() []*swagger.Endpoint {
	var eps []*swagger.Endpoint

	eps = append(eps, http.Endpoints()...)
	return eps
}
