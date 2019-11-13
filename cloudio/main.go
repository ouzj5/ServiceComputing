package main

import (
	"cloudio/service"
)

func main() {
	server := service.NewServer()
	server.Run(":8080")
}
