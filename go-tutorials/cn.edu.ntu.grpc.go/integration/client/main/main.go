package main

import (
    "cn.edu.ntu.grpc.go/integration/client/model"
    "cn.edu.ntu.grpc.go/integration/client/protos/pb"
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/metadata"
    "google.golang.org/protobuf/types/known/timestamppb"
    "io"
    "log"
    "os"
    "time"
)

const port = ":5001"

func main() {
    // 1. create security connection to gRPC server
    creds, err := credentials.NewClientTLSFromFile("./integration/client/certs/cert.pem", "")
    if err != nil {
        log.Fatalf("create creds error: %v", err)
    }
    
    options := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
    conn, err := grpc.Dial("localhost"+port, options...)
    if err != nil {
        log.Fatalf("create connection error: %v", err)
    }
    
    // create request
    cRequest := &pb.EmployeeRequest{
        Employee: &pb.Employee{
            No:        1,
            FirstName: "zack",
            LastName:  "zhang",
            MonthSalary: &pb.MonthSalary{
                Basic: 500,
                Bonus: 520,
            },
            Status: pb.EmployeeStatus_NORMAL,
            LastModified: &timestamppb.Timestamp{Seconds:
            time.Now().Unix()},
        },
    }
    client := NewEmployeeServiceClient(conn)
    
    // 1. create
    employees, err := client.save(cRequest)
    log.Println(employees)
    
    // 2. list
    list, _ := client.getAll()
    for _, employee := range list {
        log.Println(employee)
    }
    
    // 3. upload
    client.uploadFile(employees.Id.Hex(), "./integration/README.md")
    
    // 4. save_all
    var listReqs []*pb.EmployeeRequest
    for i := 0; i < 5; i++ {
        listReqs = append(listReqs, cRequest)
    }
    es, _ := client.saveAll(listReqs)
    for _, e := range es {
        log.Println(e)
    }
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) *EmployeeServiceClient {
    aa := pb.NewEmployeeServiceClient(cc)
    return &EmployeeServiceClient{aa}
}

type EmployeeServiceClient struct {
    cc pb.EmployeeServiceClient
}

func (client *EmployeeServiceClient) saveAll(cRequests []*pb.EmployeeRequest) ([]*model.Employee, error) {
    // 1. get stream
    stream, err := client.cc.SaveAll(context.Background())
    if err != nil {
        log.Fatalf("call to save-all grpc server failed: %v")
    }
    
    // list for response
    var list []*model.Employee
    finishChannel := make(chan struct{})
    go func() {
        for {
            data, err := stream.Recv()
            if err == io.EOF {
                finishChannel <- struct{}{}
                break
            }
            
            if err != nil {
                log.Fatalf("recv data from grpc server failed: %v", err)
            }
            employee := Dto2Model(data)
            list = append(list, &employee)
        }
    }()
    
    // 2. send data
    for _, req := range cRequests {
        err = stream.Send(req)
        if err != nil {
            log.Fatalf("send data to grpc server failed: %v")
        }
    }
    stream.CloseSend()
    <-finishChannel
    
    return list, nil
}

func (client *EmployeeServiceClient) save(cRequest *pb.EmployeeRequest) (model.Employee, error) {
    cResponse, err := client.cc.Save(context.Background(), cRequest)
    if err != nil {
        log.Fatalf(err.Error())
    }
    
    return Dto2Model(cResponse), nil
}

func (client *EmployeeServiceClient) getAll() ([]model.Employee, error) {
    stream, err := client.cc.GetByAll(context.Background(), new(pb.GetByAllRequest))
    if err != nil {
        log.Fatal(err.Error())
    }
    
    var list []model.Employee
    for {
        res, err := stream.Recv()
        // If end of stream, break the loop
        if err == io.EOF {
            break
        }
        
        if err != nil {
            return nil, err
        }
        list = append(list, Dto2Model(res))
    }
    
    return list, nil
}

func (client *EmployeeServiceClient) uploadFile(id string, path string) {
    // 1. open file
    imageFile, err := os.Open(path)
    if err != nil {
        log.Fatal("open file error: %v", err)
    }
    defer imageFile.Close()
    
    // 2. prepare metadata
    md := metadata.New(map[string]string{"id": id})
    context := metadata.NewOutgoingContext(context.Background(), md)
    
    // 3. get send stream
    stream, err := client.cc.AddPhoto(context)
    if err != nil {
        log.Fatalf("call grpc add photo failed: %v", err)
    }
    
    // 4. send data
    for {
        chunk := make([]byte, 1024)
        chunkSize, err := imageFile.Read(chunk)
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("write file to grpc server failed: %v", err)
        }
        
        if chunkSize < len(chunk) {
            chunk = chunk[:chunkSize]
        }
        
        stream.Send(&pb.AddPhoneRequest{Data: chunk})
    }
    
    // 5. receive response
    
    res, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("receive data from grpc server failed: %v", err)
    }
    
    log.Printf("add phone request is success: %v\n", res.IsOk)
}

func Dto2Model(response *pb.EmployeeResponse) model.Employee {
    oid, _ := primitive.ObjectIDFromHex(response.Employee.Id)
    
    return model.Employee{
        Id:        oid,
        No:        response.Employee.No,
        FirstName: response.Employee.FirstName,
        LastName:  response.Employee.LastName,
        MonthSalary: &model.MonthSalary{
            Basic: response.Employee.MonthSalary.Basic,
            Bonus: response.Employee.MonthSalary.Bonus,
        },
        Status:       response.Employee.Status,
        LastModified: response.Employee.LastModified,
    }
}
