package main

import "github.com/gorilla/mux"
import "net/http"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":30824", nil)
}


func HomeHandler(response http.ResponseWriter , request *http.Request) {
	response.Write([]byte("Hello"))
}
