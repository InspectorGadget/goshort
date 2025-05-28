package models

type Role struct {
	ID      uint   `json:"id" gorm:"primaryKey; autoIncrement; not null"`
	Name    string `json:"name" gorm:"type:varchar(255); not null; unique"`
	AddedBy uint   `json:"-" gorm:"type:int; not null"`

	User User `json:"-" gorm:"foreignKey:AddedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AddRoleMapRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
}

type RoleMap struct {
	ID     uint `json:"id" gorm:"primaryKey; autoIncrement; not null"`
	UserID uint `json:"user_id" gorm:"type:int; not null"`
	RoleID uint `json:"role_id" gorm:"type:int; not null"`

	User User `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role Role `json:"-" gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
