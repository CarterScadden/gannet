package main

import (
	"fmt"
	"gannet/server"
	"log"
	"net/http"
)

const PORT int = 4000

func main() {
	server := server.New(PORT)
	port := fmt.Sprintf(":%d", PORT)

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, server))
}
