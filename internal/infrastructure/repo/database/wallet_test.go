package database

import (
	"context"
	"errors"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/toanppp/go-clean-tx/internal/domain"
)

func mockWalletRepo(t *testing.T) (walletRepo, sqlmock.Sqlmock, func()) {
	tx, mock, closeFunc := mockTransactor(t)
	return walletRepo{transactor: tx}, mock, closeFunc
}

func TestWalletRepo_CreateWallet(t *testing.T) {
	repo, mock, closeFunc := mockWalletRepo(t)
	defer closeFunc()

	// define expectation
	want1 := domain.Wallet{
		ID:      int64(rand.Uint32() + 1),
		Balance: int64(rand.Uint32() + 1),
	}
	want2 := domain.Wallet{
		ID:      int64(rand.Uint32() + 1),
		Balance: int64(rand.Uint32() + 1),
	}

	// https://github.com/DATA-DOG/go-sqlmock/issues/118
	// For some reason, gorm uses QueryRow instead of Exec for INSERT statement.
	// sqlmock expect INSERT statement to be called with Exec, so it cannot catch the call.
	//
	// Gorm does not use Exec because Postgres can do RETURNING in a single INSERT statement,
	// but to be able to fetch this returning value QueryRow is required rather than Exec.
	// I guess by leveraging this feature, gorm does not need to do any other operation
	// to obtain primary key of the item it just inserted.
	// To mock an insert query you can do this:
	// mock.ExpectQuery(`INSERT INTO "table" (.+) RETURNING`).
	// 	WithArgs(sqlmock.AnyArg(), value1, value2).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	// The tricky part is: to use WithArgs you must know the exact number of arguments,
	// and replace those you have no control over (like CreatedAt, which is managed automatically by gorm)
	// with sqlmock.AnyArg(). If you'd like to also match for argument ordinals, well, there is still more hard work...
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO `wallets` (`balance`) VALUES (?) RETURNING `id`")).
		WithArgs(want1.Balance).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(want1.ID))
	mock.ExpectCommit()

	// We need to expect begin and commit because
	// GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency
	// https://gorm.io/docs/transactions.html
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO `wallets` (`balance`) VALUES (?) RETURNING `id`")).
		WithArgs(want2.Balance).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(want2.ID))
	mock.ExpectCommit()

	// run
	got1, err := repo.CreateWallet(context.Background(), want1.Balance)
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	got2, err := repo.CreateWallet(context.Background(), want2.Balance)
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	// assert
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("unexpected result: got: %v, want %v", got1, want1)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("unexpected result: got: %v, want %v", got2, want2)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func TestWalletRepo_CreateWallet_Error(t *testing.T) {
	repo, mock, closeFunc := mockWalletRepo(t)
	defer closeFunc()

	// define expectation
	mockErr := errors.New(strconv.FormatInt(rand.Int63(), 10))
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO `wallets` (`balance`) VALUES (?) RETURNING `id`")).
		WithArgs(sqlmock.AnyArg()).WillReturnError(mockErr)
	mock.ExpectRollback()

	// run
	_, err := repo.CreateWallet(context.Background(), int64(rand.Uint32()+1))
	if err == nil {
		t.Fatalf("not occur error")
	}

	// assert
	if err.Error() != mockErr.Error() {
		t.Errorf("unexpected error: got %v, want %v", err, mockErr)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func TestWalletRepo_GetWalletByID(t *testing.T) {
	repo, mock, closeFunc := mockWalletRepo(t)
	defer closeFunc()

	// define expectation
	want := domain.Wallet{
		ID:      int64(rand.Uint32() + 1),
		Balance: int64(rand.Uint32() + 1),
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wallets` WHERE `wallets`.`id` = ? LIMIT 1")).
		WithArgs(want.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).AddRow(want.ID, want.Balance))

	// run
	got, err := repo.GetWalletByID(context.Background(), want.ID)
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	// assert
	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected result: got: %v, want %v", got, want)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func TestWalletRepo_GetWalletByID_Error(t *testing.T) {
	repo, mock, closeFunc := mockWalletRepo(t)
	defer closeFunc()

	// define expectation
	mockErr := errors.New(strconv.FormatInt(rand.Int63(), 10))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wallets` WHERE `wallets`.`id` = ? LIMIT 1")).
		WithArgs(sqlmock.AnyArg()).WillReturnError(mockErr)

	// run
	_, err := repo.GetWalletByID(context.Background(), int64(rand.Uint32()+1))
	if err == nil {
		t.Fatalf("not occur error")
	}

	// assert
	if err.Error() != mockErr.Error() {
		t.Errorf("unexpected error: got %v, want %v", err, mockErr)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unexpected: %v", err)
	}
}
