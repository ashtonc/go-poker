package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Env will serve as an environment that contains "global" variables. See
// http://www.alexedwards.net/blog/organising-database-access for more details.
type Env struct {
	Database *sql.DB
	// template cache
	// logger middleware
	// authentication middleware
}
