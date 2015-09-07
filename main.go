package main

import (
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8000"
	}
	
	err := http.ListenAndServe(":" + port, router)
	
	if err != nil {
		panic(err)
	}
}