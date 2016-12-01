package sqlh

import "testing"
import "gopkg.in/DATA-DOG/go-sqlmock.v1"

func TestMustPrepare(t *testing.T) {
	test := "SELECT F1 FROM T1"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectPrepare(test)
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	}()
	MustPrepare(db, test)
}

func TestMustPreparePanic(t *testing.T) {
	test := "SELECT F1 FROM T1"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer func() {
		if r := recover(); r == nil {
			t.Error(r)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	}()
	MustPrepare(db, test)
}
