package models

type Cart struct {
	Base
	UserID uint `json:"user_id"`
}

type CartItem struct {
	Base
	CartID   uint  `json:"cart_id"`
	SKUId    uint  `json:"sku_id"`
	Quantity int64 `json:"quantity"`
	Price    int64 `json:"price"`
	Currency string `json:"currency"`
}
