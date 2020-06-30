package datatest

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
type FullTag struct {
	ID        int64     `db:"id"         gorm:"id;primary_key" json:"id"`
	Code      string    `db:"code"       gorm:"code"           json:"code"`
	Address   string    `db:"address"    gorm:"address"        json:"address"`
	Name      string    `db:"name"       gorm:"name"           json:"name"`
	UpdatedBy int64     `db:"updated_by" gorm:"updated_by"     json:"updated_by"`
	UpdatedAt time.Time `db:"updated_at" gorm:"updated_at"     json:"updated_at"`
}

// IncompleteTag  a struct with incomplete t
type IncompleteTag struct {
	ID        int64     `gorm:"id;primary_key" json:"id"         db:"id"`
	Code      string    `gorm:"code"           json:"code"       db:"code"`
	Address   string    `gorm:"address"        json:"address"`
	Name      string    `gorm:"name"           json:"name"       db:"name"`
	UpdatedBy int64     `gorm:"updated_by"     json:"updated_by" db:"updated_by"`
	UpdatedAt time.Time `gorm:"updated_at"     json:"updated_at" db:"updated_at"`
}

// InvalidTag is a struct which has the invalid tag
type InvalidTag struct {
	ID        int64     `db:"id"         gorm:"id;primary_key" json:"id"`
	Code      string    `db:"code"       gorm:"code"           json:"code"`
	Address   string    `gorm:"address"        json:"address"`
	Name      string    `db:"name"       gorm:"name"           json:"name"`
	UpdatedBy int64     `db:"updated_by" gorm:"updated_by"     json:"updated_by"`
	UpdatedAt time.Time `db:"updated_at"`
}
