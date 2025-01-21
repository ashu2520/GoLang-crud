package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"fetch_metadata_using_models/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin ko initiaalize kiye and port assign kiye...
	r := gin.Default()
	RegisterRoutes(r)
	r.Run(":8080")
}

func RegisterRoutes(r *gin.Engine) {
	// route create kiye...
	r.GET("/tables", GetTablesHandler)
}

func GetTablesHandler(c *gin.Context) {
	// schema get and check kiye...
	schemaName := c.Query("schema")
	if schemaName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "schema name is required"})
		return
	}

	conn, err := net.Dial("tcp", "localhost:50051") // TCP connection bana rahe hai server ke sath...
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to server"})
		return
	}
	defer conn.Close()

	conn.Write([]byte(schemaName)) // sends the schema name to the server...
	// Write => Write function writes the data to connection
	buffer := make([]byte, 4096) // It allocate the buffer to read the connection...

	n, err := conn.Read(buffer) //Reads the data from the server...
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response from server"})
		return
	}
	var result models.DataStoreSchema
	log.Println("Received response:", string(buffer[:n]))
	// json.Unmarshal(buffer[:n], &result); => Parse the JSON response from the server into a models.DataStoresSchema...
	if err := json.Unmarshal(buffer[:n], &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse server response"})
		return
	}

	c.JSON(http.StatusOK, result) //Return the parsed result as a json response with a 200 OK status
}
