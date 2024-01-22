package repositories

import (
	"Credit-Customer-Golang-Test/models"
	"errors"
	"sync"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// consumer
func (repo *Repository) CreateConsumer(consumer models.Consumer) (models.Consumer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var existingDataConsumer models.Consumer
	result := repo.DB.Where("nik = ?", consumer.NIK).First(&existingDataConsumer)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Consumer{}, result.Error
	}
	if existingDataConsumer.ID != 0 {
		return models.Consumer{}, errors.New("NIK Number Already exist")
	}

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
