package request

import (
	"encoding/json"
	"fmt"
	"github.com/HADLakmal/api-load-test/internal/errors"
	"net/http"
)

// Unpack the request in to the given unpacker struct.
func Unpack(r *http.Request, unpacker UnpackerInterface) error {

	err := json.NewDecoder(r.Body).Decode(unpacker)

	if err != nil {
		verr := errors.New(errors.VALIDATION_ERROR, 5000, fmt.Sprintf("Require following format: %s", unpacker.RequiredFormat()), "")

		return verr
	}

	return nil
}
