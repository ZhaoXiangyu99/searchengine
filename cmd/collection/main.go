package main

import (
	"net"
	"searchengine/cmd/collection/dal"
	collectionModel "searchengine/kitex_gen/collectionModel/collectionservice"
	"searchengine/pkg/constants"
	"searchengine/pkg/middleware"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
}

func main() {
	//微服务注册
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		panic(err)
	}
	Init()
	//服务设定
	svr := collectionModel.NewServer(new(CollectionServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CollectionServiceName}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
