build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/user-service proto/user/user.proto 
# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统 
	set GOOS=linux 	
	set GOARCH=amd64 	
	go build
	
# 根据当前目录下的 Dockerfile 生成名为 user-service 的镜像 
#docker build -t consignment-service .
run:
	docker run -d -p 50051:50051 consignment-service