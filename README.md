# grpcserver

Architecture: Clean Architecture 

Prequisite:
1. setup go evn https://go.dev/doc/install
2. install grpc library https://grpc.io/docs/languages/go/quickstart/
*3. install posgreSql [library](https://www.postgresql.org/download/windows/)
4. folow https://github.com/gusaul/grpcox to install grpcox interface

Use docker postgre container:
  How to run postgresql server: `docker-compose up --build -d`
  How to stop postgresql server: `docker-compose down`

Database Config:
	dbhost   = "localhost"
	dbport   = 5656
	user     = "server"
	password = "server"
	dbname   = "grpc-server"

How to run:
1. cd directory server
2. run command `go mod vendor`
3. run `go build && ./server`
3. run grpcox and connect to server port 50051


*How to update grpc contract
quick-start: https://grpc.io/docs/languages/go/quickstart/
added new contract and compile with command:
protoc --proto_path=${GOPATH}/src --proto_path=${GOPATH}/src/github.com/protobuf -I  grpc/  grpc/contract.proto --go_out=plugins=grpc:grpc

Noted:
*: ifneed
