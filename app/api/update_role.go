package api

import (
	"TestTask/app/store/pgstore/models"
	"TestTask/app/utils/ptr"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateRoleRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateRole godoc
// @Summary Update role information
// @Description Updates the details of a role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body updateRoleRequest true "Update role"
// @Success 200
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /role/{id} [patch]
func (h handler) updateRole(c *gin.Context) {
	// Получаем ID роли из параметров запроса.
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid role ID"}) // Некорректный ID роли
		return
	}

	// Ищем роль в базе данных.
	var role *models.Role
	role, err = h.DB.Role().Find(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) // Ошибка при поиске роли
		return
	}

	// Привязываем данные из запроса к структуре.
	req := updateRoleRequest{}
	if err = c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err) // Некорректные данные запроса
		return
	}

	// Обновляем поля роли.
	updateRoleFields(role, req)

	// Сохраняем обновленную роль в базе данных.
	if err = h.DB.Role().Save(role); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) // Ошибка при сохранении роли
		return
	}

	// Возвращаем статус 200 OK.
	c.Status(http.StatusOK)
}

// updateRoleFields обновляет поля роли в зависимости от данных запроса.
func updateRoleFields(role *models.Role, params updateRoleRequest) {
	if params.Name != nil {
		role.Name = ptr.DeRef(params.Name)
	}

	if params.Description != nil {
		role.Description = ptr.DeRef(params.Description)
	}
}
