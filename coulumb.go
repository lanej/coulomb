package coulomb

import (
	"net/url"
)

// Adapter defines a Call() method that fires the request
type Adapter interface {
	Call(Env) Env
}

// Client represents a base url and a specific Adapter
type Client struct {
	URL     string
	Adapter Adapter
}

// Get performs a GET request
func (client *Client) Get(path string) Response {
	return client.processRequest(path, "GET")
}

// Put performs a PUT request
func (client *Client) Put(path string) Response {
	return client.processRequest(path, "PUT")
}

func (client *Client) processRequest(path string, method string) Response {
	url, _ := url.Parse(path)

	env := client.Adapter.Call(Env{URL: *url, Method: method})

	return Response{Env: env}
}

// Build a Client from some configuration
func Build(config map[string]interface{}) {
}
