package client

import (
	"context"
	"crypto/tls"
	"net/http"
)

// HTTPRequestExecutor defines the interface of the request executor for the HTTP transport protocol
type HTTPRequestExecutor func(ctx context.Context, req *http.Request) (*http.Response, error)

// DefaultHTTPRequestExecutor creates a HTTPRequestExecutor with the received HTTPClientFactory
func DefaultHTTPRequestExecutor(clientFactory HTTPClientFactory) HTTPRequestExecutor {
	return func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return clientFactory(ctx).Do(req.WithContext(ctx))
	}
}

// HTTPClientFactory creates http clients based with the received context
type HTTPClientFactory func(ctx context.Context) *http.Client

// NewHTTPClient just returns the http default client
func NewHTTPClient(ctx context.Context) *http.Client {
	return defaultHTTPClient
}

// NewHTTPClient just returns the http default client
func NewInsecureHTTPClient(ctx context.Context) *http.Client {
	return defaultInsecureHTTPClient
}

var defaultHTTPClient = &http.Client{}
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var defaultInsecureHTTPClient = &http.Client{Transport: tr}
