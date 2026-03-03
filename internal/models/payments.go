package models

import "time"

type Payment struct {
	Base
	OrderID   uint          `json:"order_id"`
	Amount    int64         `json:"amount"`
	Currency  string        `json:"currency"`
	Status    PaymentStatus `json:"status"`
	Provider  string        `json:"provider"`
	ExternalID string       `json:"external_id"`
}

type PaymentEvent struct {
	Base
	PaymentID uint             `json:"payment_id"`
	Type      PaymentEventType `json:"type"`
	Payload   string           `json:"payload"`
}

type LedgerEntry struct {
	Base
	OrderID  uint            `json:"order_id"`
	PaymentID uint           `json:"payment_id"`
	SellerID uint            `json:"seller_id"`
	Type     LedgerEntryType `json:"type"`
	Amount   int64           `json:"amount"`
	Currency string          `json:"currency"`
}

type Payout struct {
	Base
	SellerID   uint        `json:"seller_id"`
	Amount     int64       `json:"amount"`
	Currency   string      `json:"currency"`
	Status     PayoutStatus `json:"status"`
	ScheduledAt time.Time  `json:"scheduled_at"`
	PaidAt      time.Time  `json:"paid_at"`
}
