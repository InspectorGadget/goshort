package models

type Url struct {
	ID     uint   `json:"id" gorm:"primaryKey; autoIncrement; not null"`
	UserID uint   `json:"-" gorm:"not null"`
	Short  string `json:"short" gorm:"type:varchar(255); not null; unique"`
	Url    string `json:"url" gorm:"type:text; not null"`

	User User `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *Url) Serialize() map[string]any {
	return map[string]any{
		"id":    u.ID,
		"short": u.Short,
		"url":   u.Url,
	}
}
