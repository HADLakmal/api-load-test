package response

import "net/http"

// Send sets all required fields and write the response.
func Send(w http.ResponseWriter, payload []byte, code int) {

	// set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// set response code
	w.WriteHeader(code)

	// set payload
	w.Write(payload)
}
