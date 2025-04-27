package testjson

import (
	"testing"

	"github.com/v1v/v3/assert"
	"github.com/v1v/v3/skip"
)

func TestRelativePackagePath(t *testing.T) {
	prefix := "github.com/v1v/gotestsum/testjson"
	patchPkgPathPrefix(t, prefix)
	relPath := RelativePackagePath(prefix + "/extra/relpath")
	assert.Equal(t, relPath, "extra/relpath")

	relPath = RelativePackagePath(prefix)
	assert.Equal(t, relPath, ".")
}

func TestGetPkgPathPrefix(t *testing.T) {
	t.Run("with go path", func(t *testing.T) {
		skip.If(t, isGoModuleEnabled())
		assert.Equal(t, getPkgPathPrefix(), "github.com/v1v/gotestsum/testjson")
	})
	t.Run("with go modules", func(t *testing.T) {
		skip.If(t, !isGoModuleEnabled())
		assert.Equal(t, getPkgPathPrefix(), "github.com/v1v/gotestsum")
	})
}
