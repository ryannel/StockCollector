package nasdaqProvider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanFloatStrings(t *testing.T) {
	result, err := cleanFloatStrings("$2004,2.20")
	assert.Nil(t, err)
	assert.Equal(t, 20042.20, result)

	result, err = cleanFloatStrings("31,130,520")
	assert.Nil(t, err)
	assert.Equal(t, 31130520.0, result)
}
