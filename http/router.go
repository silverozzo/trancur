package http

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(srv CourseService, cfg Config) *gin.Engine {
	ctrl := NewController(srv, cfg)

	rtr := gin.Default()
	rtr.GET("/ping", ctrl.Ping)
	rtr.GET("/courses", ctrl.GetAllCursBySource)
	rtr.GET("/courses/:source", ctrl.GetAllCursBySource)

	return rtr
}
