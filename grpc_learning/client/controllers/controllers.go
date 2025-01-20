package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	proto "grpc_learning/proto"

	"github.com/gin-gonic/gin"
)

var client proto.UserServiceClient

// SetUserServiceClient sets the gRPC client
func SetUserServiceClient(c proto.UserServiceClient) {
	// SetUserServiceClient: ye function call hua hai clinet/main.go se
	// interface is type that lists method without providing their code
	// UserServiceClient ek interface hai, jo ki hold karta hai sare function ka declaration jo humne service mein define kiya tha...
	client = c
}

// CreateUser handles user creation
func CreateUser(c *gin.Context) {
	// Parse the JSON body
	var user proto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	// fmt.Println(user) //{{{} [] [] <nil>} 0           [] 0}
	// Context for gRPC call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call gRPC CreateUser
	res, err := client.CreateUser(ctx, &proto.CreateUserRequest{User: &user})
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Send the response
	c.JSON(http.StatusOK, gin.H{"message": res.GetMessage()})
}

func ReadUsers(c *gin.Context) {
	// Context for gRPC call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call gRPC GetUsers without pagination
	res, err := client.GetUsers(ctx, &proto.GetUsersRequest{})
	if err != nil {
		log.Printf("Failed to retrieve users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Send the response
	c.JSON(http.StatusOK, gin.H{"users": res.GetUsers()})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var req proto.UpdateUserRequest    // "UpdateUserRequest" struct ka instance bana rahe hai...
	json_err := c.ShouldBindJSON(&req) // ye bond karta hai http json request ko Go struct ke sath.
	if json_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaild Input Data"})
		return
	}

	UserId, typeErr := strconv.ParseInt(userID, 10, 32)
	if typeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id was not declared"})
		return
	}

	// Convert int64 to int32
	req.UserId = int32(UserId) // Casting the int64 to int32

	// Agar 10 second mein execute nhi hua to autmatic request ko cancel kar do.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// ctx: Its a derived context that include the timeout.
	// cancel: its a function that you can call to manually cancel the the context with no canellation.
	defer cancel()

	res, err := client.UpdateUser(ctx, &req)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "reason": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res.GetMessage()})
}

func DeleteUser(c *gin.Context) {
	// Parse the user ID from the request body
	var req proto.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Context for gRPC call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call gRPC DeleteUser
	res, err := client.DeleteUser(ctx, &req)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Send the response
	c.JSON(http.StatusOK, gin.H{"message": res.GetMessage()})
}
