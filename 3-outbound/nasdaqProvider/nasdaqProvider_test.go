package nasdaqProvider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetListing(t *testing.T) {
	symbols, err := GetSymbolList()
	assert.Nil(t, err)

	assert.NotEmpty(t, symbols)
}
