package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SetBaseUrl_GivenValidUrl_ShouldSetCorrectly(t *testing.T) {
	sut, err := NewHttpRequestBuilder("http://testUrl/testFile")
	assert.Nil(t, err)

	assert.Equal(t, "http", sut.baseUrl.Scheme)
	assert.Equal(t, "testUrl", sut.baseUrl.Host)
	assert.Equal(t, "/testFile", sut.baseUrl.Path)
	assert.Equal(t, "http://testUrl/testFile", sut.GetUrl())
}

func Test_SetBaseUrl_GivenBadUrl_ShouldError(t *testing.T) {
	_, err := NewHttpRequestBuilder("http//google.com")
	assert.Error(t, err)

	err = nil

	_, err = NewHttpRequestBuilder("google.com")
	assert.Error(t, err)

	err = nil

	_, err = NewHttpRequestBuilder("/foo/bar")
	assert.Error(t, err)
}

func Test_AddParams_GivenParams_ShouldBeAddedToUrl(t *testing.T) {
	sut, err := NewHttpRequestBuilder("http://testUrl")
	assert.Nil(t, err)

	sut.AddQueryParameter("key1", "value1")
	sut.AddQueryParameter("key2", "value2")
	url := sut.GetUrl()
	_ = url

	test := sut.params["key1"][0]
	_ = test

	assert.Equal(t, "value1", sut.params["key1"][0])
	assert.Equal(t, "value2", sut.params["key2"][0])
	assert.Equal(t, "http://testUrl?key1=value1&key2=value2", sut.GetUrl())
}