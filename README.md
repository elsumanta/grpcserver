# grpcserver

Architecture: Clean Architecture 

Prequisite:
1. setup go evn
2. install grpc library https://grpc.io/docs/languages/go/quickstart/
3. install posgreSql library
4. folow https://github.com/gusaul/grpcox to install grpcox interface

Folder GRPC:
quick-start: https://grpc.io/docs/languages/go/quickstart/
added new contract and compile with command:
protoc --proto_path=${GOPATH}/src --proto_path=${GOPATH}/src/github.com/protobuf -I  grpc/  grpc/contract.proto --go_out=plugins=grpc:grpc

How to run postgresql server:
1. `docker-compose up --build -d`

How to stop postgresql server:
1. `docker-compose down`

How to run:
1. to directory server
2. run command `go mod vendor`
3. run `go build && ./server`
3. run grpcox and connect to server port 50051
