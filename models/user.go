package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"type:int; primaryKey; autoIncrement; not null"`
	Username  string    `json:"username" gorm:"type:varchar(255); not null; unique"`
	Password  string    `json:"password" gorm:"type:text; not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime; not null"`

	Urls []Url `json:"urls" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) Serialize() map[string]any {
	return map[string]any{
		"id":         u.ID,
		"username":   u.Username,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}
