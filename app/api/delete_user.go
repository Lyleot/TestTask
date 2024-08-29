package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /user/{id} [delete]
func (h handler) deleteUser(c *gin.Context) {
	// Получаем ID пользователя из параметра запроса.
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"}) // Ошибка — 400 Bad Request.
		return
	}

	// Удаляем пользователя из базы данных; ошибка — 500 Internal Server Error.
	if err = h.DB.User().Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Успешное удаление — 200 OK.
	c.Status(http.StatusOK)
}
