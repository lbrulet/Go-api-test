package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/Go-api-test/pkg/models"
	"github.com/lbrulet/Go-api-test/pkg/user"
	"net/http"
	"strconv"
)

type UserEndpointService struct {
	
}

func NewUserEndpointService() *UserEndpointService {
	return &UserEndpointService{}
}

func (e *UserEndpointService) CreateUser(userService *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		payload := &models.User{}

		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		if err := userService.InsertUser(payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": payload})
	}
}

func (e *UserEndpointService) GetUsers(userService *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func (e *UserEndpointService) UpdateUser(userService *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		payload := &models.User{}

		if err := c.ShouldBind(payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		_, err := userService.GetUserByID(payload.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}

		_ = userService.UpdateUserByID(payload)

		c.JSON(http.StatusOK, gin.H{"data": payload})
	}
}

func (e *UserEndpointService) DeleteUser(userService *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		_, err = userService.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		_ = userService.DeleteUserByID(userID)


		c.JSON(http.StatusNoContent, gin.H{"data": nil})
	}
}
