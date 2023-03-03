package server

import (
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	go func() {
		HttpServerRun()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	HttpServerStop()
}
