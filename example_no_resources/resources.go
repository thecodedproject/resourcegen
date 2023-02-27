package example_no_resources

import (
)

//go:generate go run ../main.go --struct_name=resources
type resources struct {
}

func New() *resources {

	return nil
}
