package server

import (
	"context"
	"fmt"
	"github.com/HADLakmal/api-load-test/internal/http/router"
	"github.com/HADLakmal/api-load-test/internal/util/config"
	"github.com/HADLakmal/api-load-test/internal/util/container"
	"github.com/tryfix/log"
	"net/http"
	"strconv"
	"time"
)

// HTTPServer implementes the base type for Http server
type HTTPServer struct {
	server    *http.Server
	logger    log.PrefixedLogger
	config    *config.Config
	container *container.Container
}

// NewHTTPServer creates a new HttpServer
func NewHTTPServer(config *config.Config, container *container.Container) *HTTPServer {
	return &HTTPServer{
		server:    nil,
		logger:    container.Adapters.Log,
		config:    config,
		container: container,
	}
}

// Init initializes the server
func (srv *HTTPServer) Init() error {

	// initialize the router
	r := router.Init(srv.container)

	address := srv.config.AppConfig.Host + ":" + strconv.Itoa(srv.config.AppConfig.Port)

	server := &http.Server{
		Addr: address,

		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,

		// pass our instance of gorilla/mux in
		Handler: r,
	}

	// run our server in a goroutine so that it doesn't block
	go func() {

		err := server.ListenAndServe()
		if err != nil {
			srv.logger.Error("http.server.Init", err)
		}
	}()

	srv.server = server
	srv.logger.Info("http.server.Init", fmt.Sprintf("HTTP server listening on %s", address))

	return nil

}

// ShutDown releases all http connections gracefully and shut down the server
func (srv *HTTPServer) ShutDown(ctx context.Context) {

	go func() {

		srv.logger.Warn("http.server.ShutDown", "Stopping HTTP Server")
		srv.server.SetKeepAlivesEnabled(false)

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		err := srv.server.Shutdown(ctx)
		if err != nil {
			srv.logger.Error("http.server.ShutDown", "Unable to stop HTTP server")

		}
	}()
}
