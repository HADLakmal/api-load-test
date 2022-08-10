package error

import (
	"context"
	"github.com/HADLakmal/api-load-test/internal/errors"
	"github.com/HADLakmal/api-load-test/internal/http/response"
	"github.com/tryfix/log"
	"net/http"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, w http.ResponseWriter, logger log.PrefixedLogger) {

	e, ok := err.(*errors.Error)

	if ok {
		switch e.Type {
		case errors.SERVER_ERROR, errors.ADAPTER_ERROR, errors.DATA_ERROR, errors.SERVICE_ERROR:
			log.ErrorContext(ctx, "http.error.handler.Handle", "Internal Server Error", err)
			response.Send(w, format(err), http.StatusInternalServerError)
			break
		case errors.MIDDLEWARE_ERROR:
			log.ErrorContext(ctx, "http.error.handler.Handle", "Invalid Request", err)
			response.Send(w, format(err), http.StatusInternalServerError)
			break
		case errors.DOMAIN_ERROR:
			log.ErrorContext(ctx, "http.error.handler.Handle", "Domain Constraint Violation", err)
			response.Send(w, format(err), http.StatusBadRequest)
			break
		case errors.VALIDATION_ERROR:
			log.ErrorContext(ctx, "http.error.handler.Handle", "Validation Structure Error", err)
			response.Send(w, format(err), http.StatusUnprocessableEntity)
			break
		default:
			log.ErrorContext(ctx, "http.error.handler.Handle", "Unknown Error", err)
			response.Send(w, format(err), http.StatusInternalServerError)
			break
		}
	} else {
		log.ErrorContext(ctx, "http.error.handler.Handle", "Unknown Error", err)
		response.Send(w, format(err), http.StatusInternalServerError)
	}

	return
}
