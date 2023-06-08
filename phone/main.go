package main

import (
	"fmt"
	"phone/pkg/db"
)

func main() {
	dbConn := db.NewConnection()
	defer dbConn.Connection.Close()
	fmt.Println("Hello world")
}
