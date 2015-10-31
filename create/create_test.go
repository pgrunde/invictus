package create

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasIllegalFilename(t *testing.T) {
	assert.False(t, hasIllegalFilename("myproject"))
	assert.False(t, hasIllegalFilename("my-api"))
	assert.False(t, hasIllegalFilename("handle_it"))
	assert.False(t, hasIllegalFilename(strings.Repeat("a", 255)))

	assert.True(t, hasIllegalFilename("Wel/p"))
	assert.True(t, hasIllegalFilename("?Welp"))
	assert.True(t, hasIllegalFilename("Welp<"))
	assert.True(t, hasIllegalFilename("double>whamm>y"))
	assert.True(t, hasIllegalFilename("double>whamm>y"))
	assert.True(t, hasIllegalFilename("sl\\ashes"))
	assert.True(t, hasIllegalFilename("a:anda|"))
	assert.True(t, hasIllegalFilename("*wit\"i"))
	assert.True(t, hasIllegalFilename("myproj.doc"))
	assert.True(t, hasIllegalFilename(strings.Repeat("a", 256)))
}
