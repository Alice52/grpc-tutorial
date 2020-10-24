package main

import (
    enumpb "cn.edu.ntu.grpc.go/enum"
    "log"
)

func main() {
    em := NewEnumMessage()
    log.Println(em)
    name := enumpb.Gender_name[int32(em.Gender)]
    value := enumpb.Gender_value[em.Gender.String()]
    log.Printf("name: %v, value: %v\n", name, value)
}

func NewEnumMessage() *enumpb.EnumMessage {
    em := enumpb.EnumMessage{
        Id:     132,
        Gender: enumpb.Gender_FEMALE,
    }
    em.Gender = enumpb.Gender_MALE
    
    return &em
}
