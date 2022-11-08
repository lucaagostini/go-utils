package simplehttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HttpParams struct {
	queryParams map[string]string
	headers     map[string]string
}

func Get[T any](client *http.Client, url string, param HttpParams) (*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Set query paramateres
	if param.queryParams != nil {
		q := req.URL.Query()
		for k, v := range param.queryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	// Set headers
	if param.headers != nil {
		for k, v := range param.headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(T)
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, nil
	}
	return data, nil
}
