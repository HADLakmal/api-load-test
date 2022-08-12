package adaptors

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/HADLakmal/api-load-test/internal/errors"
	"github.com/HADLakmal/api-load-test/internal/http/request/unpackers"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

func Init() Client {
	return Client{}
}

func (a Client) FetchGeneric(payload unpackers.Generic, req *http.Request) error {
	a.HttpClient = &http.Client{}
	r, err := createRequest(payload, req)

	res, err := a.HttpClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted {
		return errors.New(errors.SERVICE_ERROR, 5000, fmt.Sprintf("Unsuccessful response code from service - %d", res.StatusCode), "")
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(errors.SERVICE_ERROR, 5000, "Error reading response body", "")
	}

	return nil
}

func createRequest(payload unpackers.Generic, r *http.Request) (*http.Request, error) {
	body, err := json.Marshal(payload.Payload)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	req, err := http.NewRequestWithContext(context.Background(), payload.Method, payload.BaseURL+payload.Path, reader)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	for m, v := range r.URL.Query() {
		query.Set(m, v[0])
	}
	req.URL.RawQuery = query.Encode()
	for m, v := range r.Header {
		req.Header.Add(m, v[0])
	}
	return req, nil
}
