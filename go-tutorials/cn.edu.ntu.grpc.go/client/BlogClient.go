package main

import (
	blogpb "cn.edu.ntu.grpc.go/proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

const port = ":50051"

func main() {
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("localhost"+port, options...)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	var client = NewBlogServiceClient(conn)

	// create request
	cRequest := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "zack",
			Title:    "title",
			Content:  "content",
		},
	}

	// 0. use generated code
	c := blogpb.NewBlogServiceClient(conn)
	b, e := c.CreateBlog(context.Background(), cRequest)
	if e != nil {
		log.Fatalln(err.Error())
	}
	log.Println(b.GetBlog())

	// 1. create
	for i := 0; i < 5; i++ {
		client.create(cRequest)
	}

	// 5. query list
	list, err := client.getAll(new(blogpb.ListBlogsRequest))
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, blog := range list {
		log.Println(blog)
	}

	oid, err := client.create(cRequest)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 2. update
	uBlog := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       oid,
			AuthorId: "zack-updated",
			Title:    "title-updated",
			Content:  "content-updated",
		}}

	client.update(uBlog)

	// 3. get
	gBlog := &blogpb.GetBlogRequest{Id: oid}
	blog, err := client.get(gBlog)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(blog)

	// 4. delete
	isSuccess, err := client.delete(&blogpb.DeleteBlogRequest{Id: oid})
	if err != nil || isSuccess != true {
		log.Fatalf("error occur: %v, isSuccess: %v\n", err.Error(), isSuccess)
	}
	log.Printf("delete is success: %v\n", isSuccess)

	log.Println("success")
}

type BlogServiceClient struct {
	cc blogpb.BlogServiceClient
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) *BlogServiceClient {
	aa := blogpb.NewBlogServiceClient(cc)
	return &BlogServiceClient{aa}
}

func (client *BlogServiceClient) create(request *blogpb.CreateBlogRequest) (string, error) {
	blog, err := client.cc.CreateBlog(context.Background(), request)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(blog)

	return blog.GetBlog().Id, err
}

func (client *BlogServiceClient) update(request *blogpb.UpdateBlogRequest) (string, error) {
	blog, err := client.cc.UpdateBlog(context.Background(), request)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(blog)

	return blog.GetBlog().Id, err
}

func (client *BlogServiceClient) get(request *blogpb.GetBlogRequest) (*blogpb.Blog, error) {
	blog, err := client.cc.GetBlog(context.Background(), request)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(blog)

	return blog.GetBlog(), err
}

func (client *BlogServiceClient) getAll(request *blogpb.ListBlogsRequest) ([]*blogpb.Blog, error) {
	stream, err := client.cc.ListBlogs(context.Background(), request)
	if err != nil {
		println(stream)
	}

	var list []*blogpb.Blog
	for {
		res, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		list = append(list, res.GetBlog())
	}

	return list, nil
}

func (client *BlogServiceClient) delete(request *blogpb.DeleteBlogRequest) (bool, error) {
	blog, err := client.cc.DeleteBlog(context.Background(), request)

	if err != nil {
		return false, err
	}

	return blog.GetSuccess(), nil
}
