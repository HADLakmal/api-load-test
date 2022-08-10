package controllers

import (
	"github.com/HADLakmal/api-load-test/internal/domain/adaptors"
	"github.com/HADLakmal/api-load-test/internal/http/request"
	"github.com/HADLakmal/api-load-test/internal/http/request/unpackers"
	"github.com/HADLakmal/api-load-test/internal/http/response"
	"github.com/HADLakmal/api-load-test/internal/util/container"
	"github.com/tryfix/log"
	"net/http"
	"sync"
)

type Generic struct {
	Container *container.Container
}

// NewGenericController returns a base type for this controllers
func NewGenericController(container *container.Container) *Generic {
	return &Generic{
		Container: container,
	}
}

func (ctl *Generic) Execute(w http.ResponseWriter, r *http.Request) {
	a := unpackers.Generic{}
	err := request.Unpack(r, &a)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	var success int
	cl := adaptors.Init()
	wg := sync.WaitGroup{}
	for x := 0; x < int(a.Parallelism); x++ {
		go func() {
			wg.Add(1)
			errGen := cl.FetchGeneric(a, r)
			if errGen != nil {
				log.Error(errGen.Error())
			}
			cl.HttpClient.CloseIdleConnections()
			success++
			wg.Done()
		}()
	}
	wg.Wait()

	response.Send(w, response.Transform(struct {
		Success int
	}{success}), http.StatusOK)
}
