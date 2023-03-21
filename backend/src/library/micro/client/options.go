package client

import (
	"../../micro/grpcx/resolver"
	"../../micro/options"
)

type Options struct {
	target       string
	authority    string
	Timeout      time.Duration
	resolverOpts []options.Options
}

type Option func(*Options)

func target(serviceName string) Option {
	return func(o *Options) {
		resolv := options.NewOptions()
		for _, op := range o.resolverOpts {
			op(&resolv)
		}

		c, err := resolver.NewEtcdClient(resolv)
		defer c.Close()
		if err != nil {
			logger.Errorf("Etcd client generating error: %v", err)
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
		defer cancel()

		var schema string
		resp, err := c.KV.Get(ctx, resolver.SchemaKey, clientv3.WithSerializable())
		if err != nil {
			logger.Errorf("grpc client: can not getting prefix from etcd: %v", err)
			return
		} else {
			schema = string(resp.Kvs[0].Value)
		}
		o.target = fmt.Sprintf("%s://%s/%s", schema, o.authority, util.Escape(serviceName))
	}
}

func Timeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.timeout = timeout
	}
}

func Authority(authority string) Option {
	return func(opts *Options) {
		opts.authority = authority
	}
}

func ResolverOptions(ropts ...options.Option) Option {
	return func(opts *Options) {
		opts.resolverOpts = ropts
	}
}

func newOptions(opt ...Option) Options {
	opts := Options{
		authority:    "@",
		timeout:      options.DefaultTimeout,
		resolverOpts: make([]options.Option, 0),
	}

	for _, o := range opt {
		o(&opts)
	}

	return options
}
