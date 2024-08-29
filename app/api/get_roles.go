package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetListRoleResponse []GetRoleResponse

// GetRoles godoc
// @Summary Get all roles
// @Description Retrieves a list of all roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {array} models.Role
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /role/all [get]
func (h handler) getRoles(c *gin.Context) {
	// Получаем все роли из базы данных; ошибка — 500 Internal Server Error.
	roles, err := h.DB.Role().FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Формируем ответ с ролями.
	response := make(GetListRoleResponse, 0, len(roles))
	for i := range roles {
		response = append(response, buildGetRoleResponse(&roles[i]))
	}

	// Отправляем список ролей в ответе; статус — 200 OK.
	c.JSON(http.StatusOK, response)
}
