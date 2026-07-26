package repository

import (
	"github.com/yazu-codes/scanme.git/internal/model"
	"gorm.io/gorm"
)

type MenuOwnerRepository struct {
	DB *gorm.DB
}

func (m *MenuOwnerRepository) GetMenuOwnerByMenuId(id uint) (model.MenuOwner, error) {
	var menuOwner model.MenuOwner
	if err := m.DB.First("menu_id = ?", id).Find(&menuOwner).Error; err != nil {
		return model.MenuOwner{}, err
	}
	return menuOwner, nil
}

func (m *MenuOwnerRepository) CreateMenuOwner(menuOwner *model.MenuOwner) (*model.MenuOwner, error) {
	if err := m.DB.Create(menuOwner).Error; err != nil {
		return nil, err
	}
	return menuOwner, nil
}

func (m *MenuOwnerRepository) UpdateMenuOwner(menuOwner *model.MenuOwner) (*model.MenuOwner, error) {
	if err := m.DB.Save(menuOwner).Error; err != nil {
		return nil, err
	}
	return menuOwner, nil
}

func (m *MenuOwnerRepository) DeleteMenuOwner(id uint) error {
	if err := m.DB.Delete(&model.MenuOwner{}, id).Error; err != nil {
		return err
	}
	return nil
}
