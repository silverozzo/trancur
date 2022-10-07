package http

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"trancur/config"
)

type Server struct {
	srv     *http.Server
	infoLog *log.Logger
	errLog  *log.Logger
}

func NewServer(rtr *gin.Engine, cfg *config.Config, infoLog, errLog *log.Logger) *Server {
	srv := &http.Server{
		Addr:    cfg.GetSelfHttpPort(),
		Handler: rtr,
	}

	self := &Server{
		srv:     srv,
		infoLog: infoLog,
		errLog:  errLog,
	}

	return self
}

func (srv *Server) Run() {
	srv.infoLog.Println("поднимаем http сервер")

	if err := srv.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		srv.errLog.Println("еггог:", err)
	}
}

func (srv *Server) Shutdown(ctx context.Context) {
	srv.infoLog.Println("опускаем http сервер")

	if err := srv.srv.Shutdown(ctx); err != nil {
		srv.errLog.Println("еггог:", err)
	}
}
