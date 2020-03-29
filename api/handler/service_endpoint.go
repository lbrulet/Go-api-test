package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lbrulet/Go-api-test/pkg/models"
	"github.com/lbrulet/Go-api-test/pkg/user"
)

//UserEndpointService is a user endpoint service structure
type UserEndpointService struct {
	userService *user.Service
}

//NewUserEndpointService is the constructor of UserEndpointService
func NewUserEndpointService(userService *user.Service) *UserEndpointService {
	return &UserEndpointService{userService: userService}
}

//CreateUser is an endpoint to create a user
func (e *UserEndpointService) CreateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		payload := &models.User{}

		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		if err := e.userService.InsertUser(payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": payload})
	}
}

//GetUsers is an endpoint to get all users
func (e *UserEndpointService) GetUsers() func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := e.userService.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

//UpdateUser is an endpoint to update a user
func (e *UserEndpointService) UpdateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		payload := &models.User{}

		if err := c.ShouldBind(payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		_, err := e.userService.GetUserByID(payload.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		_ = e.userService.UpdateUserByID(payload)

		c.JSON(http.StatusOK, gin.H{"data": payload})
	}
}

//DeleteUser is an endpoint to delete a user
func (e *UserEndpointService) DeleteUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		_, err = e.userService.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		_ = e.userService.DeleteUserByID(userID)

		c.JSON(http.StatusNoContent, gin.H{"data": nil})
	}
}
