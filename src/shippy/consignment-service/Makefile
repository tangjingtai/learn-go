build:
# protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/consignment-service proto/consignment/consignment.proto
# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
	set GOOS=linux 	
	set GOARCH=amd64 	
	go build -o consignment-service .
	
# 根据当前目录下的 Dockerfile 生成名为 consignment-service 的镜像 
	docker build -t 192.168.1.101:5000/consignment-service:v1 .
	docker push 192.168.1.101:5000/consignment-service:v1
#run:
#	docker run -d -p 50051:50051 consignment-service