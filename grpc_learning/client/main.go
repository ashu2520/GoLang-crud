package main

import (
	"log"

	"grpc_learning/client/controllers"
	proto "grpc_learning/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Initialize gRPC client and inject it into controllers
	client := proto.NewUserServiceClient(conn)
	controllers.SetUserServiceClient(client)

	// Initialize Gin router
	r := gin.Default()
	setupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run Gin server: %v", err)
	}
}

// SetupRoutes defines all the routes
func setupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/getusers", controllers.ReadUsers)
	r.PUT("/update-users/:id", controllers.UpdateUser)
	r.DELETE("/delete-users", controllers.DeleteUser)
}

// protoc --go-grpc_out=. --go_out=. *.proto
