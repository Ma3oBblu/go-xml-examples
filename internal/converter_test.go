package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertIdsToXmlForOrder(t *testing.T) {
	raw, err := ConvertIdsToXmlForNews([]string{"a", "b", "c"})

	assert.NoError(t, err)
	assert.Equal(t, "<news><id>a</id><id>b</id><id>c</id></news>", raw)
}
