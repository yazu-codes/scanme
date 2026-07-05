package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yazu-codes/scanme.git/internal/model"
	"github.com/yazu-codes/scanme.git/internal/service"
)

type PublicHandler struct {
	service     *service.MenuService
	codeService *service.CardMenuCodeService
}

func NewPublicHandler(service *service.MenuService, codeService *service.CardMenuCodeService) *PublicHandler {
	return &PublicHandler{service: service, codeService: codeService}
}

func (h *PublicHandler) GetMenus(c *gin.Context) {
	menus, err := h.service.GetAllMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func (h *PublicHandler) GetMenuByName(c *gin.Context) {
	name := c.Param("name")

	menu, err := h.service.GetMenuByUrlName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func (h *PublicHandler) GetMenuNameByCode(c *gin.Context) {
	code := c.Param("code")
	menuId, err := h.codeService.GetMenuIdByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	menuName, err := h.service.GetMenuNameById(uint(menuId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menuName": menuName})
}

func (h *PublicHandler) UpdateMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateMenu(&menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)
}

func (h *PublicHandler) DeleteMenuById(c *gin.Context) {
	id := c.Param("id")
	idtouint, err := strconv.ParseInt(id, 10, 64)
	properid := uint(idtouint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}
	err = h.service.DeleteMenu(properid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}

func (h *PublicHandler) SuspendMenuById(c *gin.Context) {
	id := c.Param("id")
	idtouint, err := strconv.ParseInt(id, 10, 64)
	properid := uint(idtouint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}
	err = h.service.SuspendMenu(properid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu suspended successfully"})
}

func (h *PublicHandler) EnableMenuById(c *gin.Context) {
	id := c.Param("id")
	idtouint, err := strconv.ParseInt(id, 10, 64)
	properid := uint(idtouint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}
	err = h.service.EnableMenu(properid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu enabled successfully"})
}

func (h *PublicHandler) CreateMenu(c *gin.Context) {
	fmt.Println("AA")
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateMenu(&menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, menu)
}

func (h *PublicHandler) CreateCardMenuCode(c *gin.Context) {
	createdCardMenuCode, err := h.codeService.CreateCardMenuCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdCardMenuCode)
}

func (h *PublicHandler) AddMenuItems(c *gin.Context) {
	var menuItems []model.MenuItem
	if err := c.ShouldBindJSON(&menuItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.service.AddMenuItems(&menuItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, menuItems)
}

func (h *PublicHandler) AddMenuConfiguration(c *gin.Context) {
	var config model.MenuConfiguration
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.service.AddMenuConfiguration(&config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, config)
}

func (h *PublicHandler) AddMenuOwner(c *gin.Context) {
	var owner model.MenuOwner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.service.AddMenuOwner(&owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, owner)
}

func (h *PublicHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome!",
	})
}

func (h *PublicHandler) Login(c *gin.Context) {
	// Normally you'd validate credentials and generate a JWT.
	c.JSON(http.StatusOK, gin.H{
		"token": "my-secret-token",
	})
}
