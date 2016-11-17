package sqlh

import "database/sql"

// MustPrepare is like DB.Prepare but panics if the SQL cannot be parsed.
// It simplifies safe initialization of global variables holding prepared statements.
func MustPrepare(db *sql.DB, sql string) (stmt *sql.Stmt) {
	var err error
	if stmt, err = db.Prepare(sql); err != nil {
		panic(`sqlhelper: Prepare(` + sql + `): ` + err.Error())
	}
	return
}
