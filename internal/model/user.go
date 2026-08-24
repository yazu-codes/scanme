package model

type User struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Email           string `gorm:"uniqueIndex" json:"email"`
	Password        string `json:"-"` // Never expose password in JSON
	Name            string `json:"name"`
	Role            string `json:"role"`
	AssociatedMenus string `json:"menus"`
}
