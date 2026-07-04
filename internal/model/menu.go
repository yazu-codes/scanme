package model

type Menu struct {
	ID                int64             `json:"id" gorm:"primaryKey"`
	MenuItems         []MenuItem        `json:"menu_items" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
	MenuOwner         MenuOwner         `json:"menu_owner" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
	MenuConfiguration MenuConfiguration `json:"menu_configuration" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
	Suspended         bool              `json:"suspended"`
}

type MenuItem struct {
	ID                   int64   `json:"id" gorm:"primaryKey"`
	Name                 string  `json:"name" gorm:"not null"`
	Price                float64 `json:"price" gorm:"not null"`
	Description          string  `json:"description" gorm:"not null"`
	PictureURL           string  `json:"picture_url" gorm:"not null"`
	DisplayOrderPosition int     `json:"display_order_position" gorm:"not null"`
	Category             string  `json:"category" gorm:"not null"`
	Allergens            string  `json:"allergens"`
	MenuID               int64   `json:"menu_id" gorm:"column:menu_id;not null"`
}

type MenuOwner struct {
	ID                 int64  `json:"id" gorm:"primaryKey"`
	Name               string `json:"menu_owner_name" gorm:"column:menu_owner_name;uniqueIndex;not null"`
	UrlName            string `json:"menu_owner_url_name" gorm:"column:menu_owner_url_name;uniqueIndex;not null"`
	Phone              string `json:"menu_owner_phone" gorm:"column:menu_owner_phone"`
	LogoURL            string `json:"menu_owner_logo_url" gorm:"column:menu_owner_logo_url;not null"`
	PlaceBackgroundURL string `json:"menu_owner_place_background_url" gorm:"column:menu_owner_place_background_url"`
	Slogan             string `json:"menu_owner_slogan" gorm:"column:menu_owner_slogan"`
	MenuID             int64  `json:"menu_id" gorm:"column:menu_id"`
}

type MenuConfiguration struct {
	ID              int64  `json:"id" gorm:"primaryKey"`
	BackgroundColor string `json:"background_color"`
	FontColor       string `json:"font_color"`
	FontFamily      string `json:"font_family"`
	FontSize        int    `json:"font_size"`
	MenuID          int64  `json:"menu_id" gorm:"column:menu_id"`
}

type CardMenuCode struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	MenuID int64  `json:"menu_id" gorm:"column:menu_id"`
	Code   string `json:"code" gorm:"uniqueIndex;not null"`
}
