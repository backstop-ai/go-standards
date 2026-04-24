package invalid

import "encoding/base64"

// Bypass attempt: credential obfuscated via base64 encoding
// ruleid: go.security.no-hardcoded-credentials
var password = func() string {
	decoded, _ := base64.StdEncoding.DecodeString("c3VwZXJzZWNyZXQ=")
	return string(decoded)
}()
