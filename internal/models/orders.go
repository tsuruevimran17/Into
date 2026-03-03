package models

import "time"

type Order struct {
	Base
	UserID       uint        `json:"user_id"`
	Status       OrderStatus `json:"status"`
	TotalAmount  int64       `json:"total_amount"`
	Currency     string      `json:"currency"`
	PlacedAt     time.Time   `json:"placed_at"`
	DeliveredAt  time.Time   `json:"delivered_at"`
}

type OrderItem struct {
	Base
	OrderID  uint  `json:"order_id"`
	SellerID uint  `json:"seller_id"`
	SKUId    uint  `json:"sku_id"`
	Quantity int64 `json:"quantity"`
	Price    int64 `json:"price"`
	Currency string `json:"currency"`
}
