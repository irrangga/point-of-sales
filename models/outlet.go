package models

type Outlet struct {
	ID         int `gorm:"primary_key" json:"id"`
	MerchantID int `json:"merchant_id"`
	ProductID  int `json:"product_id"`
	Price      int `json:"price"`
}
