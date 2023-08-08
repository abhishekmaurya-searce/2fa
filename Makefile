proto:
	protoc -Iproto/ --go_opt=module=github.com/abhishekmaurya0/2fa --go_out=. --go-grpc_opt=module=github.com/abhishekmaurya0/2fa --go-grpc_out=. proto/*.proto
client:
	go build -o bin/auth/client client/*.go
server:
	go build -o bin/auth/server cmd/main.go
runserver:
	./bin/auth/server
runclient:
	./bin/auth/client

.PHONY: proto client server runclient runserver