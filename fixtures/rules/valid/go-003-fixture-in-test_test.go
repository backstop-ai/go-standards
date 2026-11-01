package testdata

import "time"

// Test-package fixtures are excluded from GO-003: a fixed clock shared by the
// tests in this package is not runtime state reached without injection.
var fixedNow = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
