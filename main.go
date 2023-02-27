package main

import (
	"log"

	"github.com/thecodedproject/resourcegen/internal"
)

func main() {

	err := internal.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
}
