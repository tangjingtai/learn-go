build:
# protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/vessel-service proto/vessel/vessel.proto
# dep 工具暂不可用，直接手动编译
	set GOOS=linux 
	set GOARCH=amd64 
	go build -o vessel-service .

# 根据当前目录下的 Dockerfile 生成名为 vessel-service 的镜像
	docker build -t 192.168.1.101:5000/vessel-service:v1 .
	docker push 192.168.1.101:5000/vessel-service:v1

#run:
#	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service