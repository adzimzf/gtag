package example

import "time"

// OnlyOneTag is normal tag with only one tag
// and all fields has same tag
type OnlyOneTag struct {
	Name    string `json:"name"`
	Age     int64  `json:"age" `
	Ages    int64  `json:"no_ages"`
	WithTag string `json:"with_tag" `
}

// FullTag is a struct that has the tag is full and normal
// It'll be ordered by tag key
type FullTag struct {
	ID            int64     `gorm:"id;primary_key" db:"id" json:"id"`
	Code          string    `gorm:"code" db:"code" json:"code"`
	CategoryName  string    `gorm:"category_name" db:"category_name" json:"category_name"`
	OperatorName  string    `gorm:"operator_name" db:"operator_name" json:"operator_name"`
	Name          string    `gorm:"name" db:"name" json:"name"`
	OMSProductID  int64     `gorm:"oms_product_id" db:"oms_product_id" json:"oms_product_id"`
	UpdatedBy     int64     `gorm:"updated_by" db:"updated_by" json:"updated_by"`
	UpdatedAt     time.Time `gorm:"updated_at" db:"updated_at" json:"updated_at"`
	OverridePrice bool      `gorm:"override_price" db:"override_price" json:"override_price"`
}
