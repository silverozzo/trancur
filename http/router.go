package http

import (
	"github.com/gin-gonic/gin"

	"trancur/http/controller"
)

func NewRouter(srv controller.CourseService, cfg Config) *gin.Engine {
	crCtrl := controller.NewCourse(srv, cfg)

	rtr := gin.Default()
	rtr.GET("/courses", crCtrl.GetAllCursBySource)
	rtr.GET("/courses/:source", crCtrl.GetAllCursBySource)

	return rtr
}
