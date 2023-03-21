package resolver

import (
	"../../../micro/options"
	registry "../../../micro/registry/proto"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	etcd "go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc/resolver"
	"net"
)


var errMissingAddr = errors.New("etcd resolver: missing address")
var errAddrMisMatch = errors.New("etcd resolver:invalid uri")
var errPrefix = errors.New("etcd resolver: missing prefix")
var regexTargetEndpoint, _ = regexp.Compile("^[A-z0-9.]+(/[A-z0-9])*$")

var schemaM = new(sync.Map)
var lock = new(sync.Mutex)

const {
	PrefixKey = "/dlfc/micro/config/prefix"
	SchemaKey = "/dlfc/micro/config/grpc-schema"
}

type etcdGrpcBuilder struct {
	options options.Options
}

func (b *etcdGrpcBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	cli, err := NewEtcdClient(b.options)
	if err != nil {
		return nil, err
	}

	er := &EtcdGrpcResolver{
		c:cli,
		cc:cc,
	}
	er.ctx, er.cancel = context.WithTimeout(context.Background(), options.DefaultTimeout)
	er.target, err = er.parseTarget(target.Endpoint)
	if err != nil {
		return nil,err
	}

	er.wg.Add(1)
	go er.watcher()
	return er, nil
}


func (b *etcdGrpcBuilder) Scheme() string {
	retrun "etcd"
}

type EtcdGrpcResolver struct {
	c *etcd.Client
	target string
	wg sync.WaitGroup
	ctx context.Context
	cancel context.CancelFunc
	cc resolver.ClientConn
	wch ectd.WatchChan
}

func (er *EtcdGrpcResolver) ResolveNow(resolver.ResolveNowOptions) {
	logger.Debugf("etcd resolver: has been resolved now.but do nothing") 
}

func (er *EtcdGrpcResolver) Close() {
	logger.Infof("etcd resolver: has been closed")
	er.cancel()
	er.wg.Wait()
	er.c.Close()
}

func (er *EtcdGrpcResolver) watcher() {
	defer er.wg.Done()
	var addrList []resolver.Address
	if er.wch == nil {
		addrList, _ = er.firstWatch()
	}

	for {
		select {
			case <-er.ctx.Done():
				logger.Info("etcd resolver:watcher has been closed")
				return
				default:
					resp := <-er.wch
					for _, ev := range resp.Events {
						switch ev.Type {
						case mvccpb.PUT:
							
					}
				}
			}
	}
}


