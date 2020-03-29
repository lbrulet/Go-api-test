package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lbrulet/Go-api-test/api/handler"
)

//HttpService is the http server structure
type HttpService struct {
	router              *gin.Engine
	userEndpointService *handler.UserEndpointService
}

//Router is a function that return the router
func (h *HttpService) Router() *gin.Engine {
	return h.router
}

//Ping is an endpoint that return pong
func (h *HttpService) Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong"})
	}
}

//NewHttpService is the constructor of HttpService
func NewHttpService(router *gin.Engine, userEndpointService *handler.UserEndpointService) *HttpService {
	return &HttpService{
		router:              router,
		userEndpointService: userEndpointService,
	}
}

//SetupRouter is a function that will setup the HttpService router variable
func (h *HttpService) SetupRouter() {
	h.router.GET("/ping", h.Ping())

	h.router.POST("/user", h.userEndpointService.CreateUser())
	h.router.PUT("/user", h.userEndpointService.UpdateUser())
	h.router.GET("/users", h.userEndpointService.GetUsers())
	h.router.DELETE("/user/:id", h.userEndpointService.DeleteUser())
}
