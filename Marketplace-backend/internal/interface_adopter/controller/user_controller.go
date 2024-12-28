package controller

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// // UserController represents a user controller
type UserController struct {
	userService service.UserService
}

// NewUserController returns a new instance of UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (Us *UserController) RegisterUser(c *gin.Context) {
	var user entity.User

	// bind JSON data from the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call the service to register the user
	createdUser, err := Us.userService.RegisterUser(user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.RoleName, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the created user
	c.JSON(http.StatusCreated, createdUser)
}

func (Us *UserController) AuthenticateUser(c *gin.Context) {
	var user entity.User

	//  Bind JSON data from the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call the service to authenticate the user
	authenticatedUser, err := Us.userService.AuthenticateUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the authenticated user
	c.JSON(http.StatusOK, authenticatedUser)
}

func (Us *UserController) GetUserById(c *gin.Context) {
	// get the user ID from the request paramaters
	userParam := c.Param("id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call the service to get the user by ID
	user, err := Us.userService.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the user
	c.JSON(http.StatusOK, user)

}

func (Us *UserController) GetAllUsers(c *gin.Context) {
	// call the service to get the user
	user, err := Us.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the user in the response
	c.JSON(http.StatusOK, user)

}

func (Us *UserController) UpdateUser(c *gin.Context) {
	var user entity.User

	// get the user ID from the request parameters
	userParam := c.Param("id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// BIND json from the request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID
	// call the service to update the user
	if err := Us.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the user
	c.JSON(http.StatusOK, user)
}

func (Us *UserController) DeleteUser(c *gin.Context) {
	// get the user ID from the request parameters
	userParam := c.Param("id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call the service to delete the user
	if err := Us.userService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
