package repository

import (
	"github.com/yazu-codes/scanme.git/internal/model"
	"gorm.io/gorm"
)

type CardMenuCodeRepository struct {
	DB *gorm.DB
}

func (c *CardMenuCodeRepository) GetCardMenuCodeByMenuId(menuId int64) (*model.CardMenuCode, error) {
	var cardMenuCode model.CardMenuCode
	err := c.DB.Where("menu_id = ?", menuId).First(&cardMenuCode).Error
	if err != nil {
		return nil, err
	}
	return &cardMenuCode, nil
}

func (c *CardMenuCodeRepository) CreateCardMenuCode(cardMenuCode *model.CardMenuCode) (*model.CardMenuCode, error) {
	if err := c.DB.Create(cardMenuCode).Error; err != nil {
		return nil, err
	}
	return cardMenuCode, nil
}

// GetCardMenuCodeByCode retrieves a CardMenuCode by its code.
func (c *CardMenuCodeRepository) GetCardMenuCodeByCode(code string) (*model.CardMenuCode, error) {
	var cardMenuCode model.CardMenuCode
	err := c.DB.Where("code = ?", code).First(&cardMenuCode).Error
	if err != nil {
		return nil, err
	}
	return &cardMenuCode, nil
}

func (c *CardMenuCodeRepository) UpdateCardMenuCode(cardMenuCode *model.CardMenuCode) (*model.CardMenuCode, error) {
	if err := c.DB.Save(cardMenuCode).Error; err != nil {
		return nil, err
	}
	return cardMenuCode, nil
}

func (c *CardMenuCodeRepository) DeleteCardMenuCode(id int64) error {
	if err := c.DB.Delete(&model.CardMenuCode{}, id).Error; err != nil {
		return err
	}
	return nil
}
