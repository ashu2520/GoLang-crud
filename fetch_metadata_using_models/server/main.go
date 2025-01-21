package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"fetch_metadata_using_models/config"

	"fetch_metadata_using_models/models"
)

type MetadataServer struct {
	DB *sql.DB // Encapsulating the the database with the server
}

func (s *MetadataServer) GetTables(schemaName string) (*models.DataStoreSchema, error) {
	// Database query jo humein execute karna hai...
	query := `
        SELECT TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, TABLE_TYPE, ROW_COUNT, BYTES, CREATED, LAST_ALTERED, LAST_DDL
        FROM information_schema.tables
        WHERE TABLE_SCHEMA = ?`

	rows, err := s.DB.Query(query, schemaName) // Executing the Query...
	if err != nil {
		return nil, fmt.Errorf("failed to query tables: %w", err)
	}
	defer rows.Close()

	// Temporary struct define kiye hai...
	type TableInfo struct {
		TableCatalog string
		TableSchema  string
		TableName    string
		TableType    string
		RowCount     sql.NullInt64
		Bytes        sql.NullInt64
		Created      sql.NullTime
		LastAltered  sql.NullTime
		LastDDL      sql.NullTime
	}
	// Why sql.NullInt64 and sql.NullTime?
	// => To Handle nullable database fields safely...

	// creating a global array jisme humlog rows ke values ko store karenge...
	var dataTables []*models.DataTables

	for rows.Next() {
		// rows.Next() method will throught each row that comes from the Database...
		var tableInfo TableInfo // Temporary struct jo humne just upar define kiya tha... Usko instance banaya hai...
		if err := rows.Scan(
			// rows.scan method maps the rows coming to the coressponding field
			&tableInfo.TableCatalog,
			&tableInfo.TableSchema,
			&tableInfo.TableName,
			&tableInfo.TableType,
			&tableInfo.RowCount,
			&tableInfo.Bytes,
			&tableInfo.Created,
			&tableInfo.LastAltered,
			&tableInfo.LastDDL,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Map the local struct to the model
		table := &models.DataTables{
			// Now again we will map the tableInfo(instance of struct jisme humne abhi value store kiya tha ek row ka)
			// to the models.DataTables(jo base models humne banaya tha models.go mein)
			Name:          tableInfo.TableName,
			TableType:     tableInfo.TableType,
			TotalRows:     uint64OrDefault(tableInfo.RowCount),
			DataLength:    uint64OrDefault(tableInfo.Bytes),
			CreatedAt:     uint64OrDefaultTime(tableInfo.Created),
			LastUpdatedAt: uint64OrDefaultTime(tableInfo.LastAltered),
			Schema:        tableInfo.TableSchema,
			Description:   fmt.Sprintf("Table catalog: %s", tableInfo.TableCatalog),
		}

		dataTables = append(dataTables, table) // Global variable jisko humne upar define kiya tha usme values ko append kar rahe hai...
	}

	return &models.DataStoreSchema{DataTables: dataTables}, nil
}

// Helper function to handle sql.NullInt64
func uint64OrDefault(value sql.NullInt64) uint64 {
	if value.Valid {
		return uint64(value.Int64)
	}
	return 0
}

// Helper function to handle sql.NullTime
func uint64OrDefaultTime(value sql.NullTime) uint64 {
	if value.Valid {
		return uint64(value.Time.Unix())
	}
	return 0
}

func main() {
	// database se connect kar rahe hai...
	db, db_err := config.ConnectDB()
	if db_err != nil {
		log.Fatalf("Failed to connect to Snowflake: %v", db_err)
	}
	defer db.Close()

	// Yaha pr hum jo client se incoming network connection hai usko listen kiya ja raha hai...
	// client se dial kiya ja raha hai or  serve pr listen
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}
	defer listener.Close()

	db_server := &MetadataServer{DB: db} //Initializing MetadataServer instance with the established database connection...

	for {
		// The server will continously listen for the new connection
		conn, err := listener.Accept() // connection establish hua then conn object return kar dega..
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn, db_server) // Ek go routine start kar rahe hai to handle connection concurrently...
	}
}

func handleConnection(conn net.Conn, db_server *MetadataServer) {
	defer conn.Close()

	buffer := make([]byte, 1024) // This allocates the a buffer size of 1024 bytes to hold incoming data...
	n, err := conn.Read(buffer)  // Read kar raha hai jo bhi buffer mein aa raha hai... And returns a number of bytes read(n)
	if err != nil {
		log.Printf("failed to read from connection: %v", err)
		return
	}

	schemaName := string(buffer[:n])               // converts the containing the clients request into string
	result, err := db_server.GetTables(schemaName) // get table function ko call kiya..
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("error: %v", err)))
		return
	}

	response, _ := json.Marshal(result)
	// Marshal returns the json encoding
	conn.Write(response)
}
