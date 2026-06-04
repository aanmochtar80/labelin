package models

import "time"

type PrintLog struct {
	ID         uint          `gorm:"primarykey" json:"id"`
	UserID     uint          `json:"user_id"`
	User       User          `gorm:"foreignKey:UserID" json:"user"`
	TemplateID uint          `json:"template_id"`
	Template   LabelTemplate `gorm:"foreignKey:TemplateID" json:"template"`
	TotalLabels int          `json:"total_labels"`
	PrintTime  time.Time     `json:"print_time"`
}
