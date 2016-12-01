package sqlh

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

type Label struct {
	ID   int32
	Name string
}

func (l *Label) SQLScanInterface() []interface{} {
	return []interface{}{
		&l.ID,
		&l.Name,
	}
}

type Labels []*Label

func (l *Labels) SQLNewElement() SingleScannable {
	e := &Label{}
	*l = append(*l, e)
	return e
}

func TestStmtScanAll(t *testing.T) {
	test := "SELECT id, name FROM t1"

	// Prepare DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	exPrep := mock.ExpectPrepare(test)
	exPrep.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "One").AddRow(2, "Two"))

	// Perform positive tests
	stmt, err := db.Prepare(test)
	if err != nil {
		t.Fatal(err)
	}
	var labels Labels
	if err := StmtScanAll(stmt, &labels); err != nil {
		t.Error(err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	if len(labels) != 2 {
		t.Errorf("expect 2 rows, got %v", len(labels))
	} else if labels[0].ID != 1 || labels[0].Name != "One" || labels[1].ID != 2 || labels[1].Name != "Two" {
		t.Errorf("expect 1 One 2 Two, got %v %v %v %v", labels[0].ID, labels[0].Name, labels[1].ID, labels[1].Name)
	}

	// Preform negative tests 1 - invalid arguments to stmt
	if err := StmtScanAll(stmt, &labels, 0xbeef); err == nil {
		t.Error("expect error but no error")
	}

	// Additional DB preparation
	exPrep.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("1a", "One").AddRow("2b", "Two"))

	// Preform negative tests 2 - incompatible types
	if err := StmtScanAll(stmt, &labels); err == nil {
		t.Error("expect error but no error")
	}
}
