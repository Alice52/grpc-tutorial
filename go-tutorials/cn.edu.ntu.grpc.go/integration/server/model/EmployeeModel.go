package model

import (
    "cn.edu.ntu.grpc.go/integration/server/protos/pb"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type Employee struct {
    Id           primitive.ObjectID `bson:"_id,omitempty"`
    No           int32
    FirstName    string
    LastName     string
    MonthSalary  *MonthSalary
    Status       pb.EmployeeStatus
    LastModified *timestamppb.Timestamp
}

type MonthSalary struct {
    Basic float32
    Bonus float32
}
