package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"trancur/domain/service"
)

type Controller struct {
	srv *service.Course
}

func NewController(srv *service.Course) *Controller {
	ctrl := &Controller{
		srv: srv,
	}

	return ctrl
}

func (ctrl *Controller) Ping(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (ctrl *Controller) GetAllCursBySource(gctx *gin.Context) {
	data := ctrl.srv.GetCourseListBySource("foobar")

	gctx.JSON(http.StatusOK, data)
}
