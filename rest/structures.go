package rest

import "time"

// RequestInfo doc ...
type RequestInfo struct {
	Method      string
	Endpoint    string
	Headers     map[string]string
	QueryParams map[string]string
	Body        interface{}
	Timeout     time.Duration
}
