package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey; autoIncrement; not null"`
	Username  string    `json:"username" binding:"required" gorm:"type:varchar(255); not null"`
	Password  string    `json:"password" binding:"required" gorm:"type:text; not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime; not null"`
	Urls      []Url     `json:"urls" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) Values() map[string]any {
	return map[string]any{
		"id":         u.ID,
		"username":   u.Username,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
		"urls":       u.Urls,
	}
}
