package http

import (
	"log"

	"github.com/gin-gonic/gin"

	"trancur/http/controller"
)

func NewRouter(srv controller.CourseService, cfg Config, infoLog *log.Logger) *gin.Engine {
	crCtrl := controller.NewCourse(srv, cfg)

	rtr := gin.Default()
	rtr.GET("/courses", crCtrl.GetAllCursBySource)
	rtr.GET("/courses/:source", crCtrl.GetAllCursBySource)

	trCtrl := controller.NewTransit(srv, cfg, infoLog)

	rtr.POST("/transit/:source", trCtrl.Transit)

	return rtr
}
