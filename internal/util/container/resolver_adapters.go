package container

import (
	"github.com/HADLakmal/api-load-test/internal/util/config"
	"github.com/pickme-go/metrics"
	"github.com/tryfix/log"
)

var resolvedAdapters Adapters

func resolveAdapters(cfg *config.Config) Adapters {
	resolveLogAdapter(cfg.LogConfig)
	resolveMetricsAdapter()
	//resolveDBAdapter(cfg.DatabaseConfig)

	return resolvedAdapters
}

func resolveLogAdapter(cfg config.LogConfig) {

	logger := log.Constructor.PrefixedLog(
		log.WithLevel(cfg.Level),
		log.WithColors(cfg.Colors),
		log.WithFilePath(cfg.FilePathEnabled),
	)
	resolvedAdapters.Log = logger
}

func resolveMetricsAdapter() {
	reporter := metrics.PrometheusReporter(`ceylon_wings`, `backend`)
	resolvedAdapters.MetricsReporter = reporter
}

/*
func resolveMetricsAdapter() {
	reporter := adapters.NewMetricsAdapter(`geo`, `gis_api`)
	resolvedAdapters.MetricsReporter = reporter
}*/

// Resolve the database adapter.
//func resolveDBAdapter(cfg config.DatabaseConfig) {
//
//	pg := mysql.Adapter{}
//	db, _ := pg.New(cfg, resolvedAdapters.MetricsReporter)
//	resolvedAdapters.MySQL = db
//}

// Resolve the database adapter.
/*func resolveRadiusStoreAdapter(cfg config.DatabaseConfig) {
	mySQLAdapter, _ = mysql.NewAdapter()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Database.User, config.Database.Pass, config.Database.Hostname, config.Database.Port, config.Database.Db)

	err = mySQLAdapter.Connect("mysql", dsn, config.Database.MaxOpenConns)
	pg := mysql.Adapter{}
	db, _ := pg.New(cfg, resolvedAdapters.MetricsReporter)
	resolvedAdapters.RadiusStoreDBAdapter = db
}
*/
