package shared

import "time"

const (
	Port           = 8081
	IdleTimeout    = time.Minute
	RequestTimeout = time.Second * 30
	ReadTimeout    = time.Second * 30
	WriteTimeout   = time.Second * 60
)
