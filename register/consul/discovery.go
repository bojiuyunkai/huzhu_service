//地址 https://github.com/janlely/consul-go-grpc-demo/
package consul
import (
	"errors"
	
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"regexp"
	"sync"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func init(){
	logs.SetLogger(logs.AdapterFile,`{"filename":"/tmp/discovery.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}
const (
	defaultPort = "8500"//consul 默认端口
)

var (
	errMissingAddr = errors.New("consul resolver: missing address")

	errAddrMisMatch = errors.New("consul resolver: invalied uri")

	errEndsWithColon = errors.New("consul resolver: missing port after port-separator colon")

	regexConsul, _ = regexp.Compile("^([A-z0-9.]+)(:[0-9]{1,5})?/(.*)$")
)
var logger log.Logger;
func Init(logger log.Logger) {
	logger = logger;
	level.Info(logger).Log("calling", "discovery.go", "method","Init")
	resolver.Register(NewBuilder())
}

type consulBuilder struct {
}

type consulResolver struct {
	address              string//consul 地址
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	name                 string//服务名
	disableServiceConfig bool
	lastIndex            uint64
}

func NewBuilder() resolver.Builder {
	return &consulBuilder{}
}

func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	host, port, name, err := parseTarget(fmt.Sprintf("%s/%s", target.Authority, target.Endpoint))
	
	logs.Debug("Build host: ", host)
	//fmt.Println("host",host,"port",port,"name",name);
	if err != nil {
		return nil, err
	}
	cr := &consulResolver{
		address:              fmt.Sprintf("%s%s", host, port),
		name:                 name,
		cc:                   cc,
		disableServiceConfig: opts.DisableServiceConfig,
		lastIndex:            0,
	}

	cr.wg.Add(1)
	go cr.watcher()
	return cr, nil

}

func (cr *consulResolver) watcher() {
	config := api.DefaultConfig()
	config.Address = cr.address
	client, err := api.NewClient(config)
	if err != nil {
		//fmt.Println("error create consul client:",err)
		return
	}

	for {
		services, metainfo, err := client.Health().Service(cr.name, "", true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			//fmt.Println("error retrieving instances from Consul",err)
		}

		
		cr.lastIndex = metainfo.LastIndex
		var newAddrs []resolver.Address
		for _, service := range services {
			addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
			newAddrs = append(newAddrs, resolver.Address{Addr: addr})
		}
		//fmt.Printf("adding service addrs\n")
		//fmt.Printf("newAddrs: %v\n", newAddrs)
		cr.cc.NewAddress(newAddrs)
		cr.cc.NewServiceConfig(cr.name)
	}

}

func (cb *consulBuilder) Scheme() string {
	return "consul"
}

func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOptions) {
}

func (cr *consulResolver) Close() {
}

func parseTarget(target string) (host, port, name string, err error) {

	//fmt.Printf("target uri: %v\n", target)
	if target == "" {
		return "", "", "", errMissingAddr
	}
	//fmt.Println(regexConsul.MatchString(target));
	if !regexConsul.MatchString(target) {
		return "", "", "", errAddrMisMatch
	}

	groups := regexConsul.FindStringSubmatch(target)
	host = groups[1]
	port = groups[2]
	name = groups[3]
	if port == "" {
		port = defaultPort
	}
	return host, port, name, nil
}