package models

type Outlet struct {
	ID         int    `gorm:"primary_key" json:"id"`
	MerchantID int    `json:"merchant_id"`
	OutletName string `json:"outlet_name"`
}
