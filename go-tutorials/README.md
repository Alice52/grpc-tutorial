## go grpc tutorial

1. env

   - go: `go-1.13.4`
   - database: `mongodb-4.2`
   - protocol: `proto-3`
   - protoc: `libprotoc 3.13.0`
   - protoc-gen-go: `v1.25.0`

2. introduce

   - CRUD operations in mongodb with gRPC using Go

3. command

   ```shell
   go get github.com/golang/protobuf/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

   # protoc ./proto/blog.proto --go_out=./ --plugin=grpc:.
   protoc ./proto/blog.proto --go_out=plugins=grpc:./
   ```
   
4. workflow

    - Protbuf Message (Request) 
    - Regular Go Struct 
    - Convert to BSON + Mongo Action 
    - Protobuf Message (Response)   
