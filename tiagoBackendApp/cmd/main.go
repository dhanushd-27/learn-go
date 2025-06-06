package main

import (
	"fmt"
	"log"

	"github.com/dhanushd-27/learn-go/tiagoBackendApp/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on port 8080")
}