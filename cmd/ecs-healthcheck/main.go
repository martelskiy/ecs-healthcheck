package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var path string

	flag.StringVar(&path, "path", "/status", "Relative HTTP path to health check")
	flag.Parse()

	resp, err := http.Get("http://localhost" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("got unexpected status: ", resp.StatusCode)
	}
}
