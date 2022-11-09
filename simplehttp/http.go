package simplehttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HttpRequestParams struct {
	QueryParams map[string]string
	Headers     map[string]string
}

type httpResponse struct {
	statusCode int
	headers    http.Header
	body       []byte
}

func call(client *http.Client, url string, method string, param HttpRequestParams, requestBody []byte) (httpResponse, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return httpResponse{}, err
	}
	// Set query paramateres
	if param.QueryParams != nil {
		q := req.URL.Query()
		for k, v := range param.QueryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	// Set headers
	if param.Headers != nil {
		for k, v := range param.Headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return httpResponse{}, err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse{}, err
	}
	return httpResponse{body: responseBody, statusCode: resp.StatusCode, headers: resp.Header}, nil
}

func GetJson[T any](client *http.Client, url string, param HttpRequestParams) (*T, error) {
	response, err := call(client, url, "GET", param, nil)
	if err != nil {
		return nil, err
	}
	// I return a JSON serialization
	data := new(T)
	if err := json.Unmarshal(response.body, &data); err != nil {
		return nil, err
	}
	return data, nil
}
