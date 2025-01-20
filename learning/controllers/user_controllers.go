package controllers

import (
	"database/sql"
	"fmt"
	"learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InsertUser inserts a new user into the USERS table
func InsertUser(c *gin.Context) {
	// Parse the JSON body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve the database connection from context
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	// Insert query
	query := `
		INSERT INTO USERS (
			USER_ID, USER_NAME, USER_MOBILE, USER_EMAIL, USER_GENDER,
			USER_COUNTRY, USER_STATE, USER_PASSWORD
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, user.UserID, user.UserName, user.UserMobile, user.UserEmail, user.UserGender, user.UserCountry, user.UserState, user.UserPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User inserted successfully"})
}

// ReadUsers retrieves all users from the USERS table
func ReadUsers(c *gin.Context) {
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	query := "SELECT USER_ID, USER_NAME, USER_MOBILE, USER_EMAIL FROM USERS"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read users: %v", err)})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.UserMobile, &user.UserEmail); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to scan row: %v", err)})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser updates a user's status
func UpdateUser(c *gin.Context) {
	// Parse the user ID from the URL parameter
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse the request body for new status
	var input struct {
		NewStatus string `json:"new_status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	query := "UPDATE USERS SET USER_STATUS = ? WHERE USER_ID = ?"
	_, err = db.Exec(query, input.NewStatus, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	query := "DELETE FROM USERS WHERE USER_ID = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
