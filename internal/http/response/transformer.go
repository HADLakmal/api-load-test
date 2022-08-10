package response

import (
	"encoding/json"
)

// Transform transforms a dataset in to a relevent structure and marshal to JSON.
func Transform(data interface{}) []byte {

	wrapper := Data{}
	wrapper.Payload = data

	message, _ := json.Marshal(wrapper)

	return message
}

// TransformLegacy transforms a dataset in to a relevent structure and marshal to JSON.
// TODO This needs to be removed once we get clients to use the new API
func TransformLegacy(data interface{}) []byte {

	message, _ := json.Marshal(data)

	return message
}
