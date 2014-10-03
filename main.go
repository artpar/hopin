package main

import "github.com/gorilla/mux"
import (
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":30824", nil)
}


func HomeHandler(response http.ResponseWriter , request *http.Request) {
	//	vars := mux.Vars(request)
	email := request.FormValue("email")
	regId := request.FormValue("regid")
	fmt.Printf("vars - %s - %s\n", email, regId)
	if (len(email) == 0 || len(regId) == 0) {
		response.Write([]byte("Not Ok"))
		return
	}
	response.Write([]byte("Ok"))
}
