/*
@Time : 2023/3/31 13:30
@Author : sc-52766
@File : http_client.go
@Software: GoLand
*/
package api




// We can use the net/http package to implement an HTTP client with the desired features.
// Here is an example implementation:

import (
    "bytes"
    "encoding/json"
    "net/http"
    "time"
)

type HttpClient struct {
    client *http.Client
    timeout time.Duration
    retries int
}

func NewHttpClient(timeout time.Duration, retries int) *HttpClient {
    return &HttpClient{
        client: &http.Client{},
        timeout: timeout,
        retries: retries,
    }
}

func (c *HttpClient) SetHeader(req *http.Request, headers map[string]string) {
    for k, v := range headers {
        req.Header.Set(k, v)
    }
}

func (c *HttpClient) Get(url string, headers map[string]string) (*http.Response, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    c.SetHeader(req, headers)
    return c.Do(req)
}

func (c *HttpClient) Post(url string, headers map[string]string, body interface{}) (*http.Response, error) {
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, err
    }
    c.SetHeader(req, headers)
    return c.Do(req)
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
    var resp *http.Response
    var err error
    for i := 0; i < c.retries; i++ {
        resp, err = c.client.Do(req)
        if err == nil {
            break
        }
        time.Sleep(c.timeout)
    }
    return resp, err
}

// This implementation allows us to define the timeout and retry count when creating a new instance of the HttpClient.
// We can also set custom headers for each request, and send JSON data in the request body for POST requests.
// The Get and Post methods are used to send GET and POST requests respectively, and the Do method can be used to send any type of request.
// The Do method will automatically retry the request up to the specified number of times if it fails due to a network error.