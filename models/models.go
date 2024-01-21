package models

import "time"

type Consumer struct {
	ID          int       `gorm:"id" json:"id"`
	NIK         string    `gorm:"nik" json:"nik"`
	FullName    string    `gorm:"fullname" json:"fullname"`
	BirthDate   time.Time `gorm:"birthdate" json:"birthdate"`
	Salary      float64   `gorm:"salary" json:"salary"`
	KtpPhoto    string    `gorm:"ktp_photo" json:"ktp_photo"`
	SelfiePhoto string    `gorm:"selfie_photo" json:"selfie_photo"`
}

func (Consumer) TableName() string {
	return "consumer"
}

type Transaction struct {
	ID           int64   `json:"id"`
	ConsumerID   int64   `json:"consumer_id"`
	ContractNo   string  `json:"contract_no"`
	OTR          float64 `json:"otr"`
	AdminFee     float64 `json:"admin_fee"`
	Installments int     `json:"installments"`
	InterestRate float64 `json:"interest_rate"`
	AssetName    string  `json:"asset_name"`
}

func (Transaction) TableName() string {
	return "transaction"
}
