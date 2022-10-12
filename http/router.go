package http

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(srv CourseService) *gin.Engine {
	ctrl := NewController(srv)

	rtr := gin.Default()
	rtr.GET("/ping", ctrl.Ping)
	rtr.GET("/courses", ctrl.GetAllCursBySource)

	return rtr
}
