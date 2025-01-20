package config

import (
	"database/sql"
	"fmt"

	_ "github.com/snowflakedb/gosnowflake" // Snowflake driver
)

// ConnectDB establishes a connection to the Snowflake database
func ConnectDB() (*sql.DB, error) {
	// sabse phale to ek function banaye ConnectDB() usme humlog do chiz return kare rahe honge..
	// ConnectDB => ye start ho raha hai capital letter se which means it can be exprted.
	// In Golang the capitalization of the first letter of a function, vairable, type, or constant determines its visiblity.
	// And we are returning two things:-
	// 1) pointer(*sql.DB) => its a pointer to an sql.DB object. and it is representing the database connection
	// 2) error
	dsn := "username:password@VHAJAIJ-LH66728/VAPUS_PPD/PUBLIC?warehouse=COMPUTE_WH" // dsn => Data Source Name
	// ye basically, snowflake se connection ke liye hai...

	db, err := sql.Open("snowflake", dsn) // Yaha pr basically attempt kar rahe hai, to open a database connection using sql.Open
	// sql.Open => Open, sql package ke ander ka koi ek function hai...
	// Open mein do arguments de rahe hai...
	// 1) database driver name => "snowflake"
	// 2) driver specific data sourse name => dsn
	// Open return bhi do chiz kar raha hai...
	// 1) pointer to a sql.DB => jisko hum 'db' mein store kar rahe hai...
	// 2) error => jisko humlog 'err' mein store kar rah hai...
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	// * NOTE:- Open may just validate the arguments without connecting to the database...
	// sql.Open => validate the DSN and prepare the connection pool.
	// A real connection is established only when a query or db.Ping is executed.
	// Ping uses Context,Background internally.
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil // yaha pr humlog return kare hai poiter(db) and nil
	// pointer store the address value.
	// so we are returning adress of the DB.
}
