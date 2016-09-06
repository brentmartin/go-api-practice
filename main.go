package main

import (
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

