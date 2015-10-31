package create

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasIllegalFilename(t *testing.T) {
	assert.Nil(t, hasIllegalFilename("myproject"))
	assert.Nil(t, hasIllegalFilename("my-api"))
	assert.Nil(t, hasIllegalFilename("handle_it"))
	assert.Nil(t, hasIllegalFilename(strings.Repeat("a", 255)))

	assert.NotNil(t, hasIllegalFilename(""))
	assert.NotNil(t, hasIllegalFilename("^"))
	assert.NotNil(t, hasIllegalFilename("Wel/p"))
	assert.NotNil(t, hasIllegalFilename("?Welp"))
	assert.NotNil(t, hasIllegalFilename("Welp<"))
	assert.NotNil(t, hasIllegalFilename("double>whamm>y"))
	assert.NotNil(t, hasIllegalFilename("double>whamm>y"))
	assert.NotNil(t, hasIllegalFilename("sl\\ashes"))
	assert.NotNil(t, hasIllegalFilename("a:anda|"))
	assert.NotNil(t, hasIllegalFilename("*wit\"i"))
	assert.NotNil(t, hasIllegalFilename("myproj.doc"))
	assert.NotNil(t, hasIllegalFilename(strings.Repeat("a", 256)))
}
