package http

import (
	"time"
)

const (
	// defaultMaxIdleConnectionsPerHost is the default value for MaxIdleConnectionsPerHost
	defaultMaxIdleConnectionsPerHost = 10
	// defaultResponseTimeout is the default value for ResponseTimeout
	defaultResponseTimeout = 5 * time.Second
	// defaultRequestTimeout is the default value for RequestTimeout
	defaultRequestTimeout = 5 * time.Second
)

// timeoutImpl is a struct that holds the configuration for the http client
//
//	RequestTimeout: the Timeout for the request
//	ResponseTimeout: the Timeout for the response
//	MaxIdleConnections: the maximum number of idle connections
type timeoutImpl struct {
	ResponseTimeout    time.Duration
	RequestTimeout     time.Duration
	MaxIdleConnections int
	DisableTimeouts    bool
}

func newTimeouts() *timeoutImpl {
	return &timeoutImpl{
		ResponseTimeout:    defaultResponseTimeout,
		RequestTimeout:     defaultRequestTimeout,
		MaxIdleConnections: defaultMaxIdleConnectionsPerHost,
		DisableTimeouts:    false,
	}
}

type Timeout interface {
	SetRequestTimeout(time.Duration) Timeout
	SetResponseTimeout(time.Duration) Timeout
	SetMaxIdleConnections(int) Timeout
	GetMaxIdleConnections() int
	GetRequestTimeout() time.Duration
	GetResponseTimeout() time.Duration
	Disable() Timeout
	Enable() Timeout
}

// GetRequestTimeout returns the request Timeout
// if the request Timeout is not set, it returns the default request Timeout.
func (c timeoutImpl) GetRequestTimeout() time.Duration {
	if c.RequestTimeout != defaultRequestTimeout {
		return c.RequestTimeout
	}
	if c.DisableTimeouts {
		return 0
	}
	return defaultRequestTimeout
}

// GetResponseTimeout returns the response Timeout
// if the request Timeout is not set, it returns the default response Timeout
func (c timeoutImpl) GetResponseTimeout() time.Duration {
	if c.ResponseTimeout != defaultResponseTimeout {
		return c.ResponseTimeout
	}
	if c.DisableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

// GetMaxIdleConnections
//
// Returns the maximum number of idle connections
// if the value is not set, it returns the default value.
//
// default value is 10
func (c timeoutImpl) GetMaxIdleConnections() int {
	if c.MaxIdleConnections != defaultMaxIdleConnectionsPerHost {
		return c.MaxIdleConnections
	}
	return defaultMaxIdleConnectionsPerHost
}

// Disable disables the timeouts for the client
// if the value is true, the timeouts will be disabled
// and the client will not time out.
//
//	Example:
//		client.DisableTimeouts(true)
func (c timeoutImpl) Disable() Timeout {
	c.DisableTimeouts = true
	return c
}

func (c timeoutImpl) Enable() Timeout {
	c.DisableTimeouts = false
	return c
}

// SetRequestTimeout sets the request Timeout
// if the value is 0, the request will not Timeout.
//
// if the value is negative, the request will time out immediately.
//
// if the value is positive, the request will time out after the specified duration.
//
// if the value is not set, the request will time out after 5 seconds.
//
// if the value is not set and the DisableTimeouts is set to true, the request will not Timeout.
//
//	Example:
//		client.SetRequestTimeout(10 * time.Second)
func (c timeoutImpl) SetRequestTimeout(timeout time.Duration) Timeout {
	c.RequestTimeout = timeout
	return c
}

// SetResponseTimeout sets the response Timeout.
// 	If the value is 0, the response will not Timeout.
//	If the value is negative, the response will time out immediately.
//  If the value is positive, the response will time out after the specified duration.
//	If the value is not set, the response will time out after 5 seconds.
//	 If the value is not set and the DisableTimeouts is set to true, the response will not Timeout.
//
//	Example:
//		client.SetResponseTimeout(10 * time.Second)
//		client.SetResponseTimeout(0)
//		client.SetResponseTimeout(-1)
//		client.SetResponseTimeout(10 * time.Second)
func (c timeoutImpl) SetResponseTimeout(timeout time.Duration) Timeout {
	c.ResponseTimeout = timeout
	return c
}

// SetMaxIdleConnections sets the maximum number of idle connections.
// if the value is not set, the default value is 10. 
//
// if the value is 0, the maximum number of idle connections is unlimited.
//
// if the value is negative, the maximum number of idle connections is unlimited.
//
// if the value is positive, the maximum number of idle connections is the specified value.
//
//	Example:
//		client.SetMaxIdleConnections(10)
//		client.SetMaxIdleConnections(0)
//		client.SetMaxIdleConnections(-1)
func (c timeoutImpl) SetMaxIdleConnections(maxConnections int) Timeout {
	c.MaxIdleConnections = maxConnections
	return c
}
