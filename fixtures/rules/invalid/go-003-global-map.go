package testdata

// A package-level map is real mutable state and still flags, next to the
// exclusions this rule allows.
var Registry = map[string]int{}

func register(name string) {
	var n = len(Registry)
	Registry[name] = n
}
