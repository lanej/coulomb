package coulomb

import "net/url"

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
