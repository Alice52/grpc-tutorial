package main

import (
    nestpb "cn.edu.ntu.grpc.go/nest"
    "log"
)

func main() {
    dm := NewDepartment()
    log.Println(dm)
}

func NewDepartment() *nestpb.DepartmentMessage {
    return &nestpb.DepartmentMessage{
        Id:   5678,
        Name: "test department",
        Employees: []*nestpb.EmployeeMessage{
            &nestpb.EmployeeMessage{
                Id:   1,
                Name: "Daven",
            },
            {
                Id:   2,
                Name: "Zack",
            },
            {
                Id:   3,
                Name: "Tim",
            },
        },
        ParentDepartment: &nestpb.DepartmentMessage{
            Id:   1122,
            Name: "bus",
        },
    }
}
