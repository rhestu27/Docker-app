package main

import (
    "fmt"
    "net/http"
    "os"
	"log"
)

func main() {
    http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == ""{
			port = "8080"
	}
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    name := os.Getenv("NAME")
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Halo, %s!", name)
	fmt.Fprintf(w, "Welcome to Cloud Computing Class!\n")
	fmt.Fprintf(w, "This is a simple Dockerized Go app\n")
}


