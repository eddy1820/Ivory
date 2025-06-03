package repository

import (
	"gate/internal/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
	"time"
)

func TestAccountRepository_GetAccountInfoByAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	account := "eddy123"
	hashedPassword := "secret"
	email := "eddy@gmail.com"
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "account", "hashed_password", "email", "created_at", "password_changed_at"}).
		AddRow(1, account, hashedPassword, email, now, now)

	query := regexp.QuoteMeta("SELECT * FROM `accounts` WHERE account = ? ORDER BY `accounts`.`id` LIMIT ?")

	mock.ExpectQuery(query).
		WithArgs(account, 1).
		WillReturnRows(rows)

	repo := NewAccountRepository(gormDB)
	result, err := repo.GetAccountInfoByAccount(account)

	assert.NoError(t, err)
	assert.Equal(t, account, result.Account)
	assert.Equal(t, hashedPassword, result.HashedPassword)
	assert.Equal(t, email, result.Email)
	assert.Equal(t, now, result.CreatedAt)
	assert.Equal(t, now, result.PasswordChangedAt)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func TestAccountRepository_InsertAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	account := domain.Account{
		Account:        "eddy123",
		HashedPassword: "secret",
		Email:          "eddy@gmail.com",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `accounts` (`account`,`hashed_password`,`email`,`password_changed_at`,`created_at`) VALUES (?,?,?,?,?)")).
		WithArgs(account.Account, account.HashedPassword, account.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewAccountRepository(gormDB)
	err = repo.InsertAccount(account)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
