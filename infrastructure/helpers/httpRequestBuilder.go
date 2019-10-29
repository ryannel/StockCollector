package helpers

import (
	"fmt"
	"net/url"
	"path"
)

func NewHttpRequestBuilder(baseUrl string) (HttpRequestBuilder, error)  {
	base, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return HttpRequestBuilder{}, fmt.Errorf("unable to parse base URL: %w", err)
	}
	if base.Scheme == "" {
		return HttpRequestBuilder{}, fmt.Errorf("no schem provided for base URL")
	}
	if base.Host == "" {
		return HttpRequestBuilder{}, fmt.Errorf("no host provided for base URL")
	}

	return HttpRequestBuilder{
		baseUrl:    base,
		params:     url.Values{},
	}, nil
}

type HttpRequestBuilder struct{
	baseUrl *url.URL
	params url.Values
}

func (builder *HttpRequestBuilder) AddQueryParameter(key string, value string) {
	builder.params.Add(key, value)
}

func (builder *HttpRequestBuilder) AppendPath(newPath string) {
	builder.baseUrl.Path = path.Join(builder.baseUrl.Path, newPath)
}

func (builder *HttpRequestBuilder) GetUrl() string {
	builder.baseUrl.RawQuery = builder.params.Encode()
	return builder.baseUrl.String()
}

