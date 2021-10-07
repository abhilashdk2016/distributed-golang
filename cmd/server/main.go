package main

import (
	"log"

	"github.com/abhilashdk2016/distributed-golang/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
