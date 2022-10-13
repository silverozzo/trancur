package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"trancur/domain/model"
)

type Course struct {
	srv CourseService
	cfg Config
}

func NewCourse(srv CourseService, cfg Config) *Course {
	ctrl := &Course{
		srv: srv,
		cfg: cfg,
	}

	return ctrl
}

func (ctrl *Course) GetAllCursBySource(gctx *gin.Context) {
	source := gctx.Param("source")
	if len(source) == 0 {
		source = ctrl.cfg.GetDefaultCourseSource()
	}
	source = strings.ToUpper(source)

	if _, ok := model.SourceMap[source]; !ok {
		gctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": msgNotFound,
		})

		return
	}

	data, err := ctrl.srv.GetCourseListBySource(source)
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
