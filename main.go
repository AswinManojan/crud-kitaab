package main

import (
	"github.com/sample-crud-app/di"
)

func main() {
	server := di.Init()
	server.StartServer()
}
