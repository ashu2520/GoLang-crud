package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"fetch_metadata/config"
	pb "fetch_metadata/proto"

	"google.golang.org/grpc"
)

type metadataServer struct {
	pb.UnimplementedMetadataServiceServer
	db *sql.DB
}

func (s *metadataServer) GetTables(ctx context.Context, req *pb.GetTablesRequest) (*pb.GetTablesResponse, error) {

	query := `
    SELECT TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, TABLE_TYPE, ROW_COUNT, BYTES, CREATED, LAST_ALTERED, LAST_DDL
    FROM information_schema.tables
    WHERE TABLE_SCHEMA = ?;
`

	rows, err := s.db.Query(query, req.SchemaName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tables: %w", err)
	}
	defer rows.Close()

	var tables []*pb.TableInfo
	for rows.Next() {
		var table pb.TableInfo
		if err := rows.Scan(
			&table.Catalog, &table.Schema, &table.Name, &table.Type,
			&table.RowCount, &table.Bytes, &table.CreatedAt, &table.LastAltered, &table.LastDdl,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tables = append(tables, &table)
	}
	return &pb.GetTablesResponse{Tables: tables}, nil
}

func main() {
	db, err := config.ConnectDB() // Use your Snowflake connection logic
	if err != nil {
		log.Fatalf("Failed to connect to Snowflake: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterMetadataServiceServer(server, &metadataServer{db: db})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server is running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
