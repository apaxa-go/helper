package sqlh

// TODO what is it?
/*
import "database/sql"

// MultiScannableConverter represent object in that any amount of rows can be saved.
// Each rows saved into
type MultiScannableConverter interface {
	// SingleScannable allow scan single row.
	SingleScannable
	// ConvertFromSql called after each successfully row scan to convert data from sql representation to required representation.
	ConvertFromSql() error
}

func StmtScanConvertAll(stmt *sql.Stmt, dst MultiScannableConverter, args ...interface{}) error {
	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(dst.SqlScanInterface()...); err != nil {
			return err
		}
		if err := dst.ConvertFromSql(); err != nil {
			return err
		}
	}

	return rows.Err()
}
*/
