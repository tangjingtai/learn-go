package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"io/ioutil"
	"log"
	"os"
	consProto "shippy/consignment-service/proto/consignment"
)

const (
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*consProto.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *consProto.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"192.168.1.101:30500",
		}
	})
	server := micro.NewService(
		micro.Name("consignment.client"),
		micro.Registry(reg),
	)

	// 创建微服务的客户端，简化了手动 Dial 连接服务端的步骤
	client := consProto.NewShippingServiceClient("go.micro.srv.consignment", server.Client())
	// 调用 RPC
	// 将货物存储到我们自己的仓库里
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	// 新货物是否托运成功
	log.Printf("created: %t", resp.Created)
	if !resp.Created {
		return
	}
	resp, err = client.GetConsignments(context.Background(), &consProto.GetRequest{})
	if err != nil {
		log.Fatalf("get consignment error: %v", err)
	} else {
		log.Printf("get consignment :%v", resp.Consignments)
	}
}
