package api

import (
	"TestTask/app/store/pgstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Login      string `json:"login"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	RoleID     int    `json:"role_id"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param user body createUserRequest true "Create user"
// @Success 201
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Router /user [post]
func (h handler) createUser(c *gin.Context) {
	// Инициализируем запрос.
	req := createUserRequest{}

	// Привязываем JSON к запросу; ошибка — 400 Bad Request.
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Создаём пользователя в базе данных; ошибка — 404 Not Found.
	if err := h.DB.User().Create(&models.User{
		Login:      req.Login,
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		RoleID:     req.RoleID,
	}); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Успешное создание — 201 Created.
	c.Status(http.StatusCreated)
}
