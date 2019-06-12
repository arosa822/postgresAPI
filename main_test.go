package main_test

import (
	"."
	"os"
	"testomg"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}

	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}
