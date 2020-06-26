package example

import "time"

// OnlyOneTag is normal tag with only one tag
// and all fields has same tag
type OnlyOneTag struct {
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	WithTag string `json:"with_tag"`
	Ages    int64  `json:"no_ages"`
}
type OnlyOneTagWithAnonymousField struct {
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	WithTag    string `json:"with_tag"`
	OnlyOneTag `json:"no_ages"`
}

// FullTag is a struct that has the tag is full and normal
// It'll be ordered by tag key
type FullTag struct {
	ID            int64     `db:"id" gorm:"id;primary_key" json:"id"`
	Code          string    `db:"code"           gorm:"code" json:"code"`
	CategoryName  string    `db:"category_name"  gorm:"category_name"  json:"category_name"`
	OperatorName  string    `db:"operator_name"  gorm:"operator_name"  json:"operator_name"`
	Name          string    `db:"name"           gorm:"name" json:"name"`
	UpdatedBy     int64     `db:"updated_by"     gorm:"updated_by"     json:"updated_by"`
	UpdatedAt     time.Time `db:"updated_at"       json:"updated_at" gorm:"updated_at"   `
	OverridePrice bool      `gorm:"override_price"db:"override_price"  json:"override_price"`
}

// FullTag is a struct that has the tag is full and normal
// It'll be ordered by tag key
type IncompleteTag struct {
	ID            int64     `db:"id"              json:"id"`
	Code          string    `db:"code"           gorm:"code"           json:"code"`
	CategoryName  string    `db:"category_name"  gorm:"category_name"  json:"category_name"`
	Name          string    `gorm:"name" db:"name"            json:"name"`
	UpdatedBy     int64     `db:"updated_by"          json:"updated_by" gorm:"updated_by"`
	UpdatedAt     time.Time `db:"updated_at"     gorm:"updated_at"     json:"updated_at"`
	OverridePrice bool      `db:"override_price" gorm:"override_price" json:"override_price"`
	OperatorName  string    `db:"operator_name"    json:"operator_name"`
}
