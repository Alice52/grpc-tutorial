package main

import (
    blogpb "cn.edu.ntu.grpc.go/mongo/proto"
    "github.com/golang/protobuf/jsonpb"
    "github.com/golang/protobuf/proto"
    "io/ioutil"
    "log"
)

func main() {
    blog := blogpb.Blog{
        Id:       "132",
        AuthorId: "zack",
        Title:    "title",
        Content:  "content",
    }
    log.Println(blog)
    
    // protobuf
    write2FileWithProto("./serialize/message.bin", &blog)
    blog2, _ := readFromFileWithProto("./serialize/message.bin")
    log.Println(blog2)
    
    // json
    str := toJSONWithProto(&blog)
    log.Println(str)
    blog3, _ := fromJsonWithProto(str)
    log.Println(blog3)
}

func fromJsonWithProto(str string) (*blogpb.Blog, error) {
    blog := blogpb.Blog{}
    err := jsonpb.UnmarshalString(str, &blog)
    
    if err != nil {
        log.Fatalf("from json error: %v\n", err.Error())
    }
    
    return &blog, nil
}

func toJSONWithProto(pb proto.Message) string {
    
    marshaller := jsonpb.Marshaler{
        Indent: "    ",
    }
    
    str, err := marshaller.MarshalToString(pb)
    if err != nil {
        log.Fatalf("to json error: %v\n", err.Error())
    }
    
    return str
}

func write2FileWithProto(path string, pb proto.Message) error {
    dataBytes, err := proto.Marshal(pb)
    
    if err != nil {
        log.Fatalf("serialize error: %v\n", err.Error())
    }
    
    if err = ioutil.WriteFile(path, dataBytes, 0644); err != nil {
        log.Fatalf("writeto file error: %v\n", err.Error())
    }
    
    return nil
}

func readFromFileWithProto(path string) (*blogpb.Blog, error) {
    dataBytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatalf("read file error: %v\n", err.Error())
    }
    blog := &blogpb.Blog{}
    err = proto.Unmarshal(dataBytes, blog)
    
    if err != nil {
        log.Fatalf("Unmarshal data bytes error: %v\n", err.Error())
    }
    
    return blog, nil
}
