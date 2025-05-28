package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type AddTokenRequest struct {
	UserID    uint      `json:"user_id" binding:"required"`
	Token     string    `json:"token" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

type Token struct {
	ID        uint      `json:"id" gorm:"primaryKey; autoIncrement; not null"`
	UserID    uint      `json:"user_id" gorm:"type:int; not null"`
	Token     string    `json:"token" gorm:"type:text; not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"type:datetime; not null"`

	User User `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Token) Serialize(login bool) gin.H {
	if login {
		return gin.H{
			"message":    "Authenticated",
			"token":      t.Token,
			"expires_at": t.ExpiresAt,
		}
	}

	return gin.H{
		"token":      t.Token,
		"expires_at": t.ExpiresAt,
	}
}
