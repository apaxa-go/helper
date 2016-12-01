package sqlh

import "database/sql"

// SingleScannable represent object in that single row can be saved.
type SingleScannable interface {
	// SqlScanInterface return slice of interfaces which will be passed into Row.Scan at once.
	SqlScanInterface() []interface{}
}

// MultiScannable represent object in that any amount of rows can be saved.
type MultiScannable interface {
	// SqlNewElement called for each row in query result. It should returns SingleScannable object for scanning row.
	// Usually this method add new element to the underlying slice and return this element.
	SqlNewElement() SingleScannable
}

// StmtScanAll performs prepared statement stmt with arguments 'args' and stores all result rows in dst.
// StmtScanAll stop working on first error.
// Example:
//  type Label struct {
//  	Id       int32
//  	Name     string
//  }
//
//  func (l *Label) SqlScanInterface() []interface{} {
//  	return []interface{}{
//  		&l.Id,
//  		&l.Name,
//  	}
//  }
//
//  type Labels []*Label
//
//  func (l *Labels) SqlNewElement() sqlh.SingleScannable {
//	e := &Label{}
//	*l = append(*l, e)
//	return e
//  }
//  ...
//  var labels Labels
//  if err := sqlh.StmtScanAll(someStmtGetLabels, &labels, someId, someOtherParam); err != nil {
//  	return err
//  }
func StmtScanAll(stmt *sql.Stmt, dst MultiScannable, args ...interface{}) error {
	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rowContainer := dst.SqlNewElement()
		if err := rows.Scan(rowContainer.SqlScanInterface()...); err != nil {
			return err
		}
	}

	return rows.Err()
}
