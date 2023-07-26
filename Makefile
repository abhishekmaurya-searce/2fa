proto:
	protoc -Iproto/ --go_opt=module=github.com/abhishekmaurya0/2fa --go_out=. --go-grpc_opt=module=github.com/abhishekmaurya0/2fa --go-grpc_out=. proto/*.proto
mysql:
	podman run --name mysql-container -e MYSQL_ROOT_PASSWORD=secret -p :5432 -d mysql:latest  
.PHONY: proto,mysql