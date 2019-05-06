build:
# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统 
	set GOOS=linux 	
	set GOARCH=amd64 	
	go build -o consignment-client
	
# 根据当前目录下的 Dockerfile 生成名为 consignment-client 的镜像 
	docker build -t consignment-client .
run:
	docker run consignment-client