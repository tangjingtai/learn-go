// 寄货服务
package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"os"
	consPb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
)

const (
	DEFAULT_HOST            = "localhost:27017"
	DEFAULT_CONSUL_REGISTRY = "192.168.1.101:8500"
)

func main() {
	//获取容器设置的数据库地址环境变量的值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_HOST
	}
	consulRegistry := os.Getenv("ConsulRegistry")
	if consulRegistry == "" {
		consulRegistry = DEFAULT_CONSUL_REGISTRY
	}
	session, err := CreateSession(dbHost)
	//创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	defer session.Close()

	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			consulRegistry,
		}
	})
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.Registry(reg),
		//micro.Address("localhost:0"),
	)

	// 解析命令行参数
	server.Init()
	// 作为 vessel-service 的客户端
	vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	// 将 server 作为微服务的服务端
	consPb.RegisterShippingServiceHandler(server.Server(), &handler{session, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
