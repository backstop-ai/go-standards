package invalid

import (
	"database/sql"
	"fmt"
)

// Bypass attempt: using fmt.Sprintf instead of direct concatenation
// ruleid: go.security.no-sql-concatenation
func getUser(db *sql.DB, name string) error {
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name)
	_, err := db.Query(query)
	return err
}
