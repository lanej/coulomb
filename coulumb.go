package coulumb

import (
	"net/url"
)

type Env struct {
	Method          string
	RequestHeaders  map[string]string
	ResponseHeaders map[string]string
	RequestBody     string
	Url             url.URL
	ResponseBody    string
	Status          int
	Error           error
}

type Response struct {
	Env Env
}

func (env *Env) Success() bool {
	return ((env.Status > 199) && (env.Status < 300)) || (env.Status == 302)
}

type Adapter interface {
	Call(Env) Env
}

type Client struct {
	Url     string  `base url`
	Adapter Adapter `connection adapter`
}

func (client *Client) Get(path string) Response {
	url, _ := url.Parse(path)

	env := client.Adapter.Call(Env{Url: *url, Method: "GET"})

	return Response{Env: env}
}

type RackAdapter struct {
	Application func(map[string]interface{}) (status int, headers map[string]string, body string, err error)
}

type NetHttpAdapter struct {
}

func (a RackAdapter) Call(env Env) Env {
	status, headers, body, err := a.Application(map[string]interface{}{
		"rack.input":     env.RequestBody,
		"REQUEST_METHOD": env.Method,
		"SCRIPT_NAME":    "",
	})

	env.ResponseHeaders = headers
	env.Status = status
	env.ResponseBody = body
	env.Error = err

	return env
}

func (a NetHttpAdapter) Call(env Env) Env {
	return env
}

func Build(config map[string]interface{}) {
}
