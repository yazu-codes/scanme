package service

import (
	"fmt"

	"github.com/yazu-codes/scanme.git/internal/dto"
	"github.com/yazu-codes/scanme.git/internal/model"
	"github.com/yazu-codes/scanme.git/internal/repository"
	"gorm.io/gorm"
)

type MenuService struct {
	MenuRepository       *repository.MenuRepository
	MenuOwnerRepository  *repository.MenuOwnerRepository
	MenuConfigRepository *repository.MenuConfigRepository
	MenuItemRepository   *repository.MenuItemRepository
}

func (s *MenuService) AddMenuOwner(owner *model.MenuOwner) (*model.MenuOwner, error) {
	// s.MenuRepository.AddMenuOwner(owner)
	return nil, nil
}

func (s *MenuService) GetMenuNameById(id uint) (string, error) {
	fmt.Println("Getting menu name by ID:", id)
	menu, err := s.MenuRepository.GetMenuByID(id)
	if err != nil {
		return "", err
	}
	fmt.Println("Found menu:", menu)
	return menu.MenuOwner.UrlName, nil
}

func (s *MenuService) AddMenuConfiguration(configuration *model.MenuConfiguration) (*model.MenuConfiguration, error) {
	return nil, nil
}

func (s *MenuService) AddMenuItems(item *[]model.MenuItem) (*[]model.MenuItem, error) {
	return nil, nil
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{
		MenuRepository:       &repository.MenuRepository{DB: db},
		MenuOwnerRepository:  &repository.MenuOwnerRepository{DB: db},
		MenuConfigRepository: &repository.MenuConfigRepository{DB: db},
		MenuItemRepository:   &repository.MenuItemRepository{DB: db},
	}
}

func (s *MenuService) GetAllMenus() ([]model.Menu, error) {
	return s.MenuRepository.GetAllMenus()
}

func (s *MenuService) GetMenuByName(name string) (*model.Menu, error) {
	return s.MenuRepository.GetMenuByName(name)
}

func (s *MenuService) GetMenuByUrlName(urlName string) (*dto.PublicMenu, error) {
	return s.MenuRepository.GetMenuByUrlName(urlName)
}

func (s *MenuService) CreateMenu(menu *model.Menu) error {
	// _, err := s.MenuConfigRepository.CreateMenuConfiguration(&menu.MenuConfiguration)
	// if err != nil {
	// 	return err
	// }
	// _, err = s.MenuOwnerRepository.CreateMenuOwner(&menu.MenuOwner)
	// if err != nil {
	// 	return err
	// }

	// menu.MenuConfiguration = *config
	// menu.MenuOwner = *owner

	// tempItems := []model.MenuItem{}

	// for _, item := range menu.MenuItems {

	// 	_, err := s.MenuItemRepository.CreateMenuItem(&item)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// tempItems = append(tempItems, *item)
	// }

	// menu.MenuItems = tempItems

	return s.MenuRepository.CreateMenu(menu)
}

func (s *MenuService) UpdateMenu(menu *model.Menu) error {
	return s.MenuRepository.UpdateMenu(menu)
}

func (s *MenuService) DeleteMenu(id uint) error {
	return s.MenuRepository.DeleteMenu(id)
}

func (s *MenuService) SuspendMenu(id uint) error {
	return s.MenuRepository.SuspendMenu(id)
}

func (s *MenuService) EnableMenu(id uint) error {
	return s.MenuRepository.EnableMenu(id)
}
