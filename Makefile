build:
	go build -o ./bin/main.exe cmd/main.go

run: build
	./bin/main.exe
proto:
	protoc --go_out=./internal  --go-grpc_out=./internal --proto_path=./submodules/grpc-bookshelf-proto/proto book.proto