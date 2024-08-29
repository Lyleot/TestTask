package api

import (
	"TestTask/app/store/pgstore/models"
	"TestTask/app/utils/ptr"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateUserRequest struct {
	Login      *string `json:"login,omitempty"`
	FirstName  *string `json:"first_name,omitempty"`
	SecondName *string `json:"second_name,omitempty"`
	RoleID     *int    `json:"role_id,omitempty"`
}

// UpdateUser godoc
// @Summary Update user information
// @Description Updates the details of a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body updateUserRequest true "Update user"
// @Success 200
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /user/{id} [patch]
func (h handler) updateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var user *models.User
	user, err = h.DB.User().Find(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := updateUserRequest{}

	if err = c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateUserFields(user, req)

	if err = h.DB.User().Save(user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func updateUserFields(user *models.User, params updateUserRequest) {
	if params.Login != nil {
		user.Login = ptr.DeRef(params.Login)
	}

	if params.FirstName != nil {
		user.FirstName = ptr.DeRef(params.FirstName)
	}

	if params.SecondName != nil {
		user.SecondName = ptr.DeRef(params.SecondName)
	}

	if params.RoleID != nil {
		user.RoleID = ptr.DeRef(params.RoleID)
	}
}
