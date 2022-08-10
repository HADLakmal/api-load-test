package bootstrap

import (
	"fmt"
	"github.com/HADLakmal/api-load-test/internal/util/config"
	"github.com/HADLakmal/api-load-test/internal/util/container"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

const namespace = "ceylonwings_backend"

// Expose metrics as a separate Prometheus metric server.
func exposeMetrics(cfg config.AppConfig, ctr *container.Container) {

	if !cfg.Metrics.Enabled {
		return
	}

	// register mysql collectors
	registerStoreMetrics(ctr)

	// set metric exposing port and endpoint
	address := cfg.Host + ":" + strconv.Itoa(cfg.Metrics.Port)
	http.Handle(cfg.Metrics.Route, promhttp.Handler())

	// run metric server in a goroutine so that it doesn't block
	go func() {

		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Println(err)
			panic("Metric server error...")
		}
	}()

	logger.Info("bootstrap.metrics.exposeMetrics", fmt.Sprintf("Exposing metrics on %v ...", address))
}

func registerStoreMetrics(ctr *container.Container) {
	const subsystem = "store"

	//// The number of established connections both in use and idle.
	//storeSize := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	//	Namespace: namespace,
	//	Subsystem: subsystem,
	//	Name:      "store_size",
	//	Help:      "Number of cells in the store",
	//}, func() float64 {
	//	return float64(ctr.Repositories.Repository.StoreSize())
	//})
	//
	//// The number of established connections both in use and idle.
	//syncedRadiusEntries := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	//	Namespace: namespace,
	//	Subsystem: subsystem,
	//	Name:      "radius_entries",
	//	Help:      "Number of synced radius entries",
	//}, func() float64 {
	//	return float64(ctr.Repositories.RadiusRepository.SyncedItems())
	//})
	//
	//prometheus.MustRegister(storeSize)
	//prometheus.MustRegister(syncedRadiusEntries)

}
