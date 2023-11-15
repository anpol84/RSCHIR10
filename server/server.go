package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	pb "practice10/gen/proto"
	"practice10/service1"
	"practice10/service2"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Encrypt(ctx context.Context, req *pb.Crypt) (*pb.Crypt, error) {
	key := []byte("0123456789abcdef")
	msg, _ := service1.EncryptString(key, req.Msg)
	return &pb.Crypt{Msg: msg}, nil
}

func (s *testApiServer) Decrypt(ctx context.Context, req *pb.Crypt) (*pb.Crypt, error) {
	key := []byte("0123456789abcdef")
	msg, _ := service1.DecryptString(key, req.Msg)
	return &pb.Crypt{Msg: msg}, nil
}
func (s *testApiServer) GetFiles(ctx context.Context, req *pb.EmptyMessage) (*pb.FilesResponse, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	files, _ := service2.GetFiles(client, "files")
	return &pb.FilesResponse{Files: files}, nil
}

func (s *testApiServer) GetFileInfo(ctx context.Context, req *pb.IdRequest) (*pb.File, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	file := service2.GetFileInfo(client, "files", req.Id)
	return &pb.File{Id: file.Id, Name: file.Name, Description: file.Description}, nil
}

func (s *testApiServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.EmptyMessage, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ := mongo.Connect(ctx, clientOptions)

	log.Println(req)
	service2.UploadFile(client, "files", req.Name, req.Description)
	return &pb.EmptyMessage{}, nil
}

func (s *testApiServer) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.EmptyMessage, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ := mongo.Connect(ctx, clientOptions)

	log.Println(req)
	service2.UpdateFile(client, req.Id, req.Name, req.Description, "files")
	return &pb.EmptyMessage{}, nil
}

func (s *testApiServer) DeleteFile(ctx context.Context, req *pb.IdRequest) (*pb.EmptyMessage, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	log.Println(req)
	service2.DeleteFile(client, req.Id, "files")
	return &pb.EmptyMessage{}, nil
}

func main() {
	port := flag.String("port", "8081", "port to run the server on")
	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})
		log.Println("Server started on port " + *port)
		log.Fatalln(http.ListenAndServe("localhost:"+*port, mux))
	}()

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})
	log.Println("It all works")
	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalln(err)
	}
}
