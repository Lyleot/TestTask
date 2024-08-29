package api

import (
	"TestTask/app/store/pgstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRole godoc
// @Summary Create a new role
// @Description Creates a new role with the given details
// @Tags roles
// @Accept json
// @Produce json
// @Param role body createRoleRequest true "Create role"
// @Success 201
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /role [post]
func (h handler) createRole(c *gin.Context) {
	req := createRoleRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.DB.Role().Create(&models.Role{
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusCreated)
}
