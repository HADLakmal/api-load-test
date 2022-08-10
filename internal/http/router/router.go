package router

import (
	"github.com/HADLakmal/api-load-test/internal/http/controllers"
	"github.com/HADLakmal/api-load-test/internal/http/middleware"
	"github.com/HADLakmal/api-load-test/internal/util/container"
	"github.com/gorilla/mux"
	"net/http"
)

// Init initializes the router.
func Init(container *container.Container) *mux.Router {

	// create new router
	r := mux.NewRouter()

	metricsMidleware := middleware.NewMetricsMiddleware(container)

	r.Use(metricsMidleware.Middleware)

	genericController := controllers.NewGenericController(container)
	r.HandleFunc("/generic", genericController.Execute).Methods(http.MethodPost)

	return r
}
