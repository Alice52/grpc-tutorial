## go grpc tutorial

1. env

   - go: `go-1.13.4`
   - database: `mongodb-4.2`
   - protocol: `proto-3`

2. intoduce

   - CRUD operations in mongodb with gRPC using Go

3. command

   ```shell
   go get github.com/golang/protobuf/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

   protoc blog.proto --go_out=plugins=grpc:.
   ```
