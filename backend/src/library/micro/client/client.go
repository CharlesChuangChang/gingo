package client

import (
	"context"
	"sync"

	"library/micro/options"
	"../../micro/grpcx/resolver"
	"google.golang.org/gprc"
	"google.golang.org/gprc/balancer/roundrobin"
	"go.etcd.io/etcd/clientv3"
	event "../../micro/registry"
	"../../micro/util"
	"../../../../../common/logger"
	"fmt"
	"path"
)

type grpcClient struct {
	sync.RWMutex
	opts   Options
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	once   sync.Once
}

type ServiceName struct{}

func (c *grpcClient) CreateConn() (*grpc.ClientConn, error) {
	c.Lock()
	defer c.Unlock()
	c.once.Do(func() {
		resolver.Init(c.opts.resolverOpts...)
	})

	c.ctx, c.cancel = context.WithTimeout(context.Background(), c.opts.Timeout)
	sc := `{"loadBalancingPolicy": "%s", "loadBalancingConfig":[{"%s":"{}"}]}`
	var grpcOptions []grpc.DialOption
	grpcOptions = append(grpcOptions, grpc.WithInsecure())
	grpcOptions = append(grpcOptions, grpc.WithBlock())
	grpcOptions = append(grpcOptions, grpc.WithDefaultServiceConfig(fmt.Sprintf(sc, roundrobin.Name, roundrobin.Name)))

	conn, err := grpc.DialContext(c.ctx, c.opts.target, grpcOptions...)
	if err != nil {
		return nil, error
	}

	return conn, nil
}


func (c *grpcClient) Emit(eventName string, payload string, opts ...clientv3.OpOptions) error {
	resolverOpts := options.NewOptions()
	for _, op : range c.opts.resolverOpts {
		op(&resolverOpts)
	}

	cc, err := resolver.NewEtcdClient(resolverOpts)
	defer cc.Close()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.opts.timeout)
	defer cancel()

	resp, err := cc.KV.Get(ctx, event.PrefixKey, clientv3.WithSerializable())
	if err != nil {
		return err
	} else {
		eventPrefix := string(resp.Kvs[0].Value)
		eventFullPath := path.Join(eventPrefix, util.Escape(eventName))

		var lgr *clientv3.LeaseGrantResponse
		ctx1, cancel1 := context.WithTimeout(context.Background(), c.opts.timeout)
		defer cancel1()

		lgr, err = cc.Grant(ctx1, int64(event.DefaultEventNodeTTL.Seconds()))
		if err != nil {
			return err
		}

		if lgr != nil {
			logger.Debugf("Event name '%s' has been emitted with leaseId %v", eventFullPath, lgr.ID)
			ctx2, cancel2 := context.WithTimeout(context.Background(), c.opts.timeout)
			defer cancel2()
			po := append([]clientv3.OpOption{
				clientv3.WithLease(lgr.ID),
			}, opts...)
			_, err = cc.Put(ctx2, eventFullPath, payload, po...)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


func NewClient(serviceName string, opts ...Option) grpcClient {
	opts = append(opts, target(serviceName))
	options := newOptions(opts...)

	cli := &grpcClient{
		opts:options,
	}

	return *cli
}

func NewGRPCConn(serviceName string, opts ...Option) (*grpc.ClientConn, error) {
	opts = append(opts, target(serviceName))
	options := newOptions(opts...)


	cli := &grpcClient{
		opts:options,
	}

	return cli.CreateConn()
}
