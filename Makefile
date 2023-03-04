gen:
	protoc -I proto proto/*.proto --go_out=pb --go-grpc_out=pb --grpc-gateway_out=pb  --openapiv2_out=swagger
clean:
	del .\pb\*.go
server:
	go run cmd/server/main.go -port 8080
rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 0.0.0.0:8080
client:
	go run cmd/client/main.go -address 0.0.0.0:8080
build:
	go run cmd/build/main.go
test:
	go test -cover -race ./...
cert:
	cd ./cert; sh gen.sh; cd ..
.PHONY: gen clean server client test cert build
