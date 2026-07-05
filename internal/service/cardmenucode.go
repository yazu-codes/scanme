package service

import (
	"github.com/google/uuid"
	"github.com/yazu-codes/scanme.git/internal/model"
	"github.com/yazu-codes/scanme.git/internal/repository"
	"gorm.io/gorm"
)

type CardMenuCodeService struct {
	DB                     *gorm.DB
	CardMenuCodeRepository *repository.CardMenuCodeRepository
}

func NewCardMenuCodeService(db *gorm.DB) *CardMenuCodeService {
	return &CardMenuCodeService{
		DB:                     db,
		CardMenuCodeRepository: &repository.CardMenuCodeRepository{DB: db},
	}
}

func (s *CardMenuCodeService) GetMenuIdByCode(code string) (int64, error) {
	cardMenuCode, err := s.CardMenuCodeRepository.GetCardMenuCodeByCode(code)
	if err != nil {
		return 0, err
	}
	return cardMenuCode.MenuID, nil
}

func (s *CardMenuCodeService) GetCardMenuCodeByMenuId(menuId int64) (*model.CardMenuCode, error) {
	cardMenuCode, err := s.CardMenuCodeRepository.GetCardMenuCodeByMenuId(menuId)
	if err != nil {
		return nil, err
	}
	return cardMenuCode, nil
}

func (s *CardMenuCodeService) CreateCardMenuCode() (*model.CardMenuCode, error) {
	cardMenuCode := &model.CardMenuCode{}
	cardMenuCode.Code = uuid.New().String()
	createdCardMenuCode, err := s.CardMenuCodeRepository.CreateCardMenuCode(cardMenuCode)
	if err != nil {
		return nil, err
	}
	return createdCardMenuCode, nil
}

func (s *CardMenuCodeService) UpdateCardMenuCode(cardMenuCode *model.CardMenuCode) (*model.CardMenuCode, error) {
	updatedCardMenuCode, err := s.CardMenuCodeRepository.UpdateCardMenuCode(cardMenuCode)
	if err != nil {
		return nil, err
	}
	return updatedCardMenuCode, nil
}

func (s *CardMenuCodeService) DeleteCardMenuCode(id int64) error {
	err := s.CardMenuCodeRepository.DeleteCardMenuCode(id)
	if err != nil {
		return err
	}
	return nil
}
