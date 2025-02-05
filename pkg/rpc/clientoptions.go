package rpc

import (
	"net/url"
	"time"

	"github.com/chia-network/go-chia-libs/pkg/config"
	"github.com/chia-network/go-chia-libs/pkg/httpclient"
	"github.com/chia-network/go-chia-libs/pkg/rpcinterface"
)

// WithAutoConfig automatically loads chia config from CHIA_ROOT
func WithAutoConfig() rpcinterface.ConfigOptionFunc {
	return func() (*config.ChiaConfig, error) {
		return config.GetChiaConfig()
	}
}

// WithManualConfig allows supplying a manual configuration for the RPC client
func WithManualConfig(cfg config.ChiaConfig) rpcinterface.ConfigOptionFunc {
	return func() (*config.ChiaConfig, error) {
		return &cfg, nil
	}
}

// WithBaseURL sets the host for RPC requests
func WithBaseURL(url *url.URL) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		return c.SetBaseURL(url)
	}
}

// WithCache specify a duration http requests should be cached for
// If unset, cache will not be used
func WithCache(validTime time.Duration) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		typed, ok := c.(*httpclient.HTTPClient)
		if ok {
			typed.SetCacheValidTime(validTime)
		}

		return nil
	}
}

// WithTimeout sets the timeout for the requests
func WithTimeout(timeout time.Duration) rpcinterface.ClientOptionFunc {
	return func(c rpcinterface.Client) error {
		typed, ok := c.(*httpclient.HTTPClient)
		if ok {
			typed.Timeout = timeout
		}

		return nil
	}
}
