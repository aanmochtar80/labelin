package models

import "time"

type LabelTemplate struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Width        float64   `json:"width"`         // in mm
	Height       float64   `json:"height"`        // in mm
	MarginTop    float64   `json:"margin_top"`    // in mm
	MarginLeft   float64   `json:"margin_left"`   // in mm
	SpacingX     float64   `json:"spacing_x"`     // in mm
	SpacingY     float64   `json:"spacing_y"`     // in mm
	Columns      int       `json:"columns"`
	Rows         int       `json:"rows"`
	ElementsJSON string    `gorm:"type:text" json:"elements_json"` // Store JSON mapping for fields
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
