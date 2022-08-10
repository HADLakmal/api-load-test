package error

import (
	"encoding/json"
	"github.com/HADLakmal/api-load-test/internal/errors"
	"github.com/HADLakmal/api-load-test/internal/http/response"
	"github.com/HADLakmal/api-load-test/internal/http/response/transformers"
	"github.com/iancoleman/strcase"
	"strings"
)

// Format error details.
func format(err error) []byte {

	wrapper := response.Error{}
	var payload interface{}

	e, ok := err.(*errors.Error)

	if ok {
		switch e.Type {
		case errors.SERVER_ERROR, errors.MIDDLEWARE_ERROR, errors.ADAPTER_ERROR, errors.DATA_ERROR, errors.SERVICE_ERROR, errors.DOMAIN_ERROR:
			payload = formatCustomError(*e)
			break
		case errors.VALIDATION_ERROR:
			payload = formatValidationStructureError(err)
			break
		default:
			payload = formatUnknownError(err)
			break
		}
	} else {
		payload = formatUnknownError(err)
	}

	wrapper.Payload = payload
	message, _ := json.Marshal(wrapper)

	return message
}

// Format custom errors.
func formatCustomError(err errors.Error) transformers.ErrorTransformer {

	payload := transformers.ErrorTransformer{}
	payload.Msg = err.Msg
	payload.Code = err.Code
	payload.Type = int(err.Type)
	payload.Trace = err.Details

	return payload
}

// Format validation structure errors.
// These occur when the format of the sent data structure does not match the expected format.
func formatValidationStructureError(err error) transformers.ValidationErrorTransformer {

	payload := transformers.ValidationErrorTransformer{}
	payload.Type = "Validation Errors"
	payload.Trace = err.Error()

	return payload
}

// Format validation errors.
// These are errors thrown when field wise validations happen against a data structure.
func formatValidationErrors(p map[string]string) []byte {

	wrapper := response.Error{}

	payload := transformers.ValidationErrorTransformer{}
	payload.Type = "Validation Errors"
	payload.Trace = formatValidationPayload(p)

	wrapper.Payload = payload
	message, _ := json.Marshal(wrapper)

	return message
}

// Do a final round of formatting to validation errors.
func formatValidationPayload(p map[string]string) map[string]string {

	ep := make(map[string]string)

	for k, v := range p {
		strcase.ToSnake(strings.Join(strings.Split(k, ".")[1:], "."))
		ep[k] = v
	}

	return ep
}

// Format errors of unhandled types.
func formatUnknownError(err error) transformers.ErrorTransformer {

	payload := transformers.ErrorTransformer{}
	payload.Type = int(errors.UNKNOWN_ERROR)
	payload.Msg = err.Error()

	return payload
}
