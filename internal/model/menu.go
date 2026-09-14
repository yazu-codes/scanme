package model

import "github.com/yazu-codes/scanme.git/internal/dto"

type Menus []Menu

func (m *Menus) ExtractPlaceInformation() []dto.PublicMenuOwner {
	owners := make(MenuOwners, 0, len(*m))

	for _, menu := range *m {
		owners = append(owners, menu.MenuOwner)
	}

	return owners.ToDTO()
}

type Menu struct {
	ID                int64             `json:"id" gorm:"primaryKey"`
	MenuItems         []MenuItem        `json:"menu_items" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
	MenuOwner         MenuOwner         `json:"menu_owner" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
	MenuConfiguration MenuConfiguration `json:"menu_configuration" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
	Suspended         bool              `json:"suspended"`
	YummEligible      bool              `json:"yumm_eligible" gorm:"column:yumm_eligible"`
	UserID            uint              `json:"-"`
}

type MenuItem struct {
	ID                   int64   `json:"id" gorm:"primaryKey"`
	Name                 string  `json:"name" gorm:"not null"`
	NameEn               string  `json:"name_en" gorm:""`
	Price                float64 `json:"price" gorm:"not null"`
	Description          string  `json:"description" gorm:"not null"`
	DescriptionEn        string  `json:"description_en" gorm:""`
	PictureURL           string  `json:"picture_url" gorm:"not null"`
	DisplayOrderPosition int     `json:"display_order_position" gorm:"not null"`
	Category             string  `json:"category" gorm:"not null"`
	Allergens            string  `json:"allergens"`
	Enabled              bool    `json:"enabled"`
	MenuID               int64   `json:"menu_id" gorm:"column:menu_id;not null"`
}

type MenuOwner struct {
	ID                 int64        `json:"id" gorm:"primaryKey"`
	Name               string       `json:"menu_owner_name" gorm:"column:menu_owner_name;uniqueIndex;not null"`
	UrlName            string       `json:"menu_owner_url_name" gorm:"column:menu_owner_url_name;uniqueIndex;not null"`
	Phone              string       `json:"menu_owner_phone" gorm:"column:menu_owner_phone"`
	LogoURL            string       `json:"menu_owner_logo_url" gorm:"column:menu_owner_logo_url;not null"`
	PlaceBackgroundURL string       `json:"menu_owner_place_background_url" gorm:"column:menu_owner_place_background_url"`
	Slogan             string       `json:"menu_owner_slogan" gorm:"column:menu_owner_slogan"`
	SloganEn           string       `json:"menu_owner_slogan_en" gorm:"column:menu_owner_slogan_en"`
	MenuID             int64        `json:"menu_id" gorm:"column:menu_id"`
	ReviewLinks        []ReviewLink `json:"review_links" gorm:"constraint:OnDelete:CASCADE;foreignKey:menu_id"`
}

func (mo *MenuOwner) ReviewLinksToDTO() []dto.PublicReviewLink {
	publicLinks := make([]dto.PublicReviewLink, 0, len(mo.ReviewLinks))
	for _, rl := range mo.ReviewLinks {
		publicLinks = append(publicLinks, dto.PublicReviewLink{
			URL:      rl.URL,
			Title:    rl.Title,
			ImageURL: rl.ImageURL,
		})
	}
	return publicLinks
}

type ReviewLink struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	MenuID   int64  `json:"menu_id" gorm:"column:menu_id"`
	URL      string `json:"url" gorm:"not null"`
	Title    string `json:"title" gorm:"not null"`
	ImageURL string `json:"image_url" gorm:"not null"`
}

type ReviewLinks []ReviewLink

func (rls *ReviewLinks) ToPublicReviewLinks() []dto.PublicReviewLink {
	publicLinks := make([]dto.PublicReviewLink, 0, len(*rls))
	for _, rl := range *rls {
		publicLinks = append(publicLinks, dto.PublicReviewLink{
			URL:      rl.URL,
			Title:    rl.Title,
			ImageURL: rl.ImageURL,
		})
	}
	return publicLinks
}

type MenuOwners []MenuOwner

func (m *MenuOwners) ToDTO() []dto.PublicMenuOwner {
	result := make([]dto.PublicMenuOwner, 0, len(*m))

	for _, owner := range *m {
		result = append(result, dto.PublicMenuOwner{
			Name:               owner.Name,
			Phone:              owner.Phone,
			LogoURL:            owner.LogoURL,
			Slogan:             owner.Slogan,
			SloganEn:           owner.SloganEn,
			PlaceBackgroundURL: owner.PlaceBackgroundURL,
			UrlName:            owner.UrlName,
		})
	}

	return result
}

type MenuConfiguration struct {
	ID              int64  `json:"id" gorm:"primaryKey"`
	BackgroundColor string `json:"background_color"`
	FontColor       string `json:"font_color"`
	FontFamily      string `json:"font_family"`
	FontSize        int    `json:"font_size"`
	MenuID          int64  `json:"menu_id" gorm:"column:menu_id"`
	CategoryOrder   string `json:"category_order" gorm:"column:category_order"`
	Theme           string `json:"theme" gorm:"column:theme"`
}

type CardMenuCode struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	MenuID    int64  `json:"menu_id" gorm:"column:menu_id"`
	Code      string `json:"code" gorm:"uniqueIndex;not null"`
	CustomURL string `json:"custom_url" gorm:"column:custom_url"`
}
