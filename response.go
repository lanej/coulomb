package coulomb

// Response wraps the env with some helper fucntions
type Response struct {
	Env Env
}

// Success returns true if 2xx or redirect
func (r Response) Success() bool {
	return ((r.Env.Status > 199) && (r.Env.Status < 300)) || (r.Env.Status == 302)
}

// Body returns the response body.
func (r Response) Body() string {
	return r.Env.ResponseBody
}
