gen_proto:
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/transmitter.proto

protoc-install:
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

run-server:
	go mod tidy
	go run ./cmd/server/main.go

run-client:
	go mod tidy
	go run ./cmd/client/main.go

run-db:
	docker-compose up -d

build-client:
	go mod tidy
	go build -o client ./cmd/client/main.go

build-server:
	go mod tidy
	go build -o server ./cmd/server/main.go