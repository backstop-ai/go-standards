package testdata

import (
	"embed"
	"io/fs"
)

// GO-003 (no-global-mutable-state) targets package-level state. None of these is
// that: a local variable, a blank-identifier interface assertion, and a typed
// re-export of a read-only embedded filesystem.

var _ fs.FS = embed.FS{}

var rootFS embed.FS

var ReexportFS embed.FS = rootFS

func count(xs []int) int {
	var total = 0
	for _, x := range xs {
		total += x
	}
	return total
}

type summer struct{}

func (summer) sum(xs []int) (out int) {
	var acc = 0
	f := func() { var step = 1; acc += step }
	for range xs {
		f()
	}
	return acc
}
