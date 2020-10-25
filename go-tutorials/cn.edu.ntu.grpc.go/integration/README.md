## issue list
1. import other proto

    - download the source from [github](https://github.com/protocolbuffers/protobuf)
    - use follow command
    
        ```shell
        import "google/protobuf/timestamp.proto";
        protoc -I E:/dev/Go/include/protobuf/src/ -I protos/ ./protos/*.proto --go_out=plugins=grpc:./pb
        ```
      
    - `-I` is to specify the import proto path, which is `proto_included` path.
    - it's proto_included path + import path should be right.
    
        ```js
        // -I: E:/dev/Go/include/protobuf/src/
        // import: google/protobuf/timestamp.proto
        // so E:/dev/Go/include/protobuf/src/google/protobuf/timestamp.proto file is existence.
        ```
      
2. During actual development, the client and server are completely independent
    
    - proto
    - model
    - certs
    
3. issue list: // TODO:
    
    - certs and key are not worked properly
    - ~~InsertedID will be binary data.~~: please use official lib
    - cannot use `*pb.EmployeeResponse` as master in **Dto2Model** method