package rest

import "time"

// DefaultTimeout request doc ...
var DefaultTimeout = 25 * time.Second

// RequestInfo doc ...
type RequestInfo struct {
	Method   string
	Endpoint string
	Headers  map[string]string
	Query    map[string]string
	Body     interface{}
	Timeout  time.Duration
}
