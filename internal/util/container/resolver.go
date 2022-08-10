package container

import "github.com/HADLakmal/api-load-test/internal/util/config"

// Resolve resolves the entire container.
// The order of resolution is very important. Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
// 		- Adapters
// 		- Repositories
// 		- Services
func Resolve(cfg *config.Config) *Container {

	return &Container{
		Adapters: resolveAdapters(cfg),
		//Services:     resolveServices(cfg),
	}
}
