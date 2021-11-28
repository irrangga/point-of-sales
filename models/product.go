package models

type Product struct {
	ID           int    `gorm:"primary_key" json:"id"`
	MerchantID   int    `json:"merchant_id"`
	ProductName  string `json:"product_name"`
	SKU          string `json:"sku"`
	DisplayImage string `json:"display_image"`
}
