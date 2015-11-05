package coulomb

// RackAdapter passed the request throught the specified Application
type RackAdapter struct {
	Application func(map[string]interface{}) (status int, headers map[string]string, body string, err error)
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
