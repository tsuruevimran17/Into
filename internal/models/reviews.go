package models

type Review struct {
	Base
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	OrderID   uint   `json:"order_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}

type ProductRating struct {
	Base
	ProductID    uint    `json:"product_id"`
	Average      float64 `json:"average"`
	ReviewsCount int64   `json:"reviews_count"`
}
