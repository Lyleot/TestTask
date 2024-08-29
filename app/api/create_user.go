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
	req := createUserRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.DB.User().Create(&models.User{
		Login:      req.Login,
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		RoleID:     req.RoleID,
	}); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusCreated)
}
