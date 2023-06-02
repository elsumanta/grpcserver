# grpcserver
Folder GRPC:
quick-start: https://grpc.io/docs/languages/go/quickstart/
added new contract and compile with command:
protoc --proto_path=${GOPATH}/src --proto_path=${GOPATH}/src/github.com/protobuf -I  grpc/  grpc/contract.proto --go_out=plugins=grpc:grpc


