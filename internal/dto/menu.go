package dto

// Public-facing shapes — no DB identifiers exposed
type PublicMenu struct {
	MenuOwner         PublicMenuOwner         `json:"menu_owner"`
	MenuConfiguration PublicMenuConfiguration `json:"menu_configuration"`
	MenuItems         []PublicMenuItem        `json:"menu_items"`
}

type PublicMenuOwner struct {
	Name               string `json:"menu_owner_name"`
	Phone              string `json:"menu_owner_phone"`
	LogoURL            string `json:"menu_owner_logo_url"`
	Slogan             string `json:"menu_owner_slogan"`
	PlaceBackgroundURL string `json:"menu_owner_place_background_url"`
}

type PublicMenuConfiguration struct {
	BackgroundColor string `json:"background_color"`
	FontColor       string `json:"font_color"`
	FontFamily      string `json:"font_family"`
	FontSize        int    `json:"font_size"`
}

type PublicMenuItem struct {
	Name                 string  `json:"name"`
	Price                float64 `json:"price"`
	Description          string  `json:"description"`
	PictureURL           string  `json:"picture_url"`
	Category             string  `json:"category"`
	Allergens            string  `json:"allergens"`
	DisplayOrderPosition int     `json:"display_order_position"`
}
