package test

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Get is a helper that makes a HTTP GET request against an endpoint and
// returns an unmarshaled response. Takes the URL, data source string, and
// container for the unmarshaled output as parameters.
func Get(url string, dataSource string, out interface{}) (*http.Response, error) {
	return makeRequest("GET", url, dataSource, nil, out)
}

// Post is a helper that makes a HTTP POST request against an endpoint and
// returns an unmarshaled response. Takes the URL, JSON payload, data source
// string, and container for the unmarshaled output as parameters.
func Post(url string, payload interface{}, dataSource string, out interface{}) (*http.Response, error) {
	return makeRequest("POST", url, dataSource, payload, out)
}

func makeRequest(method string, url string, dataSource string, payload interface{}, out interface{}) (*http.Response, error) {
	jsonContent, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	payloadBytes := bytes.NewBuffer(jsonContent)
	req, err := http.NewRequest(method, url, payloadBytes)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("FC-Data-Source", dataSource)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(out)
	return res, err
}
