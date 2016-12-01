// Package sqlh provides some helper functions and types to simplify working with sql package.
package sqlh

import "database/sql"

// MustPrepare is like DB.Prepare but panics if the SQL cannot be parsed.
// It simplifies safe initialization of global variables holding prepared statements.
func MustPrepare(db *sql.DB, sql string) (stmt *sql.Stmt) {
	var err error
	if stmt, err = db.Prepare(sql); err != nil {
		panic(err.Error())
	}
	return
}
