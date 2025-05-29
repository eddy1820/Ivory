package mysql

import (
	"gate/internal/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
)

func TestUserRepository_GetUserById(t *testing.T) {
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

	user := domain.User{Id: 1, AccountId: 10, Gender: "male", Name: "Eddy", Address: "xxxxx"}

	rows := sqlmock.NewRows([]string{"id", "account_id", "gender", "name", "address"}).
		AddRow(user.Id, user.AccountId, user.Gender, user.Name, user.Address)

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE account_id = ? ORDER BY `users`.`id` LIMIT ?")

	mock.ExpectQuery(query).
		WithArgs(user.AccountId, 1).
		WillReturnRows(rows)

	repo := NewUserRepository(gormDB)
	result, err := repo.GetUserById(user.AccountId)

	assert.NoError(t, err)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.AccountId, result.AccountId)
	assert.Equal(t, user.Gender, result.Gender)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Address, result.Address)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_InsertUser(t *testing.T) {
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

	user := domain.User{AccountId: 10, Gender: "male", Name: "Eddy", Address: "xxxxx"}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`account_id`,`gender`,`name`,`address`) VALUES (?,?,?,?)")).
		WithArgs(user.AccountId, user.Gender, user.Name, user.Address).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewUserRepository(gormDB)
	err = repo.InsertUser(user)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
