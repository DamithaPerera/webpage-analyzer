package services

import "net/http"

//interface for making HTTP GET requests
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

//the default implementation of HTTPClient using the http package
type DefaultHTTPClient struct{}

//Get sends an HTTP GET request to the given URL and returns the response
func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
