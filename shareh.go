package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		http.FileServer(http.Dir("."))),
	)
}
