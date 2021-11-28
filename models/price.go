package models

type Price struct {
	ID        int `gorm:"primary_key" json:"id"`
	ProductID int `json:"product_id"`
	OutletID  int `json:"outlet_id"`
	Price     int `json:"price"`
}
