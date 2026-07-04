package repository

import (
	"errors"

	"github.com/yazu-codes/scanme.git/internal/dto"
	"github.com/yazu-codes/scanme.git/internal/model"
	"gorm.io/gorm"
)

type MenuRepository struct {
	DB *gorm.DB
}

func (m *MenuRepository) GetAllMenus() ([]model.Menu, error) {
	var menus []model.Menu
	if err := m.DB.Preload("MenuItems").Preload("MenuOwner").Preload("MenuConfiguration").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (m *MenuRepository) GetMenuByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	if err := m.DB.First(&menu, id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (m *MenuRepository) GetMenuByName(name string) (*model.Menu, error) {
	var menu model.Menu
	if err := m.DB.Where("menu_owner_name = ?", name).First(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (m *MenuRepository) GetMenuByUrlName(urlName string) (*dto.PublicMenu, error) {
	var menu model.Menu

	err := m.DB.
		Joins("JOIN menu_owners ON menu_owners.menu_id = menus.id").
		Where("menu_owners.menu_owner_url_name = ? AND menus.suspended = false", urlName).
		Preload("MenuOwner").
		Preload("MenuConfiguration").
		Preload("MenuItems").
		First(&menu).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	items := make([]dto.PublicMenuItem, 0, len(menu.MenuItems))
	for _, it := range menu.MenuItems {
		items = append(items, dto.PublicMenuItem{
			Name:                 it.Name,
			Price:                it.Price,
			Description:          it.Description,
			PictureURL:           it.PictureURL,
			DisplayOrderPosition: it.DisplayOrderPosition,
			Category:             it.Category,
			Allergens:            it.Allergens,
		})
	}

	public := &dto.PublicMenu{
		MenuOwner: dto.PublicMenuOwner{
			Name:               menu.MenuOwner.Name,
			Phone:              menu.MenuOwner.Phone,
			LogoURL:            menu.MenuOwner.LogoURL,
			Slogan:             menu.MenuOwner.Slogan,
			PlaceBackgroundURL: menu.MenuOwner.PlaceBackgroundURL,
		},
		MenuConfiguration: dto.PublicMenuConfiguration{
			BackgroundColor: menu.MenuConfiguration.BackgroundColor,
			FontColor:       menu.MenuConfiguration.FontColor,
			FontFamily:      menu.MenuConfiguration.FontFamily,
			FontSize:        menu.MenuConfiguration.FontSize,
		},
		MenuItems: items,
	}

	return public, nil
}

func (m *MenuRepository) CreateMenu(menu *model.Menu) error {
	// if menu.MenuConfiguration.ID == 0 {
	// 	if err := m.DB.Create(&menu.MenuConfiguration).Error; err != nil {
	// 		return err
	// 	}
	// }

	// if menu.MenuOwner.ID == 0 {
	// 	if err := m.DB.Create(&menu.MenuOwner).Error; err != nil {
	// 		return err
	// 	}
	// }

	// if menu.MenuItems != nil {
	// 	for i := range menu.MenuItems {
	// 		menu.MenuItems[i].MenuID = menu.ID
	// 	}
	// }

	if err := m.DB.Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func (m *MenuRepository) UpdateMenu(menu *model.Menu) error {
	tx := m.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&model.Menu{}).
		Where("id = ?", menu.ID).
		Update("suspended", menu.Suspended).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Make sure the foreign key is right regardless of what the client sent
	menu.MenuOwner.MenuID = menu.ID
	if err := tx.Model(&model.MenuOwner{}).
		Where("menu_id = ?", menu.ID).
		Omit("id").
		Omit("menu_id").
		Updates(&menu.MenuOwner).Error; err != nil {
		tx.Rollback()
		return err
	}

	menu.MenuConfiguration.MenuID = menu.ID
	if err := tx.Model(&model.MenuConfiguration{}).
		Where("menu_id = ?", menu.ID).
		Omit("id").
		Omit("menu_id").
		Updates(&menu.MenuConfiguration).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range menu.MenuItems {
		item.MenuID = menu.ID
		if item.ID == 0 {
			if err := tx.Create(&item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := tx.Model(&model.MenuItem{}).
				Where("id = ?", item.ID).
				Omit("id").
				Omit("menu_id").
				Updates(&item).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (m *MenuRepository) DeleteMenu(id uint) error {
	if err := m.DB.Delete(&model.Menu{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *MenuRepository) SuspendMenu(id uint) error {
	if err := m.DB.Model(&model.Menu{}).Where("id = ?", id).Update("suspended", true).Error; err != nil {
		return err
	}
	return nil
}

func (m *MenuRepository) EnableMenu(id uint) error {
	if err := m.DB.Model(&model.Menu{}).Where("id = ?", id).Update("suspended", false).Error; err != nil {
		return err
	}
	return nil
}
