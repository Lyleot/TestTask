package api

import (
	"TestTask/app/store/pgstore/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetRoleResponse struct {
	ID          int       `json:"role_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetRole godoc
// @Summary Get role information
// @Description Retrieves role information by ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.Role
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Router /role/{id} [get]
func (h handler) getRole(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid role ID"})
		return
	}

	var role *models.Role
	role, err = h.DB.Role().Find(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, buildGetRoleResponse(role))
}

func buildGetRoleResponse(role *models.Role) (result GetRoleResponse) {
	result = GetRoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}

	return result
}
