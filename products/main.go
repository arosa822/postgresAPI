package main

import (
	"fmt"
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	fmt.Printf("\nConnecting to DB...\n")
	fmt.Printf("%s\n", os.Getenv("APP_DB_USERNAME"))

	a.Run(":5432")

}
