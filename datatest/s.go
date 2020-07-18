package datatest

import "time"

type Sample struct {
	ID        int64     `gorm:"id;primary_key"db:"id"json:"id"`
	Code      string    `db:"code" json:"code" gorm:"code" `
	Address   string    `db:"address" gorm:"address" json:"address"`
	Name      string    `db:"name" gorm:"name" json:"name"`
	UpdatedBy int64     ` gorm:"updated_by"db:"updated_by"json:"updated_by"`
	UpdatedAt time.Time `db:"updated_at" gorm:"updated_at"json:"updated_at"`
}
