package nasdaqProvider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetListing(t *testing.T) {
	sut := Nasdaq{}
	symbols, err := sut.GetSymbolList()
	assert.Nil(t, err)

	assert.NotEmpty(t, symbols)
}
