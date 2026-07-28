package testdata

import (
	"os"
	"testing"
)

// TestIgnoredErrorInTestFile carries the EXACT discarded-second-value shape that
// the GO-010 negative fixture (go-010-ignored-error.go) uses, but in a *_test.go
// file. It is a POSITIVE fixture: go.core.no-ignored-errors excludes test files,
// because the rule matches a discarded second value rather than a discarded
// error and cannot tell the two apart. Genuine unchecked errors in tests are
// covered by the go-toolchain errcheck pass, which knows the real types.
//
// This is the exclusion's falsifier — before the paths:exclude was added, this
// file produced a finding and this fixture failed.
func TestIgnoredErrorInTestFile(t *testing.T) {
	_, _ = os.ReadFile("missing.txt")
}
