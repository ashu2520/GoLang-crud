package main

import (
	"context"
	"database/sql"
	"fmt"
	proto "grpc_learning/proto"
	"log"
	"time"

	_ "github.com/snowflakedb/gosnowflake"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	DB *sql.DB
}

// CreateUser implementation
func (s *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user := req.GetUser()
	log.Printf("Request Body %v", user)
	fmt.Println(user)
	query := `
		INSERT INTO USERS (USER_ID, USER_NAME, USER_MOBILE, USER_EMAIL, USER_GENDER, USER_COUNTRY, USER_STATE, USER_PASSWORD)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.DB.Exec(query, user.UserId, user.UserName, user.UserMobile, user.UserEmail, user.UserGender, user.UserCountry, user.UserState, user.UserPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	return &proto.CreateUserResponse{Message: "User inserted successfully"}, nil
}

// GetUsers implementation
func (s *UserServiceServer) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	query := "SELECT USER_ID, USER_NAME, USER_MOBILE, USER_EMAIL FROM USERS"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to read users: %v", err)
	}
	defer rows.Close()

	var users []*proto.User
	for rows.Next() {
		var user proto.User
		if err := rows.Scan(&user.UserId, &user.UserName, &user.UserMobile, &user.UserEmail); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, &user)
	}

	return &proto.GetUsersResponse{Users: users}, nil
}

// UpdateUser implementation
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	fmt.Println(*req.NewStatus, time.Now(), req.UserId)
	query := "UPDATE USERS SET USER_STATUS = ?, USER_UPDATED_AT = ? WHERE USER_ID = ?"
	_, err := s.DB.Exec(query, req.NewStatus, time.Now(), req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return &proto.UpdateUserResponse{Message: "User updated successfully"}, nil
}

// DeleteUser implementation
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	query := "DELETE FROM USERS WHERE USER_ID = ?"
	_, err := s.DB.Exec(query, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}

	return &proto.DeleteUserResponse{Message: "User deleted successfully"}, nil
}
