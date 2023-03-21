package Options

import (
	"context"
	"crypto/tls"
	"time"
)

const (
	defaultAdress  = "127.0.0.1:2379"
	DefaultTimeout = 5 * time.Second
)

type Options struct {
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config
	Context   context.Context
}

type Option func(*Options)

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

func TLSConfig(tls *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = tls
	}
}

type AuthKey struct{}

type AuthCreds struct {
	Username string
	Password string
}

func Auth(username string, password string) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, AuthKey{}, &AuthCreds{Username: username, Password: password})
	}
}

func NewOptions() Options {
	return Options{
		Addrs:   []string{defaultAdress},
		Timeout: DefaultTimeout,
	}
}
