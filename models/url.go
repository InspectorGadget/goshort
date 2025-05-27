package models

type Url struct {
	ID     uint   `json:"id" gorm:"primaryKey; autoIncrement; not null"`
	UserID uint   `json:"-" gorm:"not null"`
	Short  string `json:"short" binding:"required" gorm:"type:varchar(255); not null; unique"`
	Url    string `json:"url" binding:"required" gorm:"type:text; not null"`
}
