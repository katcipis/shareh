package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var dir string
	var port int

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.IntVar(&port, "port", 8080, "port to listen")
	flag.StringVar(&dir, "dir", wd, "directory to serve files")
	flag.Parse()

	log.Printf("Serving %q listening on %d\n", dir, port)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		http.FileServer(http.Dir("."))),
	)
}
