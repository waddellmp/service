package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/waddellmp/service/foundation/logger"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.New("SALES_API")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// ------------------------------------------------------------------------
	// GOMAXPROCS

	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// ------------------------------------------------------------------------

	shutdown := make(chan os.Signal, 1)

	// SIGTERM is returned from kubernetes pod shutdown request
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our shutdown signal
	<-shutdown

	// Log the shutdown signal
	log.Infow("shutdown", "status", "shutdown started", "signal", shutdown)
	defer log.Infow("shutdown", "status", "shutdown coplete", "signal", shutdown)

	return nil
}
