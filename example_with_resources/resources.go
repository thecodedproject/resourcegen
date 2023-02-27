package example_with_resources

import (
	"database/sql"
)

//go:generate go run ../main.go --struct_name=resources
type resources struct {
	myInt int
	aFloat float64
	someDBConnection *sql.DB
}

func New() *resources {

	return &resources{
		myInt: 1,
		aFloat: 2.0,
		// in reality we'd create a DB connection here
		someDBConnection: nil,
	}
}
