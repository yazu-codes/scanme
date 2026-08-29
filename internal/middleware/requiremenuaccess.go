package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yazu-codes/scanme.git/internal/utils"
)

func RequireMenuAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			JWT middleware must already have validated the token
			and populated these values.
		*/
		roleValue, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing authentication context",
				},
			)
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid authentication context",
				},
			)
			return
		}

		/*
			Admins may update any menu.
		*/
		if strings.EqualFold(role, "admin") {
			c.Next()
			return
		}

		/*
			Ordinary users must have menu associations.
		*/
		menusValue, exists := c.Get("menus")
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": "no menu associations",
				},
			)
			return
		}

		associatedMenus, ok := menusValue.(string)
		if !ok {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": "invalid menu associations",
				},
			)
			return
		}

		/*
			Read the request body.
		*/
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"error": "could not read request body",
				},
			)
			return
		}

		/*
			Restore body immediately so the next middleware/handler
			can still bind it.
		*/
		c.Request.Body = io.NopCloser(
			bytes.NewBuffer(bodyBytes),
		)

		/*
			We only care about the menu ID here.
			No need to unmarshal the entire menu object.
		*/
		var request struct {
			ID uint `json:"id"`
		}

		if err := json.Unmarshal(
			bodyBytes,
			&request,
		); err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"error": "invalid request body",
				},
			)
			return
		}

		if request.ID == 0 {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"error": "missing menu id",
				},
			)
			return
		}

		/*
			Your JWT stores something like:

			menus: "1,4,7"
		*/
		menuIDs, err := utils.StringToIntArray(
			associatedMenus,
		)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": "invalid menu associations",
				},
			)
			return
		}

		allowed := false

		for _, menuID := range menuIDs {
			fmt.Println("comparing:", menuID, request.ID)
			if uint(menuID) == request.ID {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": fmt.Sprint(associatedMenus, request.ID),
				},
			)
			return
		}

		c.Next()
	}
}
