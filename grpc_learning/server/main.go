package main

import (
	"database/sql"
	"grpc_learning/config"

	proto "grpc_learning/proto"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Initialize the database connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Start the gRPC server and pass the database connection
	if err := StartGRPCServer(db); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func StartGRPCServer(db *sql.DB) error {
	// Create a new gRPC server instance(jaise object is an instance of class)
	// serrver instance ka matlab hai real time mein server ke liye banana...
	grpcServer := grpc.NewServer() //  NewServer creates a grpc server which has no service registered and has not started to accept the requst yet.
	// grpcServer ek thrha ka pointer hai...
	// type of grpcServer is *grpc.Server

	// Register the UserServiceServer with the gRPC server
	proto.RegisterUserServiceServer(grpcServer, &UserServiceServer{DB: db})
	// humne kya kiya tha??
	// humne ek service define kari thi "user.proto" file mein UserService ke naam se...

	// RegisterUserServiceServer function defined hai "user_grpc.pb.go" mein..
	// func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer)

	// ab grpc.ServiceRegister kya karta hai???
	// => ye allow karta hai humein to register with grpc server.
	// ServiceRegister makes it possible to use others types of server too...

	// ab UserServiceServer ye kya karta hai???
	// ye UserServiceServer ek dataType hai which is defined through struct keyword in "user_service.go"...
	// service ke ander humne jitni bhi function banayi thi un sabko db sath connect kar raha hai...

	// over all UserServiceServer is an interface for your grpc Service.

	// Start listening on a specific port
	lis, err := net.Listen("tcp", ":50051")
	// its just opens a TCP listener on port 50051
	if err != nil {
		return err
	}

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50051...")
	errors := grpcServer.Serve(lis)
	if errors != nil {
		return errors
	}
	// grpcServer.Serve(lis); => it starts the GRPC server and binds it to the listener.
	return nil
}

// go run (Get-ChildItem -Filter *.go).FullName  => for running all the files in a package with same
