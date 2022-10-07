package http

import (
	"github.com/gin-gonic/gin"

	"trancur/domain/service"
)

func NewRouter(srv *service.Course) *gin.Engine {
	ctrl := NewController(srv)

	rtr := gin.Default()
	rtr.GET("/ping", ctrl.Ping)
	rtr.GET("/courses", ctrl.GetAllCursBySource)

	return rtr
}
