package fcHttp

import "context"

type HttpClientInterface interface {
}

type HttpClient struct {
}

func NewHttpClient(ctx context.Context) (HttpClientInterface, error) {
	return &HttpClient{}, nil
}

func (h *HttpClient) Get(ctx context.Context) error {
	return nil
}
