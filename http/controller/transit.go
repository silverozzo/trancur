package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"trancur/domain/model"
)

type Transit struct {
	srv     CourseService
	cfg     Config
	infoLog *log.Logger
}

type TransitForm struct {
	InputCurrency  string  `json:"input"  binding:"required"`
	OutputCurrency string  `json:"output" binding:"required"`
	Value          float64 `json:"value"  binding:"required"`
}

func NewTransit(srv CourseService, cfg Config, infoLog *log.Logger) *Transit {
	ctrl := &Transit{
		srv:     srv,
		infoLog: infoLog,
	}

	return ctrl
}

func (ctrl *Transit) Transit(gctx *gin.Context) {
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

	var frm TransitForm

	if err := gctx.ShouldBindJSON(&frm); err != nil {
		gctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": msgBadRequest,
		})

		return
	}

	res, err := ctrl.srv.Transit(source, frm.InputCurrency, frm.Value, frm.OutputCurrency)
	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": msgInternal,
		})

		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"input":  frm.InputCurrency,
			"output": frm.OutputCurrency,
			"value":  frm.Value,
			"result": res,
		},
		"error": nil,
	})
}
