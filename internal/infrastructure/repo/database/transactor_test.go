package database

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	// init sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// expect check version query
	rows := sqlmock.NewRows([]string{"database/sql/driver.Rows"}).AddRow(true)
	mock.ExpectQuery("select sqlite_version()").WithArgs().WillReturnRows(rows)

	// init gorm repo
	dialector := sqlite.Dialector{
		Conn: db,
	}

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("cannot initialize gorm session: %v", err)
	}

	return gormDB, mock, func() { _ = db.Close() }
}

func mockTransactor(t *testing.T) (transactor, sqlmock.Sqlmock, func()) {
	db, mock, closeFunc := mockDB(t)
	return transactor{db: db}, mock, closeFunc
}

func TestTransactor_WithinTransaction(t *testing.T) {
	repo, mock, closeFunc := mockTransactor(t)
	defer closeFunc()

	// define expectation
	mock.ExpectBegin()
	mock.ExpectCommit()

	// run test
	if err := repo.WithinTransaction(context.Background(), func(ctx context.Context) error {
		return nil
	}); err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unexpected: %v", err)
	}
}
