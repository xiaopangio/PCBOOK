gen:
	protoc -I proto proto/*.proto --go_out=pb --go-grpc_out=pb
clean:
	del .\pb\*.go
server:
	go run cmd/server/main.go -port 8080
client:
	go run cmd/client/main.go -address 0.0.0.0:8080 
test:
	go test -cover -race ./...
cert:
	cd ./cert; ./gen.sh; cd ..
.PHONY: gen clean server client test cert
