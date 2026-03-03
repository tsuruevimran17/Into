package models

type Seller struct {
	Base
	UserID uint         `json:"user_id"`
	Status SellerStatus `json:"status"`
}

type SellerProfile struct {
	Base
	SellerID     uint   `json:"seller_id"`
	DisplayName  string `json:"display_name"`
	Description  string `json:"description"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
}
