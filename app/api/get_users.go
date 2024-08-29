package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetListUserResponse []GetUserResponse

// GetUsers godoc
// @Summary Get all users
// @Description Retrieves a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /user/all [get]
func (h handler) getUsers(c *gin.Context) {
	// Получаем всех пользователей из базы данных; ошибка — 500 Internal Server Error.
	users, err := h.DB.User().FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Формируем ответ с пользователями.
	response := make(GetListUserResponse, 0, len(users))
	for i := range users {
		response = append(response, buildGetUserResponse(&users[i]))
	}

	// Отправляем список пользователей в ответе; статус — 200 OK.
	c.JSON(http.StatusOK, response)
}
