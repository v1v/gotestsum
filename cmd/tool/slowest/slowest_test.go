package slowest

import (
	"bytes"
	"testing"

	"github.com/v1v/v3/env"
	"github.com/v1v/v3/golden"
)

func TestUsage_WithFlagsFromSetupFlags(t *testing.T) {
	env.PatchAll(t, nil)

	name := "gotestsum tool slowest"
	flags, _ := setupFlags(name)
	buf := new(bytes.Buffer)
	usage(buf, name, flags)

	golden.Assert(t, buf.String(), "cmd-flags-help-text")
}
