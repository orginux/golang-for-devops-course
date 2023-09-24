package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: ./http-get <url>")
		os.Exit(1)
	}
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println("URL is not valid:", err)
		os.Exit(1)
	}
	response, err := http.Get()
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
}
