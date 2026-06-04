package models

import "time"

type Guest struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	FullName    string    `gorm:"type:varchar(255);not null" json:"full_name"`
	PrefixTitle string    `gorm:"type:varchar(50)" json:"prefix_title"`
	SuffixTitle string    `gorm:"type:varchar(50)" json:"suffix_title"`
	Institution string    `gorm:"type:varchar(255)" json:"institution"`
	Address     string    `gorm:"type:text" json:"address"`
	City        string    `gorm:"type:varchar(100)" json:"city"`
	Province    string    `gorm:"type:varchar(100)" json:"province"`
	PostalCode  string    `gorm:"type:varchar(20)" json:"postal_code"`
	Phone       string    `gorm:"type:varchar(50)" json:"phone"`
	CategoryID  *uint     `json:"category_id"`
	Category    *Category `gorm:"foreignKey:CategoryID" json:"category"`
	TableNum    string    `gorm:"type:varchar(50)" json:"table_num"`
	Notes       string    `gorm:"type:text" json:"notes"`
	PrintStatus bool      `gorm:"default:false" json:"print_status"`
	QRCode      string    `gorm:"type:varchar(255)" json:"qrcode"` // Optional for later
	Token       string    `gorm:"type:varchar(255)" json:"token"` // Optional for later
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
