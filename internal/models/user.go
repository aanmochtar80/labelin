package models

import "time"

type User struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Username     string    `gorm:"type:varchar(100);unique;not null" json:"username"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	RoleID       uint      `json:"role_id"`
	Role         Role      `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
