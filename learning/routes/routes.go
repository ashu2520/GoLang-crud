package routes

import (
	"learning/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes
func SetupRoutes(r *gin.Engine) {
	// r => variable name
	// *gin.Engine => type of the variable 'r'
	r.POST("/users", controllers.InsertUser)
	r.GET("/getusers", controllers.ReadUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
}

// package routes

// import (
// 	"learning/controllers"

// 	"github.com/gin-gonic/gin"
// )

// // SetupRoutes defines all API endpoints for the application.
// func SetupRoutes(r *gin.Engine) {
// 	// Group routes under a common path
// 	api := r.Group("/api/v1") // Versioning the API
// 	{
// 		// Define routes and associate them with handler functions
// 		api.GET("/users", controllers.GetUsers)       // Fetch all users
// 		api.POST("/users", controllers.CreateUser)    // Create a new user
// 		api.GET("/users/:id", controllers.GetUser)    // Fetch user by ID
// 		api.PUT("/users/:id", controllers.UpdateUser) // Update user by ID
// 		api.DELETE("/users/:id", controllers.DeleteUser) // Delete user by ID
// 	}

// 	// Example of another route group for authentication
// 	auth := r.Group("/auth")
// 	{
// 		auth.POST("/login", controllers.Login)
// 		auth.POST("/register", controllers.Register)
// 	}
// }
