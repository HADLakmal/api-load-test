package container

import (
	"github.com/pickme-go/metrics"
	"github.com/tryfix/log"
)

// Container holds all resolved dependencies that needs to be injected at run time.
type Container struct {
	Adapters Adapters
}

// Adapters hold resolved adapter instances.
// These are wrappers around third party libraries. All adapters will be of a corrosponding adapter interface type.
type Adapters struct {
	Log             log.PrefixedLogger
	MetricsReporter metrics.Reporter
}

// Services hold resolved service instances.
// These are abstractions to third party APIs. All services will be of a corrosponding service interface type.
/*type Services struct {
	TestSvc services.RouteServiceInterface
}*/
