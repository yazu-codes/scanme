package repository

import (
	"github.com/yazu-codes/scanme.git/internal/model"
	"gorm.io/gorm"
)

type MenuItemRepository struct {
	DB *gorm.DB
}

func (m *MenuItemRepository) GetMenuItemsByMenuId(id uint) ([]model.MenuItem, error) {
	var menuItems []model.MenuItem
	if err := m.DB.Where("menu_id = ?", id).Find(&menuItems).Error; err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (m *MenuItemRepository) CreateMenuItem(item *model.MenuItem) (*model.MenuItem, error) {
	if err := m.DB.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (m *MenuItemRepository) UpdateMenuItem(item *model.MenuItem) (*model.MenuItem, error) {
	if err := m.DB.Save(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (m *MenuItemRepository) DeleteMenuItem(id uint) error {
	if err := m.DB.Delete(&model.MenuItem{}, id).Error; err != nil {
		return err
	}
	return nil
}
