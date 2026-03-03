package models

type SellerStatus string

const (
	SellerStatusPending   SellerStatus = "pending"
	SellerStatusActive    SellerStatus = "active"
	SellerStatusSuspended SellerStatus = "suspended"
)

type Role string

const (
	RoleUser   Role = "user"
	RoleSeller Role = "seller"
	RoleAdmin  Role = "admin"
)

type ProductStatus string

const (
	ProductStatusDraft    ProductStatus = "draft"
	ProductStatusActive   ProductStatus = "active"
	ProductStatusArchived ProductStatus = "archived"
)

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusClosed    OrderStatus = "closed"
	OrderStatusCanceled  OrderStatus = "canceled"
)

type PaymentStatus string

const (
	PaymentStatusIntent    PaymentStatus = "intent"
	PaymentStatusCaptured  PaymentStatus = "captured"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusCanceled  PaymentStatus = "canceled"
)

type PayoutStatus string

const (
	PayoutStatusPending PayoutStatus = "pending"
	PayoutStatusPaid    PayoutStatus = "paid"
	PayoutStatusFailed  PayoutStatus = "failed"
)

type LedgerEntryType string

const (
	LedgerEntryCredit   LedgerEntryType = "credit"
	LedgerEntryDebit    LedgerEntryType = "debit"
	LedgerEntryFee      LedgerEntryType = "fee"
	LedgerEntryHold     LedgerEntryType = "hold"
	LedgerEntryRelease  LedgerEntryType = "release"
)

type AttributeValueType string

const (
	AttributeValueString AttributeValueType = "string"
	AttributeValueNumber AttributeValueType = "number"
	AttributeValueBool   AttributeValueType = "bool"
	AttributeValueEnum   AttributeValueType = "enum"
)

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeVideo MediaType = "video"
)

type PaymentEventType string

const (
	PaymentEventCreated  PaymentEventType = "created"
	PaymentEventCaptured PaymentEventType = "captured"
	PaymentEventFailed   PaymentEventType = "failed"
	PaymentEventRefunded PaymentEventType = "refunded"
)
