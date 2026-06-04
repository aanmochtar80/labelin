package models

import "time"

type Setting struct {
	ID                uint      `gorm:"primarykey" json:"id"`
	EventName         string    `gorm:"type:varchar(255)" json:"event_name"`
	EventLogo         string    `gorm:"type:varchar(255)" json:"event_logo"`
	FooterText        string    `gorm:"type:text" json:"footer_text"`
	DefaultTemplateID *uint     `json:"default_template_id"`
	UpdatedAt         time.Time `json:"updated_at"`
}
