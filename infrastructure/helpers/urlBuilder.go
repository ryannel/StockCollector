package helpers

import (
	"fmt"
	"net/url"
	"path"
)

func NewUrlBuilder(baseUrl string) (UrlBuilder, error)  {
	base, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return UrlBuilder{}, fmt.Errorf("unable to parse base URL: %w", err)
	}
	if base.Scheme == "" {
		return UrlBuilder{}, fmt.Errorf("no schem provided for base URL")
	}
	if base.Host == "" {
		return UrlBuilder{}, fmt.Errorf("no host provided for base URL")
	}

	query, err := url.ParseQuery(base.RawQuery)
	if base.Host == "" {
		return UrlBuilder{}, fmt.Errorf("unable to pass query params")
	}

	return UrlBuilder{
		baseUrl:    base,
		params:     query,
	}, nil
}

type UrlBuilder struct{
	baseUrl *url.URL
	params url.Values
}

func (builder *UrlBuilder) AddQueryParameter(key string, value string) {
	builder.params.Add(key, value)
}

func (builder *UrlBuilder) AppendPath(newPath string) {
	builder.baseUrl.Path = path.Join(builder.baseUrl.Path, newPath)
}

func (builder *UrlBuilder) GetUrl() string {
	builder.baseUrl.RawQuery = builder.params.Encode()
	return builder.baseUrl.String()
}

