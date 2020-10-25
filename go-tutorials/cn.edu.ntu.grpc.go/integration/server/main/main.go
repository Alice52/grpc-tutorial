package main

import (
    "cn.edu.ntu.grpc.go/config"
    "cn.edu.ntu.grpc.go/integration/server/model"
    "cn.edu.ntu.grpc.go/integration/server/protos/pb"
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "io"
    "log"
    "net"
)

const databaseName = "tutorials"
const collectionName = "employee"
const port = ":5001"

var db *mongo.Client
var coll *mongo.Collection
var mongoCtx context.Context

var mongoConfig = config.MongoConfig{
    MongoHost: config.GetEnv("MONGO_HOST", "101.132.45.28"),
    MongoPort: config.GetEnv("MONGO_PORT", "27017"),
    MongoDb:   config.GetEnv("MONGO_DB", databaseName),
    Username:  config.GetEnv("MONGO_USER", "root"),
    Password:  config.GetEnv("MONGO_PASS", "Yu1252068782?"),
}

func main() {
    // 1. start server and listen for serving requests
    listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatal(err.Error())
    }
    
    // 2. init mongo connection
    fmt.Println("Connecting to MongoDB...")
    mongoCtx = context.Background()
    uri := fmt.Sprintf("mongodb://%v:%v@%v:27017", mongoConfig.Username, mongoConfig.Password, mongoConfig.MongoHost)
    db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(uri))
    coll = db.Database(mongoConfig.MongoDb).Collection(collectionName)
    
    // 3. register service to grpc  server.
    // 3.1 generate security server connection
    creds, err := credentials.NewServerTLSFromFile("./integration/server/certs/cert.pem",
        "./integration/server/certs/key.pem")
    if err != nil {
        log.Fatal(err.Error())
    }
    options := []grpc.ServerOption{grpc.Creds(creds)}
    server := grpc.NewServer(options...)
    // 3.2 register service to server
    pb.RegisterEmployeeServiceServer(server, new(employeeService))
    // 3.3 start server
    log.Println("grpc server starting...")
    server.Serve(listener)
}

// this struct implemented EmployeeServiceServer.
type employeeService struct {
}

func (e employeeService) GetByNo(ctx context.Context, request *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
    panic("implement me")
}

func (e employeeService) GetByAll(request *pb.GetByAllRequest, stream pb.EmployeeService_GetByAllServer) error {
    data := &model.Employee{}
    cursor, err := coll.Find(context.Background(), bson.M{})
    if err != nil {
        return err
    }
    
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        err := cursor.Decode(data)
        if err != nil {
            return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
        }
        // If no error is found send blog over stream
        stream.Send(&pb.EmployeeResponse{
            Employee: &pb.Employee{
                Id:        data.Id.Hex(),
                No:        data.No,
                FirstName: data.FirstName,
                LastName:  data.LastName,
                MonthSalary: &pb.MonthSalary{
                    Basic: data.MonthSalary.Basic,
                    Bonus: data.MonthSalary.Bonus,
                },
                Status:       data.Status,
                LastModified: data.LastModified,
            },
        })
    }
    
    return nil
}

func (e employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
    // 1. get metadata
    md, ok := metadata.FromIncomingContext(stream.Context())
    if ok {
        log.Printf("employee : %s\n", md["id"][0])
    }
    
    // 2. receive file data
    img := []byte{}
    for {
        data, err := stream.Recv()
        if err == io.EOF {
            log.Printf("file size: %d\n", len(img))
            break
        }
        if err != nil {
            return err
        }
        log.Printf("receive data size: %d\n", len(data.Data))
        img = append(img, data.Data...)
    }
    
    // 3. update database according to the metadata
    oid, _ := primitive.ObjectIDFromHex(md["id"][0])
    filter := bson.M{"_id": oid}
    update := bson.M{
        "file": img,
    }
    result := coll.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": update})
    log.Println(result)
    
    return stream.SendAndClose(&pb.AddPhoneResponse{
        IsOk: true,
    })
}

func (e employeeService) Save(ctx context.Context, request *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
    res, err := coll.InsertOne(context.Background(), Dto2Model(request))
    if err != nil {
        return nil, status.Errorf(
            codes.Internal,
            fmt.Sprintf("Internal error when creating employee: %v", err),
        )
    }
    
    request.Employee.Id = res.InsertedID.(primitive.ObjectID).Hex()
    return &pb.EmployeeResponse{
        Employee: request.Employee,
    }, nil
}

func (e employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
    for {
        data, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        employee := Dto2Model(data)
        res, err := coll.InsertOne(context.Background(), employee)
        if err != nil {
            return status.Errorf(codes.Internal, "create record error: %v", err)
        }
        employee.Id = res.InsertedID.(primitive.ObjectID)
        stream.Send(Model2Dto(&employee))
    }
    
    return nil
}

func (e employeeService) GenerateToken(ctx context.Context, request *pb.TokenRequest) (*pb.TokenResponse, error) {
    panic("implement me")
}

func Dto2Model(request *pb.EmployeeRequest) model.Employee {
    oid, _ := primitive.ObjectIDFromHex(request.Employee.Id)
    
    return model.Employee{
        Id:        oid,
        No:        request.Employee.No,
        FirstName: request.Employee.FirstName,
        LastName:  request.Employee.LastName,
        MonthSalary: &model.MonthSalary{
            Basic: request.Employee.MonthSalary.Basic,
            Bonus: request.Employee.MonthSalary.Bonus,
        },
        Status:       request.Employee.Status,
        LastModified: request.Employee.LastModified,
    }
}

func Model2Dto(m *model.Employee) *pb.EmployeeResponse {
    return &pb.EmployeeResponse{
        Employee: &pb.Employee{
            Id:        m.Id.Hex(),
            No:        m.No,
            FirstName: m.FirstName,
            LastName:  m.LastName,
            MonthSalary: &pb.MonthSalary{
                Basic: m.MonthSalary.Basic,
                Bonus: m.MonthSalary.Bonus,
            },
            Status:       m.Status,
            LastModified: m.LastModified,
        },
    }
}
