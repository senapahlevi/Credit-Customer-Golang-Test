package repositories

import (
	"Credit-Customer-Golang-Test/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

func TestCreateConsumer(t *testing.T) {
	db, mock, err := SetupMock()
	assert.NoError(t, err)

	repo := NewRepository(db)

	consumer := models.Consumer{
		NIK:         "940123123",
		FullName:    "sena",
		Salary:      20000,
		KtpPhoto:    "img/a.png",
		SelfiePhoto: "img/a.png",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `consumer`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := repo.CreateConsumer(consumer)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// func TestCreateTransaction(t *testing.T) {
// 	db, mock, err := SetupMock()
// 	assert.NoError(t, err)

// 	repo := NewRepository(db)

// 	transaction := models.Transaction{
// 		// Isi dengan data dummy
// 	}

// 	// Setup expectations
// 	mock.ExpectBegin()
// 	mock.ExpectExec("INSERT INTO `transactions`").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// Test
// 	result, err := repo.CreateTransaction(transaction)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, result)

// 	// Pastikan semua ekspektasi terpenuhi
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }
