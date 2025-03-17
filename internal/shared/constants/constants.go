package constants

import "time"

const (
	Port           = 8081
	IdleTimeout    = time.Minute
	RequestTimeout = time.Second * 3
	ReadTimeout    = time.Second * 3
	WriteTimeout   = time.Second * 30
)
