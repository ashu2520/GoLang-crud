package models // ye basically represent kar raha hai that this file belongs to models
// models package is typically defined when we have to data structure or schema...

// User represents the USERS table schema
type User struct {
	UserID       int
	UserName     string
	UserMobile   string
	UserEmail    string
	UserGender   string
	UserCountry  string
	UserState    string
	UserStatus   string
	UserPassword string
	UserTerms    string
	UserType     string
	UserRoleID   int
}

// struct is used to define the custom defined data type...
