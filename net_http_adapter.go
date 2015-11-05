package coulomb

// NetHTTPAdapter passes the request through golang net/http
type NetHTTPAdapter struct {
}

// Call forms the net/http request and executes
func (a NetHTTPAdapter) Call(env Env) Env {
	return env
}
