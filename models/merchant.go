package models

type Merchant struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	MerchantName string `json:"merchant_name"`
}
