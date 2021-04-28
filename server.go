package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Version = "undefined"

const TEMPLATE_FILE = "templates.json"

func main() {
	fmt.Println("Version:", Version)

	port := flag.String("p", "8080", "port to serve on")
	filePath := flag.String("f", TEMPLATE_FILE, "path to "+TEMPLATE_FILE+" file")
	flag.Parse()

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		log.Fatal(*filePath + " file not found")
	}

	http.HandleFunc("/"+TEMPLATE_FILE, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *filePath)
	})

	log.Println("Serving on HTTP port:", *port)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
