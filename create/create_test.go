package create

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewProject(t *testing.T) {
	var err error
	err = NewProject("myproject", "", "", "")
	require.Nil(t, err, "valid myproject call should not error")

	err = NewProject("myproject", "", "", "")
	require.NotNil(t, err, "attempting to create a project with an existing name should error")

	err = NewProject("invalid:character", "", "", "")
	require.NotNil(t, err, "attempting to create a project with an invalid name should error")

	os.RemoveAll("myproject")
}

func Test_NewCreateSettings(t *testing.T) {
	var s = "myproject"
	var dbname = "deebee"
	var user = "user"
	var pwd = "pwd"
	settings := NewCreateSettings(s, dbname, user, pwd)
	assert.Equal(t, settings.ProjectName, s)
	assert.Equal(t, settings.DbName, dbname)
	assert.Equal(t, settings.DbUser, user)
	assert.Equal(t, settings.DbPassword, pwd)
}

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
	assert.NotNil(t, hasIllegalFilename("com1"))
	assert.NotNil(t, hasIllegalFilename("lpt7"))
	assert.NotNil(t, hasIllegalFilename(strings.Repeat("a", 256)))
}
