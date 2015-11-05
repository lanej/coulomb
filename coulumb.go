package coulumb

import (
	"net/url"
)

// Env represents accumulated configuration
type Env struct {
	Method          string
	RequestHeaders  map[string]string
	ResponseHeaders map[string]string
	RequestBody     string
	URL             url.URL
	ResponseBody    string
	Status          int
	Error           error
}

// Response wraps the env with some helper fucntions
type Response struct {
	Env Env
}

// Success returns true if 2xx or redirect
func (r Response) Success() bool {
	return ((r.Env.Status > 199) && (r.Env.Status < 300)) || (r.Env.Status == 302)
}

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
	url, _ := url.Parse(path)

	env := client.Adapter.Call(Env{URL: *url, Method: "GET"})

	return Response{Env: env}
}

// RackAdapter passed the request throught the specified Application
type RackAdapter struct {
	Application func(map[string]interface{}) (status int, headers map[string]string, body string, err error)
}

// NetHTTPAdapter passes the request through golang net/http
type NetHTTPAdapter struct {
}

// Call builds a rack-compliant env and calls the Application
func (a RackAdapter) Call(env Env) Env {
	rackEnv := map[string]interface{}{
		"rack.input":     env.RequestBody,
		"REQUEST_METHOD": env.Method,
		"SCRIPT_NAME":    "",
	}

	status, headers, body, err := a.Application(rackEnv)

	env.ResponseHeaders = headers
	env.Status = status
	env.ResponseBody = body
	env.Error = err

	return env
}

// Call forms the net/http request and executes
func (a NetHTTPAdapter) Call(env Env) Env {
	return env
}

// Build a Client from some configuration
func Build(config map[string]interface{}) {
}
