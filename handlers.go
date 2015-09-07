package main

import (
	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

// 404 error handler
func NotFoundHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
    response.WriteHeader(http.StatusNotFound)
	
	notFoundError := NotFoundError { Error: "404" }
	
	err := json.NewEncoder(response).Encode(notFoundError);
	
	if err != nil {
		panic(err)
	}
}

// Index
func IndexHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
    response.WriteHeader(http.StatusOK)
}

// Ex: /search/wot
func SearchHandler(response http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//query := vars["query"]
	
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
    response.WriteHeader(http.StatusOK)
	
	videos := Videos {
		Video {
			Name: "Sample Video",
			Id: "62yAkwym",
			Url: "http://mover.uz/watch/62yAkwym/",
			DirectUrl: "http://v.mover.uz/62yAkwym_s.mp4",
			VideoHostName: "mover",
		},
	}
	
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
    response.WriteHeader(http.StatusOK)
	
	err := json.NewEncoder(response).Encode(videos);
	
	if err != nil {
		panic(err)
	}
}

// Ex: /view/mover/62yAkwym
func ViewHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	videoHost := vars["videoHost"]
	videoId := vars["videoId"]
	
	video := Video {
		Name: "Sample Video",
		Id: videoId,
		Url: "http://mover.uz/watch/62yAkwym/",
		DirectUrl: "http://v.mover.uz/62yAkwym_s.mp4",
		VideoHostName: videoHost,
	}
	
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
    response.WriteHeader(http.StatusOK)
	
	err := json.NewEncoder(response).Encode(video);
	
	if err != nil {
		panic(err)
	}
}