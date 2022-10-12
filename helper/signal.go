package helper

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

//	AwaitSignal
//	скопипизжено: github.com/vivid-money/article-golang-di
func AwaitSignal(ctx context.Context) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sig:
	}
}
