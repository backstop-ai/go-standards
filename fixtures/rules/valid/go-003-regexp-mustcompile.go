package testdata

import "regexp"

// Immutable-by-construction package-level compiled regexes are exempt from GO-003
// (no-global-mutable-state): compiled once at init, never reassigned, and
// *regexp.Regexp is safe for concurrent use — not the mutable-global hazard the rule
// targets. Recompiling per-call or threading them via DI would be strictly worse.
var namePartPattern = regexp.MustCompile(`^[A-Za-z0-9-]+$`)
var posixPattern = regexp.MustCompilePOSIX(`^[a-z]+$`)
