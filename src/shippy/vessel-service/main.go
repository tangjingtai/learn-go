// 货轮服务，提供运输获取的货轮
package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"os"
	vesselProto "shippy/vessel-service/proto/vessel"
)

const (
	DEFAULT_HOST = "localhost:27017"
)

func main() {
	// 获取容器设置的数据库地址环境变量的值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_HOST
	}
	session, err := CreateSession(dbHost)
	// 创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	// 停留在港口的货船，先写死
	repo := &VesselRepository{session.Copy()}
	CreateDummyData(repo)

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"192.168.1.101:8500",
		}
	})
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
		micro.Registry(reg),
		//micro.Address("localhost:0"),
	)

	// 解析命令行参数
	server.Init()
	vesselProto.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CreateDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*vesselProto.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}
