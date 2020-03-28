package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/Go-api-test/pkg/user"
)

type HttpService struct {
	router *gin.Engine
	userService *user.Service
}

func (h *HttpService) Router() *gin.Engine {
	return h.router
}

func NewHttpService(router *gin.Engine, userService *user.Service) *HttpService {
	return &HttpService{router: router, userService: userService}
}

func (h *HttpService) SetupRouter() {
	h.router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

