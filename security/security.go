package security

import (
	"github.com/openimchat/go-weapp/v3/request"
)

type Security struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	combineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewSecurity(request *request.Request, combineURI func(url string, req interface{}, withToken bool) (string, error)) *Security {
	sm := Security{
		request:    request,
		combineURI: combineURI,
	}

	return &sm
}
