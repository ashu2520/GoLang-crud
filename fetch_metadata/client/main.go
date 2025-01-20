package main

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "fetch_metadata/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcClient pb.MetadataServiceClient

func SetMetadataServiceClient(c pb.MetadataServiceClient) {
	grpcClient = c
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewMetadataServiceClient(conn)
	SetMetadataServiceClient(client) // grpc client initialization

	r := gin.Default()
	setupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run Gin server: %v", err)
	}
	log.Println("REST API is running on port 8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/tables", GetTables) // creating routes
}

func GetTables(c *gin.Context) {
	// Query parameter se schemaName lena
	schemaName := c.Query("schema")
	if schemaName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Schema name is required"})
		return
	}

	// 10 seconds ka timeout context setup karna
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// gRPC client call karna
	resp, err := grpcClient.GetTables(ctx, &pb.GetTablesRequest{SchemaName: schemaName})
	if err != nil {
		log.Printf("Failed to fetch tables: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to fetch tables",
			"reason": err.Error(),
		})
		return
	}

	// Response ko JSON format mein bhejna
	c.JSON(http.StatusOK, gin.H{"tables": resp.Tables})
}

// protoc --go-grpc_out=. --go_out=. *.proto
