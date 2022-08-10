package bootstrap

import (
	"context"
	httpServer "github.com/HADLakmal/api-load-test/internal/http/server"
	"github.com/HADLakmal/api-load-test/internal/util/config"
	"github.com/HADLakmal/api-load-test/internal/util/container"
	"github.com/tryfix/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger log.PrefixedLogger

// Start starts the application bootstrap process
func Start() {

	// parse all configurations
	cfg := config.Parse()

	// resolve the container using parsed configurations
	ctr := container.Resolve(cfg)

	logger = ctr.Adapters.Log

	//Initialize HTTP server

	httpSrv := httpServer.NewHTTPServer(cfg, ctr)
	err := httpSrv.Init()
	if err != nil {
		panic(err)
	}

	// expose application metrics
	exposeMetrics(cfg.AppConfig, ctr)

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT)

	// Block until we receive our signal
	<-c

	// Shutdown Http server
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	httpSrv.ShutDown(ctx)

	<-ctx.Done()

	// Destruct other respouces and stop the service
	Destruct(ctr)

	logger.Warn("bootstrap.init.Start", "Service shutdown gracefully...")
	os.Exit(0)

}

// Destruct gracefully close all additional resources.
func Destruct(ctr *container.Container) {
	logger.Info("bootstrap.Destruct", "Closing database connections...")
}
