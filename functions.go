package consumer

import (
	"log"
	"time"
)

// DefaultTimeout request doc ...
var DefaultTimeout = 25 * time.Second

// PrintError doc...
var PrintError func(...interface{}) = log.Print

// Print doc ...
var Print func(string, ...interface{}) = log.Printf

// Fatal doc ...
var Fatal func(...interface{}) = log.Fatal

// LogResponse doc ...
func LogResponse(buff []byte) {
	if len(buff) > 2000 {
		Print("[REST integration] RESPONSE: %v", string(buff[:1000]), " ••• SKIPED ••• ", string(buff[len(buff)-1000:]))
	} else {
		Print("[REST integration] RESPONSE: %v", string(buff))
	}
}

// LogRequest doc ...
func LogRequest(buff []byte) {
	if len(buff) > 2000 {
		Print("[REST integration] REQUEST: %v", string(buff[:1000]), " ••• SKIPED ••• ", string(buff[len(buff)-1000:]))
	} else {
		Print("[REST integration] REQUEST: %v", string(buff))
	}
}
