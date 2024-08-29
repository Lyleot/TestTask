package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteRole godoc
// @Summary Delete a role
// @Description Deletes a role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /role/{id} [delete]
func (h handler) deleteRole(c *gin.Context) {
	// Получаем ID роли из параметра запроса.
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid role ID"}) // Ошибка — 400 Bad Request.
		return
	}

	// Удаляем роль из базы данных; ошибка — 500 Internal Server Error.
	if err = h.DB.Role().Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Успешное удаление — 200 OK.
	c.Status(http.StatusOK)
}
