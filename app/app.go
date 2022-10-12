package app

import (
	"context"

	"trancur/config"
	"trancur/domain/service"
	"trancur/heartbeat"
	"trancur/helper"
	"trancur/http"
	"trancur/log"
)

func Run(ctx context.Context) {
	infoLog := log.NewInfoLog()
	errLog := log.NewErrorLog()

	infoLog.Println("начали")

	cfg := config.NewConfig(errLog)

	crSrv := service.NewCourse()

	//	запускаем обновление от ЦБ РФ
	rusHb := heartbeat.NewRusHeartbeat(crSrv, infoLog, errLog)
	go rusHb.StartBeat(ctx)

	thHb := heartbeat.NewThHeartbeat(crSrv, infoLog, errLog)
	go thHb.StartBeat(ctx)

	//	запускаем http сервер
	rtr := http.NewRouter(crSrv, cfg)
	httpSrv := http.NewServer(rtr, cfg, infoLog, errLog)
	go httpSrv.Run()
	defer httpSrv.Shutdown(ctx)

	helper.AwaitSignal(ctx)

	infoLog.Println("закончили")
}
