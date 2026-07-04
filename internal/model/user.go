package model

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"` // e.g., "admin", "user"
}
