package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"trancur/domain/model"
)

var (
	msgNotFound = "не найден указанный источник"
	msgInternal = "зайдите попозже"
)

type CourseService interface {
	GetCourseListBySource(string) (*model.ExchangeList, error)
}

type Controller struct {
	srv CourseService
}

func NewController(srv CourseService) *Controller {
	ctrl := &Controller{
		srv: srv,
	}

	return ctrl
}

func (ctrl *Controller) Ping(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (ctrl *Controller) GetAllCursBySource(gctx *gin.Context) {
	data, err := ctrl.srv.GetCourseListBySource("RUS")
	if errors.Is(err, model.ErrNotFound) {
		gctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": msgNotFound,
		})

		return
	}

	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": msgInternal,
		})

		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"data":  data,
		"error": nil,
	})
}
