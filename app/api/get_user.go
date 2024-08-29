package api

import (
	"TestTask/app/store/pgstore/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	ID         int       `json:"user_id"`
	Login      string    `json:"login"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	RoleID     int       `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// GetUser godoc
// @Summary Get user information
// @Description Retrieves user information by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Router /user/{id} [get]
func (h handler) getUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var user *models.User
	user, err = h.DB.User().Find(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, buildGetUserResponse(user))
}

func buildGetUserResponse(user *models.User) (result GetUserResponse) {
	result = GetUserResponse{
		ID:         user.ID,
		Login:      user.Login,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		RoleID:     user.RoleID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return result
}
