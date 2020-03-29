package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lbrulet/Go-api-test/api/handler"
	"github.com/lbrulet/Go-api-test/pkg/user"
)

type HttpService struct {
	router              *gin.Engine
	userService         *user.Service
	userEndpointService *handler.UserEndpointService
}

func (h *HttpService) Router() *gin.Engine {
	return h.router
}

func (h *HttpService) Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}

func NewHttpService(router *gin.Engine, userService *user.Service, userEndpointService *handler.UserEndpointService) *HttpService {
	return &HttpService{
		router:              router,
		userService:         userService,
		userEndpointService: userEndpointService,
	}
}

func (h *HttpService) SetupRouter() {
	h.userService.Migrate()

	h.router.GET("/ping", h.Ping())
	h.router.POST("/user", h.userEndpointService.CreateUser(h.userService))
	h.router.PUT("/user", h.userEndpointService.UpdateUser(h.userService))
	h.router.GET("/users", h.userEndpointService.GetUsers(h.userService))
	h.router.DELETE("/user/:id", h.userEndpointService.DeleteUser(h.userService))
}
