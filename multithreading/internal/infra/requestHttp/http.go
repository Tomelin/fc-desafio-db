package requestHttp

import (
	"context"
	//"golang.org/x/net/http2"
	"net/http"
)

type HttpVersion uint

const (
	v1 HttpVersion = iota + 1
	v2
)

type RequestHttpInterface interface {
	Get(ctx context.Context, url string) (*http.Response, error)
}

type RequestHttp struct {
	Ctx     context.Context
	Version HttpVersion
}

func NewRequestHttp(ctx context.Context, version ...HttpVersion) (RequestHttpInterface, error) {

	v := HttpVersion(v1)
	if len(version) > 0 {
		v = version[0]
	}
	return &RequestHttp{Ctx: ctx, Version: v}, nil
}

func (r *RequestHttp) Get(ctx context.Context, url string) (*http.Response, error) {

	client := http.Client{}
	//if r.Version == 2 {
	//	client.Transport = &http2.Transport{
	//		AllowHTTP:          true,
	//		DisableCompression: false,
	//	}
	//}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
