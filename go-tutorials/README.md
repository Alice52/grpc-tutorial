## go grpc tutorial

1. env

   - go: `go-1.13.4`
   - database: `mongodb-4.2`
   - protocol: `proto-3`
   - protoc: `libprotoc 3.13.0`
   - protoc-gen-go: `v1.25.0`

     ```shell
     export PATH=$PATH:$(go env GOPATH)/bin

     brew install protoc # macos
     go install github.com/golang/protobuf/protoc-gen-go@latest
     go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
     ```

2. introduce

   - CRUD operations in mongodb with gRPC using Go

3. protoc generate

   ```shell
   # protoc ./proto/blog.proto --go_out=./ --plugin=grpc:.
   protoc ./proto/blog.proto --go_out=plugins=grpc:./
   ```

4. workflow

   - Protbuf Message (Request)
   - Regular Go Struct
   - Convert to BSON + Mongo Action
   - Protobuf Message (Response)

5. security

   ```shell
    # https://blog.csdn.net/weixin_30531261/article/details/80891360
    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "//CN=localhost"
   ```

   - [concept](https://blog.csdn.net/earbao/article/details/82958518)
