package config

import (
	"database/sql"
	"fmt"

	_ "github.com/snowflakedb/gosnowflake" // Snowflake driver
)

// ConnectDB establishes a connection to the Snowflake database
func ConnectDB() (*sql.DB, error) {
	// Define the Data Source Name (DSN) for the Snowflake connection
	dsn := "username:password@VHAJAIJ-LH66728/VAPUS_PPD/PUBLIC?warehouse=COMPUTE_WH"

	// Open the database connection
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Ping the database to ensure the connection is valid
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Connected to Snowflake database successfully")
	return db, nil
}
