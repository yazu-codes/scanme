package repository

import (
	"github.com/yazu-codes/scanme.git/internal/model"
	"gorm.io/gorm"
)

type MenuConfigRepository struct {
	DB *gorm.DB
}

func (m *MenuConfigRepository) GetMenuConfigByID(id uint) (*model.MenuConfiguration, error) {
	var config model.MenuConfiguration
	if err := m.DB.First(&config, id).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (m *MenuConfigRepository) CreateMenuConfiguration(configuration *model.MenuConfiguration) (*model.MenuConfiguration, error) {
	if err := m.DB.Create(configuration).Error; err != nil {
		return nil, err
	}
	return configuration, nil
}

func (m *MenuConfigRepository) UpdateMenuConfiguration(configuration *model.MenuConfiguration) (*model.MenuConfiguration, error) {
	if err := m.DB.Save(configuration).Error; err != nil {
		return nil, err
	}
	return configuration, nil
}

func (m *MenuConfigRepository) DeleteMenuConfiguration(id uint) error {
	if err := m.DB.Delete(&model.MenuConfiguration{}, id).Error; err != nil {
		return err
	}
	return nil
}
