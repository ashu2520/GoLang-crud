package config

import (
	"database/sql"
	"fmt"

	_ "github.com/snowflakedb/gosnowflake" // Snowflake driver
)

func ConnectDB() (*sql.DB, error) {
	dsn := "username:password@VHAJAIJ-LH66728/VAPUS_PPD/PUBLIC?warehouse=COMPUTE_WH"

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Connected to Snowflake database successfully")
	return db, nil
}
