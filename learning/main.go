package main

import (
	"fmt"
	"learning/config"
	"learning/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	// config package se ConnectDb() function ko call kiye...
	db, err := config.ConnectDB()
	// config.ConnectDb => do chiz return kiya hai
	// db => pointer jisme database connection ka address stored hai..
	// err => agar koi error aya to...
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	// defer db.Close() => program ke sabse last mein db connection ko close kar do...
	fmt.Println("Connected to Snowflake!")

	// Initialize Gin router
	// gin is a HTTP web framework for building RESTAPIs in GO.
	r := gin.Default() // Middleware
	// Here we are creating a GIN "router" instance, which comes with built-in middleware for logging request and recovery from panic
	// It creates the router with some predefined middleware (logging and recovery)
	// fmt.Printf("Type of r is %T\n", r)  // Type of r is *gin.Engine
	// r is *gin.Engine => yani ki 'r' pointer hai...

	// Pass the database to routes as middleware
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // it is used to new key value pair. It stores the database connection under the key "db"
		c.Next()        // It ensure the request is passed to next middleware or route handler in the chain
	})

	// Set up API routes
	routes.SetupRoutes(r)
	// routes package ke ander SetupRoutes() ka function defined hai...
	// to uss function(SetupRoutes()) mein hum 'r' yani ki router ka instance bhej rahe hai. So, that hum waha pr routes create kar sake...

	// Start the HTTP server
	r.Run(":8080") // Change port if needed
}
