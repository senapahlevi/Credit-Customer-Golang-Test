package repositories

import (
	"Credit-Customer-Golang-Test/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// consumer
func (repo *Repository) CreateConsumer(consumer models.Consumer) (models.Consumer, error) {
	if err := repo.DB.Create(&consumer).Error; err != nil {
		return models.Consumer{}, err
	}
	return consumer, nil
}

// transaction
func (repo *Repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := repo.DB.Create(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}
