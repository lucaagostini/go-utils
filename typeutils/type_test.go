package typeutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsString(t *testing.T) {
	assert.Equal(t, IsString(""), true)
	assert.Equal(t, IsString("ciao"), true)
	assert.Equal(t, IsString(1), false)
	assert.Equal(t, IsString(map[string]interface{}{"a": "b"}), false)
}
