package models

type Category struct {
	Base
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ParentID  *uint  `json:"parent_id"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
}

type Attribute struct {
	Base
	Name         string             `json:"name"`
	ValueType    AttributeValueType `json:"value_type"`
	IsFilterable bool               `json:"is_filterable"`
}

type CategoryAttribute struct {
	CategoryID uint `json:"category_id"`
	AttributeID uint `json:"attribute_id"`
	IsRequired bool `json:"is_required"`
}

type Product struct {
	Base
	SellerID    uint          `json:"seller_id"`
	CategoryID  uint          `json:"category_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      ProductStatus `json:"status"`
}

type SKU struct {
	Base
	ProductID uint   `json:"product_id"`
	Code      string `json:"code"`
	Price     int64  `json:"price"`
	Currency  string `json:"currency"`
}

type ProductMedia struct {
	Base
	ProductID uint      `json:"product_id"`
	URL       string    `json:"url"`
	Type      MediaType `json:"type"`
	SortOrder int       `json:"sort_order"`
}

type Inventory struct {
	Base
	SKUId     uint  `json:"sku_id"`
	Quantity  int64 `json:"quantity"`
	Reserved  int64 `json:"reserved"`
}
