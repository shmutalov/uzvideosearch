package main

import (
	"net/http"
	
	"github.com/gorilla/mux"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("UzVideoSearch"))
}

func SearchHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	query := vars["query"]
	
	response.Write([]byte(query))
}

func ViewHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	
	response.Write([]byte(name))
}

func main() {
	r := mux.NewRouter()
	
	// Index
	r.HandleFunc("/", IndexHandler)
	
	// Search <query> string
	r.HandleFunc("/search/{query}/", SearchHandler)
	
	// View <name> video
	r.HandleFunc("/view/{name}/", ViewHandler)
	
	err := http.ListenAndServe(":8000", r)
	
	if err != nil {
		panic(err)
	}
}