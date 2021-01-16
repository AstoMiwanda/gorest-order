package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var addr = os.Getenv("HTTP_ADDR_PORT")
	log.Printf("Running http server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, RestRouter()))
}
