# 路径
TOP_PATH = $(shell pwd) 
# s 启动服务类型 w-代表web g-代表grpc
s  = w
# 启动服务
run:
	rm -rf $(shell pwd)/protouser/*.go
	protoc --proto_path=$(shell pwd)/proto/ --micro_out=$(shell pwd)/protouser --go_out=$(shell pwd)/protouser $(shell pwd)/proto/*proto
	go run main.go $s
# grpc 客户端测试命令
co:
	go run $(shell pwd)/client/cliord.go
